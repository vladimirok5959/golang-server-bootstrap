default: help

help:
	@echo "make help  - help text"
	@echo "make test  - run tests"
	@echo "make build - build demo"
	@echo "make run   - build and run demo"
	@echo "make clean - remove builded file"

test:
	@go test -v ./...

build:
	@go build -o ./golang-server-bootstrap

run: build
	@./golang-server-bootstrap --color=always

clean:
	@rm -r ./golang-server-bootstrap
