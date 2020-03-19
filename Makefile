all: run

run: 
	go run main.go

docs: 
	~/$HOME/go/bin/swag init