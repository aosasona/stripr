.PHONY: all
all:
	@echo "make build"
	@echo "make clean"

build:
	go build -o stripr .
	make move

move:
	test -d ./bin || mkdir ./bin
	mv ./stripr ./bin/stripr

clean:
	rm -rf ./bin

# for development
scan-single:
	go run . -target=./example/index.js scan

scan:
	go run . -target=./example scan
