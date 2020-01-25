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

	stmt, _ := db.Prepare("UPDATE users set name = ? where id = ?")
	stmt.Exec("Laislinha", 1)
}
