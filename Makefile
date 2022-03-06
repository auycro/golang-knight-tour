GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=knighttour

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)

all: help

## Build:
build:
	mkdir -p out/bin
	## $(GOCMD) build -o out/bin/$(BINARY_NAME) cmd/main.go
	GO111MODULE=on $(GOCMD) build -mod vendor -o out/bin/$(BINARY_NAME) cmd/main.go

run:
	$(GOCMD) run cmd/main.go

clean:
	rm -rf out

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)