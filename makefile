.PHONY: default run build test docs clean
APP_NAME=opportunities

# Tasks
default: run-with-docs 

run:
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.gon
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -r $(APP_NAME)
	@rm -rf ./docs