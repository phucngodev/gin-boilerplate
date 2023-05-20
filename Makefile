export TARGET_WEB=web

.PHONY: build
build: wire
	go build -o ./bin/$(TARGET_WEB) ./cmd/$(TARGET_WEB)

.PHONY: wire
wire:
	wire ./...
