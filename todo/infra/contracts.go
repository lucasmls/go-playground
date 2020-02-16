package infra

import (
	"database/sql"
)

//PostgresProvider ...
type PostgresProvider interface {
	Execute(query string, args ...interface{}) (*sql.Rows, error)
	ExecuteOne(query string, args ...interface{}) (*sql.Row, error)

	BeginTransaction() (PgTransaction, error)
	CommitTransaction(transaction PgTransaction) error

	ExecuteTransaction(transaction PgTransaction, query string, args ...interface{}) (sql.Result, error)
	ExecuteTransactionAndQuery(transaction PgTransaction, query string, args ...interface{}) (*sql.Rows, error)

	RollbackTransaction(transaction PgTransaction) error
}

// JwtProvider ...
type JwtProvider interface {
	GenerateJWT(userID string) (string, error)
}

// PgTransaction ...
type PgTransaction struct {
	Tx *sql.Tx
}
