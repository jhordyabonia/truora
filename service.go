package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	s "strings"
	//"honnef.co/go/js/dom"
)

type Endpoints struct {
	IpAddress         string
	ServerName        string
	StatusMessage     string
	Grade             string
	GradeTrustIgnored string
	HasWarnings       string
	IsExceptional     string
	Progress          string
	Duration          string
	Delegation        string
}

/*
Represantacion de la estructura de datos resultados de  la  api
https://api.ssllabs.com/api/v3/analyze
*/
type Result struct {
	Host            string
	Port            string
	Protocol        string
	IsPublic        string
	Status          string
	StartTime       string
	TestTime        string
	EngineVersion   string
	CriteriaVersion string
	Endpoints       []Endpoints
}
type OutServer struct {
	Address   string
	Ssl_grade string
	Country   string
	Owner     string
}

/*
represeantacion de los datos, que conforman el detalle de cada dominio
*/
type Out struct {
	Servers            []OutServer
	Servers_changed    bool
	Ssl_grade          string
	Previous_ssl_grade string
	Logo               string
	Title              string
	Is_down            bool
}

/*
Lee del html de la pagina alojada en 'host', los datos
logo, title y si esta o no en linea
*/
func Html(host string) (is_down bool, title string, logo string) {
	/*html, is_down := GetPage(host)
	document := dom.GetWindow().FrameElement()
	document.SetInnerHTML(html)
	titleElement := document.QuerySelector("title")
	title := titleElement.OuterHTML()
	fmt.Println(title)*/
	return
}

/*
Compara si el objeto in1 traido del web-service,
a cambiado con respecto al objeto almacenado en base de datos
*/
func Compare(in1, in2 Out) (out Out) {
	out = in1
	out.Servers_changed = len(in1.Servers) != len(in2.Servers)

	if in1.Ssl_grade != in2.Ssl_grade {
		out.Ssl_grade = in2.Ssl_grade
		out.Servers_changed = true
	}
	out.Previous_ssl_grade = in2.Ssl_grade

	if !out.Servers_changed {
		for i := 0; i < len(out.Servers); i++ {
			if in1.Servers[i].Address != in2.Servers[i].Address {
				out.Servers_changed = true
			} else if in1.Servers[i].Ssl_grade != in2.Servers[i].Ssl_grade {
				out.Servers_changed = true
			} else if in1.Servers[i].Country != in2.Servers[i].Country {
				out.Servers_changed = true
			} else if in1.Servers[i].Owner != in2.Servers[i].Owner {
				out.Servers_changed = true
			}
		}
	}
	return
}

/*
Ejecuta el comando whois y extrae los valores de
country y woner, para el host, pasado en el paramtro
*/
func getOwner(host string) (country string, woner string) {

	out, err := exec.Command("whois", host).Output()
	if err != nil {
		log.Fatal(err)
	}
	var a1, a2 = "Registrant Country:", "Registrant Phone:"
	if s.Contains(string(out), a1) && s.Contains(string(out), a2) {
		country = s.Split(string(out), a1)[1]
		country = s.Split(country, a2)[0]
		country = s.TrimSpace(country)
	}
	var b1, b2 = "Reseller:", "Domain Status:"
	if s.Contains(string(out), b1) && s.Contains(string(out), b2) {
		woner = s.Split(string(out), b1)[1]
		woner = s.Split(woner, b2)[0]
		woner = s.TrimSpace(woner)
	}

	return
}

/*
Obtine el html resultante, de una direccion web (url)
*/
func GetPage(url string) (body string, errOut bool) {
	url = "http://" + s.Replace(url, "http://", "", 1)
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	errOut = !s.Contains(resp.Status, "200")
	if errOut {
		return
	}
	_body, _ := ioutil.ReadAll(resp.Body)
	body = string(_body)
	return
}

/* Sustituto temporal de Html*/
func scrap(host string) (is_down bool, title string, logo string) {
	body, is_down := GetPage(host)
	var a1, a2 = "<title>", "</title>"
	if s.Contains(body, a1) && s.Contains(body, a2) {
		title = s.Split(body, a1)[1]
		title = s.Split(title, a2)[0]
		title = s.TrimSpace(title)
	}
	var b1, b2 = "<link", "rel=\"shortcut icon\""
	if s.Contains(body, b1) && s.Contains(body, b2) {
		logo0 := s.Split(body, b1)
		for i := 0; i < len(logo0); i++ {
			if s.Contains(logo0[i], b2) {
				logo1 := s.Split(string(logo0[i]), "\"")
				for ii := 0; ii < len(logo1); ii++ {
					if s.Contains(logo1[ii], ".png") {
						logo = logo1[ii]
					} else if s.Contains(logo1[ii], ".ico") {
						logo = logo1[ii]
					}
				}
			}
		}
	}
	return
}
func ApiAnalyce(host string) (out Out) {
	url := "https://api.ssllabs.com/api/v3/analyze?host=" + host
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	data := &Result{}
	json.Unmarshal([]byte(string(body)), data)

	leng_servers := len(data.Endpoints)
	out.Servers = make([]OutServer, leng_servers)

	country, woner := getOwner(host)
	for i := 0; i < leng_servers; i++ {
		if out.Ssl_grade == "" {
			out.Ssl_grade = data.Endpoints[i].Grade
		}
		out.Servers[i] = OutServer{}
		out.Servers[i].Address = data.Endpoints[i].IpAddress
		out.Servers[i].Ssl_grade = data.Endpoints[i].Grade

		//country, woner := getOwner(data.Endpoints[i].ServerName)
		out.Servers[i].Country = country
		out.Servers[i].Owner = woner
	}
	out.Is_down, out.Title, out.Logo = scrap(host)
	return out
}
