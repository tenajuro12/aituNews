package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

type Article struct {
	Id         uint16
	Title      string
	Anons      string
	Full_text  string
	For_who    string
	Categories []string
}

var posts = []Article{}
var showPost = Article{}

func save_article(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/AituNews")
	if err != nil {
		panic(err)
	}

	title := r.FormValue("title")
	anons := r.FormValue("anons")
	fullText := r.FormValue("full_text")
	forWho := r.FormValue("for_who")
	categories := r.Form["categories"]

	_, err = db.Exec("INSERT INTO articles (title, anons, full_text, for_who) VALUES (?, ?, ?, ?)", title, anons, fullText, forWho)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get the ID of the inserted article
	var articleID int
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&articleID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Associate categories with the article
	for _, category := range categories {
		categoryID, err := getCategoryID(category)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO article_categories (article_id, category_id) VALUES (?, ?)", articleID, categoryID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getCategoryID(categoryName string) (int, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/AituNews")
	if err != nil {
		panic(err)
	}
	var categoryID int
	err = db.QueryRow("SELECT id FROM categories WHERE name = ?", categoryName).Scan(&categoryID)
	return categoryID, err
}

func allArticles() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/AituNews")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM articles")
	if err != nil {
		panic(err)
	}

	posts = []Article{}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.For_who, &post.Full_text)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
}

func ShowPosts(r *http.Request) {
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
