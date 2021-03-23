build:
	go build -v .

test:
	go test -v ./...

lint:
	go vet -v ./...
