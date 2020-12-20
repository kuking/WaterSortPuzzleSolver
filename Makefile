all: clean test build run

clean:
	go clean -testcache -cache
	rm -f WSPZ

build:
	go build -o WSPZ

test:
	go test

run:
	./WSPZ
