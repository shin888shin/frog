package db

import (
	// "log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"strings"
)

type DbCreds struct {
	User     string
	Password string
	Database string
	Host     string
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

func Connection(conn *sql.DB) (*sql.DB, error) {
	creds, err := getDbCreds()
	creds_str := creds.User + ":" + creds.Password + "@/" + creds.Database

	conn, err = sql.Open("mysql", creds_str)
	if err != nil {
		panic(err)
	}
	return conn, nil
}
