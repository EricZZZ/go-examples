build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...      # <--- 这行前面必须是 Tab，不是空格

run: build
	@./bin/ecom            # <--- 这行前面必须是 Tab，不是空格

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
