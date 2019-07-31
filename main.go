package main

import (
	_ "github.com/lib/pq"
)

func init() {

}

func main() {
	Rutes()
	/*
		http.HandleFunc("/api/analyce", func(w http.ResponseWriter, r *http.Request) {
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
	*/
}
