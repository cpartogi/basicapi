SHELL := /bin/bash

## Golang Stuff
GOCMD=go
GORUN=$(GOCMD) run

SERVICE=basicapi

init:
	$(GOCMD) mod init $(SERVICE)

test:
	$(GOCMD) test -cover ./...

run:
	$(GORUN) main.go	