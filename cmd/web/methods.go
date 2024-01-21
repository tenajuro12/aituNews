package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

type Article struct {
	Id            uint16
	Title         string
	Anons         string
	Full_text     string
	For_who       string
	CreatedAt     string
	CreatedAtTime time.Time
}

var posts = []Article{}

func save_article(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/AituNews")
	if err != nil {
		panic(err)
	}

	title := r.FormValue("title")
	anons := r.FormValue("anons")
	fullText := r.FormValue("full_text")
	forWho := r.FormValue("for_who")

	insert, err := db.Prepare("INSERT INTO articles (title, anons, full_text, for_who, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = insert.Exec(title, anons, fullText, forWho, currentTime)
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
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
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.For_who, &post.Full_text, &post.CreatedAt)
		if err != nil {
			panic(err)
		}

		post.CreatedAtTime, err = time.Parse("2006-01-02 15:04:05", post.CreatedAt)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

}

func filterArticles(forWho string) []Article {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/AituNews")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM articles WHERE for_who = ?", forWho)
	if err != nil {
		panic(err)
	}

	var filteredPosts []Article

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.For_who, &post.Full_text, &post.CreatedAt)
		if err != nil {
			panic(err)
		}

		post.CreatedAtTime, err = time.Parse("2006-01-02 15:04:05", post.CreatedAt)
		if err != nil {
			panic(err)
		}

		filteredPosts = append(filteredPosts, post)
	}

	return filteredPosts
}

func save_contact(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/AituNews")
	if err != nil {
		panic(err)
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	comment := r.FormValue("comment")

	insert, err := db.Prepare("INSERT INTO contacts (name, email, comment) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}

	_, err = insert.Exec(name, email, comment)
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
