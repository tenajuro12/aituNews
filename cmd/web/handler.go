// handlers.go
package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	allArticles()
	err := templates.ExecuteTemplate(w, "home.html", posts)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func createHandler(w http.ResponseWriter, r *http.Request) {
	err := templateCreate.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	err := templateContact.ExecuteTemplate(w, "contact.html", nil)
	if err != nil {
		log.Println(err) // This will print the error message in your console
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	forWho := params["for_who"]
	filteredArticles := filterArticles(forWho)
	err := templateFilter.ExecuteTemplate(w, "filtered.html", filteredArticles)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
