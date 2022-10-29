.PHONY: all
all:
	@echo "make build"
	@echo "make clean"


.PHONY: build
build:
	@echo "make build-mac"
	@echo "make build-linux"
	@echo "make build-windows"
	@echo "make build-all"
	@echo "make build-auto"
	@echo "make build-release"


.PHONY: build-unix
build-unix:
	GOOS=${TARGET} GOARCH=${ARCH} go build -o stripr .
	make move DIR=${dir}

.PHONY: build-mac
build-mac:
	@echo ">> Building for Mac AMD64"
	make build-unix TARGET=darwin ARCH=amd64
	@echo ">> Building for Mac ARM"
	make build-unix TARGET=darwin ARCH=arm64

.PHONY: build-linux
build-linux:
	@echo ">> Building for Linux AMD64"
	make build-unix TARGET=linux ARCH=amd64
	@echo ">> Building for Linux ARM"
	make build-unix TARGET=linux ARCH=arm64

.PHONY: build-windows
build-windows:
	@echo ">> Building for Windows AMD64"
	GOOS=windows GOARCH=amd64 go build -o stripr.exe .
	make move DIR=build/windows

.PHONY: build-all
build-all:
	make build-mac
	make build-linux
	make build-windows

platform := $(shell uname -s)
arch := $(shell uname -m)
dir := ${TARGET}-${ARCH}
MODE ?= dev
mode := ${MODE}
ifeq (${mode}, dev)
	dir := build/${dir}
else
	dir := release
endif

.PHONY: build-auto
build-auto:
ifeq ($(platform),Darwin)
	make build-unix TARGET=darwin ARCH=${arch};
else ifeq ($(platform),Linux)
	make build-unix TARGET=linux ARCH=${arch};
else ifeq ($(platform),Windows)
	make build-windows;
endif

.PHONY: build-release
build-release:
	@echo ">> Building release version for ${platform} ${arch}"
	make build-auto MODE=release

DIR=bin #default directory
move:
	@echo ">> Moving stripr to ${DIR}"
	test -d ./${DIR} || mkdir -p ./${DIR}
	if [ -f ./stripr.exe ]; then \
  		mv ./stripr.exe ./${DIR}/stripr.exe; \
  	elif [ -f ./stripr ]; then \
  	  	mv ./stripr ./${DIR}/stripr; \
  	fi;

clean:
	@echo ">> Cleaning up"
	rm -rf ./bin ./build ./release ./stripr ./stripr.exe




# for development
scan-single:
	go run . -target=./example/index.js scan

scan:
	go run . -target=./example scan
