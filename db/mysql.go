package db

import (
  "fmt"
  // "log"
  "io/ioutil"
  "strings"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"

)

type DbCreds struct {
  User string
  Password string
  Database string
  Host string
}
func getDbCreds() (DbCreds, error) {
  var creds DbCreds
  fileName := "config/local_mysql.txt"
  contentBytes, err := ioutil.ReadFile(fileName)

  if err != nil {
    return creds, err
  }
  result := string(contentBytes)
  lines := strings.Split(result, "\n")

  for i, line := range lines {
    val := strings.TrimSpace(line)
    if i == 0 {
      creds.User = val
    } else if i == 1 {
      creds.Password = val
    } else if i == 2 {
      creds.Database = val
    } else {
      creds.Host = val
    }
  }
  return creds, nil

}
func Zzz() (*sql.DB, error) {


  creds, err := getDbCreds()

  creds_str := creds.User + ":" + creds.Password + "@/" + creds.Database
  fmt.Println("Creds:", creds_str)

  db, err := sql.Open("mysql", creds_str)
  // defer db.Close()
  fmt.Printf("Initializezzzxx mysql %T - %v\n", db, db)

  // dbConfig, err = db.DatabaseSetup()
  if err != nil {
    panic(err)
  }
  return db, nil
}

// func DatabaseSetup() (Database, error) {
//   var db Database
//   fileName := "config/local_mysql.txt"
//   contentBytes, err := ioutil.ReadFile(fileName)

//   if err != nil {
//     return db, err
//   }

//   result := string(contentBytes)
//   lines := strings.Split(result, "\n")

//   for i, line := range lines {
//     fmt.Println("---> line:", line)
//     val := strings.TrimSpace(line)
//     if i == 0 {
//       db.User = val
//     } else if i == 1 {
//       db.Password = val
//     } else if i == 2 {
//       db.Database = val
//     } else {
//       db.Host = val
//     }
//     fmt.Printf("---> db: %v\n", db)

//   }
//   return db, nil
// }