package userrepository

import (
	"fmt"
	"log"

	"github.com/lucasmls/todo/domain"
	"github.com/lucasmls/todo/infra"
	"golang.org/x/crypto/bcrypt"
)

// ClientInput ...
type ClientInput struct {
	Pg infra.PostgresProvider
}

// Client ...
type Client struct {
	in ClientInput
}

// NewClient ...
func NewClient(in ClientInput) *Client {
	if in.Pg == nil {
		log.Panic("Postgres provider is required")
	}

	return &Client{in: in}
}

// Find ...
func (c Client) Find() ([]*domain.User, error) {
	var params []interface{}

	rows, err := c.in.Pg.Execute("SELECT * FROM users", params...)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Failed to fetch users")
	}

	defer rows.Close()

	users := make([]*domain.User, 0)

	for rows.Next() {
		user := new(domain.User)

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Gender, &user.Phone, &user.Password)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("Failed to parse users")
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("A error occurred when searching for users")
	}

	return users, nil
}

// FindOne ...
func (c Client) FindOne(userID int64) (*domain.User, error) {
	row, err := c.in.Pg.ExecuteOne("SELECT * FROM users WHERE ID = $1", userID)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Failed to fetch users")
	}

	user := new(domain.User)
	scanErr := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Gender, &user.Phone, &user.Password)

	if scanErr != nil {
		fmt.Println(err)
		return nil, scanErr
	}

	return user, nil
}

// FindByEmail ...
func (c Client) FindByEmail(email string) (*domain.User, error) {
	row, err := c.in.Pg.ExecuteOne("SELECT * FROM users WHERE EMAIL = $1", email)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Failed to fetch users")
	}

	user := new(domain.User)
	scanErr := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Gender, &user.Phone, &user.Password)

	if scanErr != nil {
		fmt.Println(err)
		return nil, scanErr
	}

	return user, nil
}

// Save ...
func (c Client) Save(userDto domain.User) (*domain.User, error) {
	transaction, trErr := c.in.Pg.BeginTransaction()
	if trErr != nil {
		return nil, fmt.Errorf("Failed to start the transaction")
	}

	passwordHash, hashError := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if hashError != nil {
		return nil, fmt.Errorf("Failed to bcrypt the user password")
	}

	result, execTxErr := c.in.Pg.ExecuteTransactionAndQuery(
		transaction,
		"INSERT INTO users(name, email, password, age, gender, phone) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		userDto.Name, userDto.Email, passwordHash, userDto.Age, userDto.Gender, userDto.Phone,
	)

	if execTxErr != nil {
		return nil, fmt.Errorf("Failed to execute the transaction")
	}

	result.Next()
	var userID int64

	lastIDErr := result.Scan(&userID)
	if lastIDErr != nil {
		log.Panic(lastIDErr)
		return nil, fmt.Errorf("Failed to get the lastest inserted id")
	}

	c.in.Pg.CommitTransaction(transaction)

	user, userErr := c.FindOne(userID)
	if userErr != nil {
		log.Panic(userErr)
		return nil, fmt.Errorf("Failed to find the latest user registered")
	}

	return user, nil
}
