build:
	@go build -o bin/kc-test-tech main.go

run: build
	@./bin/kc-test-tech api

migrate-up:
	@go run main.go db migrate

migrate-down:
	@go run main.go db down