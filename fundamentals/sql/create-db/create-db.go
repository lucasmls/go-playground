package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func execSQL(db *sql.DB, query string) sql.Result {
	result, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	db, err := sql.Open("mysql", "root:password@/")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	execSQL(db, "CREATE DATABASE IF NOT EXISTS crud_go")
	execSQL(db, "USE crud_go")
	execSQL(db, `
		CREATE TABLE IF NOT EXISTS users (
			id integer auto_increment,
			name varchar(80),
			PRIMARY KEY (id)
		)
	`)
}
