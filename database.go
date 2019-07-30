package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func init() {

}
func Start() {

}
func Test() {
	// Connect to the "bank" database.
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/truora?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	// Create the "truora_whois" table.
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS truora_whois (url varchar(300) PRIMARY KEY, data text)"); err != nil {
		log.Fatal(err)
	}

	// Insert two rows into the "accounts" table.
	if _, err := db.Exec(
		"INSERT INTO truora_whois (url, data) VALUES ('facebook.com', 'empty')"); err != nil {
		log.Fatal(err)
	}

	// Print out the balances.
	rows, err := db.Query("SELECT url,data FROM truora_whois")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("datas:")
	for rows.Next() {
		var url, data string
		if err := rows.Scan(&url, &data); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n %s\n", url, data)
	}
}
