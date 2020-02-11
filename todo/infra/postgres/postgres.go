package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // ...
	"github.com/lucasmls/todo/infra"
)

// Client ...
type Client struct {
	db *sql.DB
	Db *sql.DB
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

	return &Client{db: db, Db: db}
}

// Execute ...
func (c Client) Execute(query string, args ...interface{}) (*sql.Rows, error) {
	return c.db.Query(query, args...)
}

// ExecuteOne ...
func (c Client) ExecuteOne(query string, args ...interface{}) (*sql.Row, error) {
	return c.db.QueryRow(query, args...), nil
}

// BeginTransaction ...
func (c Client) BeginTransaction() (infra.PgTransaction, error) {
	transaction, err := c.db.Begin()
	return infra.PgTransaction{Tx: transaction}, err
}

// ExecuteTransaction ...
func (c Client) ExecuteTransaction(transaction infra.PgTransaction, query string, args ...interface{}) (sql.Result, error) {
	if transaction.Tx == nil {
		return nil, fmt.Errorf("The transaction is required to be executed")
	}

	result, err := transaction.Tx.Exec(query, args...)

	return result, err
}

//ExecuteTransactionAndQuery ...
func (c Client) ExecuteTransactionAndQuery(transaction infra.PgTransaction, query string, args ...interface{}) (*sql.Rows, error) {
	if transaction.Tx == nil {
		return nil, fmt.Errorf("The transaction is required to be executed")
	}

	result, err := transaction.Tx.Query(query, args...)

	return result, err
}

// CommitTransaction ...
func (c Client) CommitTransaction(transaction infra.PgTransaction) error {
	if transaction.Tx == nil {
		return fmt.Errorf("The transaction is required to be commited")
	}

	return transaction.Tx.Commit()
}

// RollbackTransaction ...
func (c Client) RollbackTransaction(transaction infra.PgTransaction) error {
	if transaction.Tx == nil {
		return fmt.Errorf("The transaction is required to execute the rollback")
	}

	transaction.Tx.Rollback()

	return nil
}
