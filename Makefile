build:
	@echo "🛠️  Building..."
	@go build -ldflags "-s -w" -o ./dist/
	@echo "✅ Finished"

migrate-up:
	@migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/mimer?sslmode=disable" up

migrate-down:
	@migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/mimer?sslmode=disable" down 1

run:
	@echo "🧌 Running..."
	DATABASE_URL=postgres://postgres:postgres@localhost:5432/mimer?sslmode=disable go run main.go
