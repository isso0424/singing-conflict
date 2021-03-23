build:
	go build -v .

test:
	go test -v ./...

server/lint:
	go vet -v ./...
