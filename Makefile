all: run

run: 
	go run main.go

docs: 
	~/go/bin/swag init

clean:
	go mod tidy