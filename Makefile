GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVENDOR=$(GOCMD) mod vendor

install:
	$(GOVENDOR)

update:
	$(GOGET) -u all
	$(GOCMD) mod tidy
