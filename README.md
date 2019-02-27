# go-hyperscript

* Simple [hyperscript](https://github.com/hyperhype/hyperscript) like script implemented in Go.
* You can use this package in these ways:
  - Write HTML using Go function.
  - Update existing DOM based on previous VNode tree.

## Usage

### Write HTML using Go function

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
	dom.NewRenderer().Render(node, body)
}
```

#### Example

* https://syumai.github.io/go-hyperscript/examples/basic/
  - [Code](https://github.com/syumai/go-hyperscript/tree/master/examples/basic/main.go)

### Update DOM

* Now go-hyperscript provides basic support for reusing and updating existing DOM.
* To use this function, please render multiple times using same renderer.
  - Hook state's changes (used in: [Simple ToDo list example](https://syumai.github.io/go-hyperscript/examples/simpletodo/))
  - Use main loop (used in: [Counter example](https://syumai.github.io/go-hyperscript/examples/counter/))

#### Examples

* [Simple ToDo list example](https://syumai.github.io/go-hyperscript/examples/simpletodo/)
  - [Code](https://github.com/syumai/go-hyperscript/tree/master/examples/simpletodo/main.go)

* [Counter example](https://syumai.github.io/go-hyperscript/examples/counter/)
  - [Code](https://github.com/syumai/go-hyperscript/tree/master/examples/counter/main.go)


## Environment

* go 1.12

## Development

```console
make build
make test
```

## Author

syumai

## License

MIT
