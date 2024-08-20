.PHONY: default run build clean
APP_NAME = "gowallet"

default: run

run:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test -v ./...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -f ./docs/swagger.json