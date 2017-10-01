package main

import (
	// "log"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shin888shin/frog/db"
	"html/template"
	"net/http"
	"strconv"
)

var err error

// var dbConfig db.Database
var dbConn *sql.DB

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/todos", handleAllTodos)
	http.ListenAndServe(":8080", nil)
}

var t *template.Template

func init() {
	fmt.Println("Initialize")

	dbConn, err = db.Connection(dbConn)
	fmt.Println("---> init dbConn:", dbConn)

	if err != nil {
		panic(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("---> handleRoot dbConn:", dbConn)

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

type Todo struct {
	Id          int
	Name        string
	Description string
}

func handleAllTodos(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("---> handleAllTodos dbConn:", dbConn)
	// todos := []Todo{}

	todos := showAllTodos()
	for _, t := range todos {
		fmt.Println("ooo> ", t)
	}
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
	rows, err := dbConn.Query("SELECT * FROM todos")
	if err != nil {
		panic(err.Error())
	}

	var count int
	for rows.Next() {
		count += 1
	}
	fmt.Println("+++> count:", count)

	rows, err = dbConn.Query("SELECT * FROM todos")
	if err != nil {
		panic(err.Error())
	}

	// fmt.Printf("---> rows: %T - %v\n", rows, rows)

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// fmt.Printf("---> columns: %T - %v\n", columns, columns)
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// fmt.Println("+++> len of values:", len(rows))
	fmt.Println("+++> count 2:", count)
	var todos = make([]Todo, 1)

	for rows.Next() {
		fmt.Println("+++> row")
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		var value string
		var todo Todo
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				// fmt.Printf("---> col: %T - %v\n", col, col)
				if i == 0 {
					value = string(col)
					x, _ := strconv.Atoi(value)
					todo.Id = x
				} else if i == 1 {
					value = string(col)
					todo.Name = value
				} else {
					value = string(col)
					todo.Description = value
				}
			}
		}
		fmt.Println("todo:", todo)
		todos = append(todos, todo)
		// fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return todos
}
