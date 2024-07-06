debug:
	air

run:
	go run cmd/main.go

tidy:
	go mod tidy

build-app:
	go build -o ./build/main cmd/main.go

docs:
	swag fmt ./
	swag init -o ./api -g init.go -d ./internal/app,./internal/controller,./internal/http/response,./internal/http/request,./internal/http/server
