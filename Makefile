GOTEST=go test
BINARY_NAME=parking

.PHONY: test build run 

build: ## Build project and put the output binary in out/bin/
	mkdir -p out/bin
	GO111MODULE=on go build -o out/bin/$(BINARY_NAME) .

run: build
	./out/bin/$(BINARY_NAME)

clean: ## Remove build related file
	rm -fr ./bin
	rm -fr ./out

## Test:
test: 
	$(GOTEST) -cover -v -race ./... 
