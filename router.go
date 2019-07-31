package main

/*
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
	list, _ := GetAll()
	out_json, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	w.Write([]byte(fmt.Sprintf("%v", string(out_json))))
}

func Analyce(w http.ResponseWriter, r *http.Request) {

	url := chi.URLParam(r, "url")
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
	//w.Write([]byte("API Analice\n"))
	//w.Write([]byte(fmt.Sprintf("Url %v", url)))
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
*/
