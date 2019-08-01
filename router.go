package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

/*pagina resultados extensos*/
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Truora-Whois"))
}

/*Lista los dominios previamente consultados*/
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

	//fmt.Fprint(w,"API List\n")
}

/*Analice un dominio*/
func Analyce(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "url")
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
