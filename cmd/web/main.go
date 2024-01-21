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

var templateContact = template.Must(template.ParseFiles(
	"ui/html/contact.html",
))
var templateFilter = template.Must(template.ParseFiles(
	"ui/html/filtered.html",
))

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", homeHandler).Methods("GET")
	rtr.HandleFunc("/create.html", createHandler).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")
	rtr.HandleFunc("/contact.html", contactHandler).Methods("GET")
	rtr.HandleFunc("/filtered/{for_who}", filterHandler).Methods("GET")
	rtr.HandleFunc("/save_contact", save_contact).Methods("POST")

	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}
