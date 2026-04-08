SERVER_OUT = bin/server
SEED_OUT = bin/seed
SERVER_MAIN = ./cmd/server
SEED_MAIN = ./cmd/seed

.PHONY: all run-server run-seed dev build-server build-seed clean test tidy

all: build-server build-seed


run-server:
	go run $(SERVER_MAIN)


run-seed:
	go run $(SEED_MAIN)


dev:
	air


build-server:
	go build -o $(SERVER_OUT) $(SERVER_MAIN)


build-seed:
	go build -o $(SEED_OUT) $(SEED_MAIN)


tidy:
	go mod tidy

test:
	go test -v ./...

clean:
	rm -rf bin/ tmp/