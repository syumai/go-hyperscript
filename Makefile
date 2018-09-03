GO_BIN = go1.11

.PHONY: build
build:
	GOOS=js GOARCH=wasm $(GO_BIN) build -o ./examples/basic/index.wasm ./examples/basic
