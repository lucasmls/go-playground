package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// User ...
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UsersHandler ...
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	stringID := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(stringID)

	switch {
	case r.Method == "GET" && id > 0:
		findUser(w, r, id)
	case r.Method == "GET":
		listUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found...")
	}
}

func findUser(w http.ResponseWriter, r *http.Request, userID int) {
	db, err := sql.Open("mysql", "root:password@/crud_go")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var user User
	db.QueryRow("SELECT id, name from users WHERE id = ?", userID).Scan(&user.ID, &user.Name)

	userJSON, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(userJSON))
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@/crud_go")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, _ := db.Query("SELECT id, name from users")
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	jsonUsers, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonUsers))
}
