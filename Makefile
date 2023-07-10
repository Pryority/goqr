build:
	go build -o bin/gobank

run: build
	air	

test:
	go test -v ./...