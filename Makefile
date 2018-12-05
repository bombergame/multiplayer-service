all: build

generate:
	go generate ./...

build:
	go build -v -o service .

clean:
	rm -rf ./service
	rm -rf ./coverage.out
