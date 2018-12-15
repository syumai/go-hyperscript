.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o ./examples/basic/index.wasm ./examples/basic
	GOOS=js GOARCH=wasm go build -o ./examples/counter/index.wasm ./examples/counter
	GOOS=js GOARCH=wasm go build -o ./examples/simpletodo/index.wasm ./examples/simpletodo

.PHONY: test
test:
	go test -v ./hyperscript \
		./examples/counter/counter
