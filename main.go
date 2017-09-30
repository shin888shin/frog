package main

import (
	// "log"
	"fmt"
	"html/template"
	"net/http"
	"github.com/shin888shin/frog/db"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// var err error
// var dbConfig db.Database
var dbConn *sql.DB

func main() {
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":8080", nil)
}

var t *template.Template

func init() {
	fmt.Println("Initialize frog")

 //  db, err := sql.Open("mysql", "root:yllis@/fish_dev")
 //  defer db.Close()
	// fmt.Printf("Initialize mysql %T - %v\n", db, db)

	dbConn, err := db.Zzz()
  if err != nil {
    panic(err)
  }

  rows, err := dbConn.Query("SELECT * FROM todos")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Printf("---> rows: %T - %v\n", rows, rows)

  columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("---> columns: %T - %v\n", columns, columns)
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

  for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
 //  db, err = mysql.DialPassword(dbConfig.Host, dbConfig.User, dbConfig.Password)  
 //  if err != nil {
 //    log.Print("oh no")
 //  }
 //  _, err4 := db.Exec("use fish_dev")
 //  if err4 != nil {
 //    log.Fatal(err)
 //  }
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	t, _ = template.ParseFiles("views/home.html", "views/layout.html")
	err := t.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
