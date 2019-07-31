package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

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

	//Rutes()
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("url")
		_data, err := Get(url)
		var data string
		if err {
			tmp, err := json.Marshal(_data)
			if err != nil {
				data = string(tmp)
			}
		} else {
			data = One(url)
		}
		Insert(url, data)
		fmt.Fprint(w, data)
	})
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		list, _ := GetAll()
		type items struct {
			Items []string
		}

		out_json, err := json.Marshal(items{Items: list})
		if err != nil {
			fmt.Println(err)
		}

		w.Write([]byte(fmt.Sprintf("%v", string(out_json))))
	})
	http.ListenAndServe(":8090", nil)

}
