# go-hyperscript

* Simple [hyperscript](https://github.com/hyperhype/hyperscript) like script implemented in golang.
* Write HTML using golang function.

## Usage

### Render HTML with wasm

```go
func main() {
	node := h.H("div", nil,
		h.H("h1", nil, h.Text("Example App")),
		h.H("strong", nil,
			h.H("font", h.Object{"color": "red"}, h.Text("Hello, world!")),
		),
		...
	)
	body := js.Global().Get("document").Get("body")
	h.Render(node, body)
}
```

#### Result

* https://syumai.github.io/go-hyperscript/examples/basic/
  - [Code](https://github.com/syumai/go-hyperscript/tree/master/examples/basic/main.go)

## Other Examples

* [Counter example](https://syumai.github.io/go-hyperscript/examples/counter/)
  - [Code](https://github.com/syumai/go-hyperscript/tree/master/examples/counter/main.go)

## Environment

* go 1.11

## Development

```console
make build
make test
```

## Author

syumai

## License

MIT
