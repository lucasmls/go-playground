package infra

import "github.com/lucasmls/todo/domain"

//PostgresProvider ...
type PostgresProvider interface {
	GetUsers() ([]*domain.User, string)
	SaveUser(domain.User) (string, string)
}
