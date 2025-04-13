.DEFAULT_GOAL := help

.PHONY: build
build: ## build project
	tinygo flash --target=arduino main.go

.PHONY: help
help: ## print this help and exit
	@echo "Usage: make [target]"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
