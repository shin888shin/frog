package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shin888shin/frog/db"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id          int
	Name        string
	Description string
}

var err error
var dbConn *sql.DB
var t *template.Template

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/todos", handleAllTodos)
	http.ListenAndServe(":8080", nil)
}

func init() {
	dbConn, err = db.Connection(dbConn)
	if err != nil {
		panic(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	t, _ = template.ParseFiles("views/home.html", "views/layout.html")
	err = t.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := showAllTodos()
	if r.URL.Path != "/todos" {
		http.NotFound(w, r)
		return
	}
	t, _ = template.ParseFiles("views/todos/index.html", "views/layout.html")
	err := t.ExecuteTemplate(w, "layout", struct{ Todos []Todo }{todos})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func showAllTodos() []Todo {
	var (
		id          int
		name        string
		description string
	)
	rows, err := dbConn.Query("SELECT id, name, description FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todos = make([]Todo, 0)
	for rows.Next() {
		err := rows.Scan(&id, &name, &description)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, Todo{Id: id, Name: name, Description: description})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return todos
}
