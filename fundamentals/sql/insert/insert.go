package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/crud_go")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	statement, _ := db.Prepare("INSERT INTO users(name) values (?)")
	result, _ := statement.Exec("Lucas")

	id, _ := result.LastInsertId()
	fmt.Println(id)

	affectedRows, _ := result.RowsAffected()
	fmt.Println(affectedRows)
}
