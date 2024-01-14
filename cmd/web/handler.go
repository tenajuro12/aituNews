// handlers.go
package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
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

func showPosts(w http.ResponseWriter, r *http.Request) {
	err := templatsh
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/AituNews")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	vars := mux.Vars(r)
	res, err := db.Query(fmt.Sprintf("SELECT * FROM articles WHERE 'id' = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.For_who, &post.Full_text)
		if err != nil {
			panic(err)
		}
		showPost = post
	}
}
