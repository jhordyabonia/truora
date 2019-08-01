package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*Metodo sin uso, alternartiva a la implementacion chi*/
func start() {
	http.HandleFunc("/api/analyce", func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("url")
		out := ApiAnalyce(url)
		old, err0 := Get(url)
		if !err0 {
			out = Compare(out, old)
		}
		tmp, err := json.Marshal(&out)
		if err != nil {
			return
		}
		data := string(tmp)
		Insert(url, data)
		fmt.Fprint(w, data)
	})
	http.HandleFunc("/api/list", func(w http.ResponseWriter, r *http.Request) {
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

func main() {
	Rutes()
}
