// cmd/web/main.go

package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"ui/html/home.html",
))

var templateCreate = template.Must(template.ParseFiles(
	"ui/html/create.html",
))

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", homeHandler).Methods("GET")
	rtr.HandleFunc("/create.html", createHandler).Methods("GET")
	rtr.HandleFunc("/post/{id:[0-9]+}", createHandler).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")

	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}
