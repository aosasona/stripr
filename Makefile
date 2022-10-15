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
