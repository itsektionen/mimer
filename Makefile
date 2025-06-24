build:
	go build -ldflags "-s -w" -o ./dist/

run:
	DATABASE_URL=postgres://postgres:postgres@localhost:5432/mimer?sslmode=disable go run main.go
