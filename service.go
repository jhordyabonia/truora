package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	s "strings"
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
	Endpoints       [3]Endpoints
}
type OutServer struct {
	Address   string
	Ssl_grade string
	Country   string
	Owner     string
}
type Out struct {
	Servers            [3]OutServer
	Ssl_grade          string
	Previous_ssl_grade string
	Logo               string
	Title              string
	Is_down            bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
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
func scrap(host string) (is_down bool, title string, logo string) {
	//url := host
	req, err := http.NewRequest("GET", host, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	is_down = !s.Contains("200", resp.Status)
	if is_down {
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var a1, a2 = "<title>", "</title>"
	if s.Contains(string(body), a1) && s.Contains(string(body), a2) {
		title = s.Split(string(body), a1)[1]
		title = s.Split(title, a2)[0]
		title = s.TrimSpace(title)
	}
	var b1, b2 = "<link", "type=\"image/x-icon\""
	if s.Contains(string(body), b1) && s.Contains(string(body), b2) {
		logo0 := s.Split(string(body), b1)
		for i := 0; i < len(logo0); i++ {
			if s.Contains(logo0[i], b2) {
				logo1 := s.Split(string(logo0[i]), "\"")
				for ii := 0; ii < len(logo1); ii++ {
					if s.Contains(logo1[ii], ".png") {
						logo = logo1[ii]
					}
				}
			}
		}
	}
	return
}
func One(host string) string {

	url := "https://api.ssllabs.com/api/v3/analyze?host=" + host
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:" , resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	data := &Result{}
	json.Unmarshal([]byte(string(body)), data)
	fmt.Printf("Endpoints[0].ServerName: %s\n", data.Endpoints[0].ServerName)

	leng_servers := len(data.Endpoints)
	//servers := [3]OutServer{}
	out := &Out{}
	//out.Servers = servers
	country, woner := getOwner(host)
	for i := 0; i < leng_servers; i++ {
		out.Servers[i] = OutServer{}
		out.Servers[i].Address = data.Endpoints[i].IpAddress
		out.Servers[i].Ssl_grade = data.Endpoints[i].Grade
		out.Servers[i].Country = country
		out.Servers[i].Owner = woner
	}
	out.Is_down, out.Title, out.Logo = scrap(host)
	out_json, err := json.Marshal(out)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(out_json)
}
