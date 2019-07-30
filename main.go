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

func init() {

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

	Test()
	//One("truora.com")
	//getOwner("truora.com")
	panic(0)

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		message := API{"Hello, world!"}
		output, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Algo salió mal")
		}
		fmt.Println("hering...")
		fmt.Fprintln(w, string(output))
	})
	http.ListenAndServe(":8090", nil)
}
