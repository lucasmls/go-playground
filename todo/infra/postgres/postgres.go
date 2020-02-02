package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // ...
	"github.com/lucasmls/todo/domain"
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
	rows, err := c.db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
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

// SaveUser ...
func (c Client) SaveUser(userDto domain.User) (string, string) {
	query := fmt.Sprintf(`
		INSERT INTO users (name, email, age, gender, phone) VALUES ('%s', '%s', %d, '%s', '%s')`,
		userDto.Name, userDto.Email, userDto.Age, userDto.Gender, userDto.Phone,
	)

	result, err := c.db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return "", "Failed to save user into db"
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return "", "Failed to count how much rows were affected"
	}

	successMessage := fmt.Sprintf(`User %s registered successfully (%d row affected)`, userDto.Name, rowsAffected)
	return successMessage, ""
}
