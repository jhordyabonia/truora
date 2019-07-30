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
	Endpoints       []Endpoints
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
	country = s.Split(string(out), "Registrant Country:")[1]
	country = s.Split(country, "Registrant Phone:")[0]
	country = s.Replace(country, " ", "", -1)

	woner = s.Split(string(out), "Reseller:")[1]
	woner = s.Split(woner, "Domain Status:")[0]
	//woner = s.Replace(woner, " ", "", -1)

	return
}
func One(host string) {

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

	leng_servers := 2 //(data.Endpoints)
	out := &Out{}
	//servers := [2]OutServer{}
	//out.Servers = servers
	country, woner := getOwner(host)
	for i := 0; i < leng_servers; i++ {
		out.Servers[i] = OutServer{}
		out.Servers[i].Address = data.Endpoints[i].IpAddress
		out.Servers[i].Ssl_grade = data.Endpoints[i].Grade
		out.Servers[i].Country = country
		out.Servers[i].Owner = woner
	}
	fmt.Printf("%v", out)
}
