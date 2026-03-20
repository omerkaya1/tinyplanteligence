.DEFAULT_GOAL := help

.PHONY: trial
trial: ## build trial programme (may contain different peripherals for testing purposes)
	@tinygo \
	flash --target=arduino ./cmd/trial/main.go

.PHONY: moisture-sensor
moisture-sensor: ## build moisture sensor programme
	@tinygo \
	flash --target=arduino ./cmd/moisture-sensor/main.go

.PHONY: help
help: ## print this help and exit
	@echo "Usage: make [target]"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
