package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/crud_go")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	transaction, _ := db.Begin()
	statement, _ := transaction.Prepare("INSERT INTO users(id, name) values(?, ?)")

	statement.Exec(2000, "Laisla")
	statement.Exec(2001, "Daniel")

	// _, err = statement.Exec(1, "Luis") // Duplicated id

	if err != nil {
		transaction.Rollback()
		log.Fatal(err)
	}

	transaction.Commit()
}
