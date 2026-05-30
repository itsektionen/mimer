DATABASE_URL=postgres://postgres:postgres@localhost:5432/mimer?sslmode=disable

build:
	@echo "🛠️  Building..."
	@go build -ldflags "-s -w" -o ./dist/
	@echo "✅ Finished"

migrate-up:
	@migrate -path ./migrations -database "$(DATABASE_URL)" up

migrate-down:
	@migrate -path ./migrations -database "$(DATABASE_URL)" down 1

run:
	@echo "🧌 Running..."
	@go tool templ generate --watch --proxy="http://localhost:8080" --cmd="go run main.go"
