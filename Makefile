BINARY := $(shell basename "$(PWD)")
SOURCES := ./

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.DEFAULT_GOAL := help

## build: Build the command line tool
build: clean
	go build -o ${BINARY} ${SOURCES}

## examples: Generates the Graph PNG examples
examples:
	rm -f ./_examples/*.png
	go run main.go -from '/***' -to '***/' ./_examples/Box.java | dot -Tpng > ./_examples/Box.java.png
	go run main.go ./_examples/docker-compose-sample.yml | dot -Tpng > ./_examples/docker-compose-sample.yml.png

## test: Starts unit test
test:
	go test -v ./... -coverprofile coverage.out

release:
	goreleaser --rm-dist --snapshot --skip-publish

## clean: Clean the binary
clean:
	rm -f $(BINARY)
