run: build
	@./bin/goth_stack

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

build:
	tailwindcss -i view/css/app.css -o public/styles.css
	@templ generate view
	@go build -o bin/goth_stack main.go

up:
	@go run cmd/migrate/main.go up

drop:
	@go run cmd/drop/main.go up

down:
	@go run cmd/migrate/main.go

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@, $(MAKECMDGOALS))

seed:
	@go run cmd/seed/main.go
