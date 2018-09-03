# go-hyperscript

[WIP]

* hyperscript implementation in golang.
* Write HTML using function.

## Usage

```go
func main() {
	node := h.H("div", nil,
		h.H("h1", nil, h.Text("Example App")),
		h.H("strong", nil,
			h.H("font", h.Object{"color": "red"}, h.Text("Hello, world!")),
		),
	)
	fmt.Println(node.ToString())
}
```

### Output

```html
<div><h1>Example App</h1><strong><font color="red">Hello, world!</font></strong></div>
```

## Author

syumai

## License

MIT
