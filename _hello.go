package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type API struct {
	Message string "json:message"
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		message := API{"Hello, world!"}
		output, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Algo salió mal")
		}
		fmt.Println("hering...")
		fmt.Println(string(output))
	})
	http.ListenAndServe(":8090", nil)
	/*
	   // Connect to the "bank" database.
	   db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/bank?sslmode=disable")
	   if err != nil {
	       log.Fatal("error connecting to the database: ", err)
	   }

	   // Create the "accounts" table.
	   if _, err := db.Exec(
	       "CREATE TABLE IF NOT EXISTS accounts (id INT PRIMARY KEY, balance INT)"); err != nil {
	       log.Fatal(err)
	   }

	   // Insert two rows into the "accounts" table.
	   if _, err := db.Exec(
	       "INSERT INTO accounts (id, balance) VALUES (1, 1000), (2, 250)"); err != nil {
	       log.Fatal(err)
	   }

	   // Print out the balances.
	   rows, err := db.Query("SELECT id, balance FROM accounts")
	   if err != nil {
	       log.Fatal(err)
	   }
	   defer rows.Close()
	   fmt.Println("Initial balances:")
	   for rows.Next() {
	       var id, balance int
	       if err := rows.Scan(&id, &balance); err != nil {
	           log.Fatal(err)
	       }
	       fmt.Printf("%d %d\n", id, balance)
	   }*/
}

/*package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Endpoints struct {
	ipAddress         string //`json:"ipAddress"`
	ServerName        string
	statusMessage     string
	grade             string
	gradeTrustIgnored string
	hasWarnings       string
	isExceptional     string
	progress          string
	duration          string
	delegation        string
}
type Result struct {
	Host            string
	port            string
	protocol        string
	isPublic        string
	status          string
	startTime       string
	testTime        string
	engineVersion   string
	CriteriaVersion string
	Endpoints       [2]Endpoints
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	// file, err := os.Open(var text.txtvar )

	// var b = make([]byte, 500)
	// var read int
	// read, err = file.Read(b)
	// check(err)
	// file.Close()

	// fmt.Printf(var hello, world\n %v\nvar , string(b[:read]))
	// var x int
	// for x = 0; x < 10; x++ {
	// 	if x%2 == 0 {
	// 		fmt.Printf(var %v  Test \nvar , x)
	// 	}
	// }
	url := "https://api.ssllabs.com/api/v3/analyze?host=truora.com"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{" title string //" Buy cheese and bread for breakfast." }`)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
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
}
*/
