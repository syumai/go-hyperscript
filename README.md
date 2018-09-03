# go-hyperscript

* Simple hyperscript like scripts implemented in golang.
* Write HTML using golang function.

## Usage

### Render to HTML using wasm

```go
func main() {
	node := h.H("div", nil,
		h.H("h1", nil, h.Text("Example App")),
		h.H("strong", nil,
			h.H("font", h.Object{"color": "red"}, h.Text("Hello, world!")),
		),
	)
	body := js.Global().Get("document").Get("body")
	h.Render(node, body)
}
```

#### Result

https://syumai.github.io/go-hyperscript/examples/basic/

## Environment

* go 1.11

## Development

```console
go get golang.org/dl/go1.11
go1.11 download

make build
```

## Author

syumai

## License

MIT
