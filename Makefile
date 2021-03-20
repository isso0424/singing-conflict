auth/build:
	cd auth; \
	go build -v .

auth/test:
	cd auth; \
	go mod download; \
	go test -v ./...

auth/lint:
	cd auth; \
	go vet -v ./...

server/build:
	cd server; \
	go build -v .

server/test:
	cd server; \
	go mod download; \
	go test -v ./...

server/lint:
	cd server; \
	go vet -v ./...
