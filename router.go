package main

import (
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

	w.Write([]byte("API list"))
}

func Analyce(w http.ResponseWriter, r *http.Request) {

	url := chi.URLParam(r, "url")
	//w.Write([]byte("API Analice\n"))
	//w.Write([]byte(fmt.Sprintf("Url %v", url)))
	fmt.Fprint(w, One(url))
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
