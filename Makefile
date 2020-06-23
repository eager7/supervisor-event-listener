#/bin/bash
# This is how we want to name the binary output
TARGET_SERVER=listener

SERVER_SRC=main.go
# These are the values we want to pass for Version and BuildTime
GITTAG=1.0.0
BUILD_TIME=`date +%Y%m%d%H%M%S`
# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${GITTAG} -X main.Build_Time=${BUILD_TIME} -s -w"

default: mod

release: server

server:
	GO111MODULE=on CGO_ENABLED=1 go build -v -mod=vendor ${LDFLAGS} -o build/${TARGET_SERVER} ${SERVER_SRC}

mod:
	export GOPROXY=https://goproxy.io &&  GO111MODULE=on go build -v ${LDFLAGS} -o build/${TARGET_SERVER} ${SERVER_SRC}

linux:
	export GOPROXY=https://goproxy.io &&  GO111MODULE=on GOOS=linux go build -v ${LDFLAGS} -o build/${TARGET_SERVER} ${SERVER_SRC}

normal:
	GO111MODULE=on go build -v ${LDFLAGS} -o build/${TARGET_SERVER} ${SERVER_SRC}

local:
	git config --global url."git@github.com:".insteadOf "https://github.com/" && GO111MODULE=on go build -v ${LDFLAGS} -o build/${TARGET_SERVER} ${SERVER_SRC}

depends:
	GO111MODULE=on go mod download

tidy:
	GO111MODULE=on go mod tidy

clean:
	-rm -rf build

check:
	GO111MODULE=on golangci-lint run ./...
