package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // ...
	"github.com/lucasmls/todo/domain"
	"log"
)

// Client ...
type Client struct {
	db *sql.DB
}

// ClientInput ...
type ClientInput struct {
	Endpoint string
}

// NewClient ...
func NewClient(in ClientInput) *Client {
	if in.Endpoint == "" {
		log.Panic("Endpoint is required")
	}

	db, err := sql.Open("postgres", in.Endpoint)
	if err != nil {
		log.Panic(err)
	}

	if pingError := db.Ping(); pingError != nil {
		log.Panic(pingError)
	}

	return &Client{db: db}
}

// GetUsers ...
func (c Client) GetUsers() ([]*domain.User, string) {
	fmt.Println("@GetUsers")

	rows, err := c.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, "Failed to fetch users"
	}

	defer rows.Close()

	users := make([]*domain.User, 0)

	for rows.Next() {
		user := new(domain.User)

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Gender, &user.Phone)
		if err != nil {
			fmt.Println(err)
			return nil, "Failed to parse users"
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, "Error"
	}

	return users, ""
}
