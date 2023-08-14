install:
	go install github.com/spf13/cobra-cli@latest

create:
	go run main.go category create --name=test

list:
	go run main.go category list