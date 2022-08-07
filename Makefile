FILE := trustme
.DEFAULT_GOAL:=help

.PHONY: fmt # run go fmt command to whole files
fmt:
	go fmt ./...

# for make run with arguments
ifeq (run,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

.PHONY: run # test run 
run:
	go run main.go $(RUN_ARGS)

.PHONY: build # build binary file to ./bin
build: fmt
	go build -ldflags="-s -w" -o bin/${FILE} main.go

.PHONY: build-arm64 # build binary file to ./bin for linux/arm64
build-arm64: fmt
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o bin/${FILE}_arm64 main.go

.PHONY: build-amd64 # build binary file to ./bin for linux/amd64
build-amd64: fmt
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/${FILE}_amd64 main.go

.PHONY: clean # remove binary build file from ./bin
clean:
	rm -rf ./bin/${FILE}

.PHONY: help # print whole command of Makefile
 help:
	@grep -E '^\.PHONY: .+ .*?# .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = "(: |#)"}; {printf "%-30s %s\n", $$2, $$3}'

.PHONY: test # run unit tests
test: 
	go test -race -covermode=atomic ./...
