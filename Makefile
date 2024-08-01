run:
	go run cmd/app/main.go

swag:
	swag init -g cmd/app/main.go -o docs
