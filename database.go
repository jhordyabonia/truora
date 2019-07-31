package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB = nil
var err error = nil

func Init() {

}
func Start() {
	db, err = sql.Open("postgres", "postgresql://jhordy@localhost:26257/truora?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS truora_whois (url varchar(300) PRIMARY KEY, data text)"); err != nil {
		log.Fatal(err)
	}
}
func GetAll() (outList []string, out []Out) {
	if db == nil {
		Start()
	}

	counter, err := db.Query("SELECT count(url) as size FROM truora_whois")
	if err != nil {
		log.Fatal(err)
	}
	var size int
	if counter.Next() {
		if err := counter.Scan(&size); err != nil {
			log.Fatal(err)
		}
	}
	defer counter.Close()

	rows, err := db.Query("SELECT url,data FROM truora_whois")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	outList = make([]string, size)
	out = make([]Out, size)
	for rows.Next() {
		var url, data string
		if err := rows.Scan(&url, &data); err != nil {
			log.Fatal(err)
		}
		size = size - 1
		out[size] = Out{}
		outList[size] = url
		json.Unmarshal([]byte(data), &out[size])
	}
	return
}

func Get(url string) (out Out, errOut bool) {
	if db == nil {
		Start()
	}
	rows, err := db.Query("SELECT url,data FROM truora_whois WHERE url LIKE '" + url + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	errOut = true
	if rows.Next() {
		var url, data string
		if err := rows.Scan(&url, &data); err != nil {
			log.Fatal(err)
		}
		errOut = data == ""
		if !errOut {
			json.Unmarshal([]byte(data), &out)
		}

	}
	return
}
func Insert(url string, data string) {
	if db == nil {
		Start()
	}
	if _, err := db.Exec(
		"INSERT INTO truora_whois (url, data) VALUES ('" + url + "', '" + data + "')"); err != nil {
		if _, err := db.Exec(
			"UPDATE truora_whois SET data = '" + data + "' WHERE url LIKE '" + url + "'"); err != nil {
			log.Fatal(err)
		}
	}
}
