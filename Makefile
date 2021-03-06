BINARY="bin/dio"

all: build test

clean:
	rm -rf bin/*

build:
	go build -v -o $(BINARY) ./src

run:
	./bin/dio

test:
	go test -v ./src/...