DATABASE_URL=postgres://postgres:postgres@localhost:5432/mimer?sslmode=disable

build:
	@echo "🛠️  Building..."
	@go build -ldflags "-s -w" -o ./dist/
	@echo "✅ Finished"

migrate-up:
	cd db/migrations && goose postgres "$(DATABASE_URL)" up
migrate-down:
	cd db/migrations && goose postgres "$(DATABASE_URL)" down

tw:
	tailwindcss -i ./tailwind.css -o ./static/index.css --watch

run:
	@go tool templ generate --watch --proxy="http://localhost:8888" --cmd="go run cmd/serve/serve.go"

dev:
	@make -j 2 tw run
