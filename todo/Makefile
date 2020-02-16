include .env
export $(shell sed 's/=.*//' .env)

server:
	@ echo
	@ echo "Starting the server..."
	@ echo
	@ go run ./cmd/server/main.go