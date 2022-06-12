package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type Db_credential struct {
	PSQL_USERNAME string
	PSQL_PASSWORD string
}

func getCredential() Db_credential {
	username := os.Getenv("PSQL_USERNAME")
	password := os.Getenv("PSQL_PASSWORD")
	return Db_credential{PSQL_USERNAME: username, PSQL_PASSWORD: password}
}

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

// func init() {
// 	var err error
// 	// cred := getCredential()
// 	// Initialize connection string.
// 	var connectionString string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "5432", "postgres", "admin", "chitchat")
// 	fmt.Println(connectionString)
// 	Db, err = sql.Open("postgres", connectionString)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from test limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("SELECT id, content, author FROM test where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	// statement := "insert into posts (content, author) values ($1, $2) returning id"
	statement := "INSERT INTO test (content, author) VALUES ($1, $2) RETURNING id;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("UPDATE test SET content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM test WHERE id = $1", post.Id)
	return
}

func main_psql_crud() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}

	fmt.Println(post)

	if err := post.Create(); err != nil {
		panic(err)
	}
	fmt.Println(post)

	readPost, err := GetPost(post.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"

	updateErr := readPost.Update()
	if updateErr != nil {
		panic(err)
	}

	posts, err := Posts(3)
	if err != nil {
		panic(err)
	}
	fmt.Println(posts)

	readPost.Delete()
}
