# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=demo

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v	
test: 
	$(GOTEST) -v ./...
cover:
	$(GOTEST) ./... -coverprofile=c.out && go tool cover -html=c.out
bench:
	$(GOTEST) ./... -bench=.	
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f c.out
run:
	$(GOBUILD) -o $(BINARY_NAME) -v 
	./$(BINARY_NAME)
lint:
	golangci-lint run