package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Truora-Whois"))
}
func List(w http.ResponseWriter, r *http.Request) {
	list, _ := GetAll()
	type items struct {
		Items []string
	}

	out_json, err := json.Marshal(items{Items: list})
	if err != nil {
		fmt.Println(err)
	}
	w.Write([]byte(fmt.Sprintf("%v", string(out_json))))
}
func compare(in1, in2 Out) (out Out) {
	out = in1
	out.Servers_changed = false
	for i := 0; i < len(out.Servers); i++ {
		if in1.Servers[i].Address != in2.Servers[i].Address {
			out.Servers_changed = true
		}
		if in1.Servers[i].Ssl_grade != in2.Servers[i].Ssl_grade {
			out.Servers_changed = true
		}
		if in1.Servers[i].Country != in2.Servers[i].Country {
			out.Servers_changed = true
		}
		if in1.Servers[i].Owner != in2.Servers[i].Owner {
			out.Servers_changed = true
		}
	}
	if in1.Ssl_grade != in2.Ssl_grade {
		out.Ssl_grade = in2.Ssl_grade
		out.Servers_changed = true
	}
	out.Previous_ssl_grade = in2.Ssl_grade
	return
}
func Analyce(w http.ResponseWriter, r *http.Request) {

	url := chi.URLParam(r, "url")
	out := One(url)
	old, err0 := Get(url)
	if !err0 {
		out = compare(out, old)
	}
	tmp, err := json.Marshal(&out)
	if err != nil {
		return
	}
	data := string(tmp)
	Insert(url, data)
	fmt.Fprint(w, data)
	//fmt.Fprint(w,"API Analice\n")
	//fmt.Fprint(w, "Url ", url)
}
func Rutes() {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/", home)
		r.With(paginate).Get("/list", List)
		r.Get("/analyce/{url}", Analyce)
	})
	http.ListenAndServe(":8090", r)
}
