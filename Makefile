.DEFAULT_GOAL := help

TARGET := arduino

.PHONY: trial
trial: ## build trial programme (may contain different peripherals for testing purposes)
	@tinygo \
	flash --target=$(TARGET) ./cmd/trial/main.go

.PHONY: moisture-sensor
moisture-sensor: ## build moisture sensor programme
	@tinygo \
	flash --target=$(TARGET) ./cmd/moisture-sensor/main.go

.PHONY: pump
pump: ## build pump programme
	@tinygo \
	flash --target=$(TARGET) ./cmd/pump/main.go

.PHONY: keyboard
keyboard: ## build keyboard programme
	@tinygo \
	flash --target=$(TARGET) ./cmd/keyboard/main.go && tinygo monitor --target=$(TARGET) --baudrate=9600

.PHONY: servo
servo: ## build servo programme
	@tinygo \
	flash --target=$(TARGET) ./cmd/servo/main.go && tinygo monitor --target=$(TARGET) --baudrate=9600

.PHONY: ultrasound-sensor
ultrasound-sensor: ## build ultrasound-sensor programme
	@tinygo \
	flash --target=$(TARGET) ./cmd/ultrasound-sensor/main.go && tinygo monitor --target=$(TARGET) --baudrate=9600

.PHONY: help
help: ## print this help and exit
	@echo "Usage: make [target]"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
