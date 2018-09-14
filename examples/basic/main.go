package main

import (
	"syscall/js"

	h "github.com/syumai/go-hyperscript"
)

var body = js.Global().Get("document").Get("body")

func mapToList(names ...string) h.VNodes {
	elements := make(h.VNodes, len(names))
	for i, name := range names {
		elements[i] = h.H("li", nil, h.Text(name))
	}
	return elements
}

func List(props h.Object) h.VNode {
	return h.H("ul", nil, mapToList(props.Strings("names")...)...)
}

func main() {
	node := h.H("div", nil,
		h.H("h1", nil, h.Text("Example App")),
		h.H("strong", nil,
			h.H("font", h.Object{"color": "red"}, h.Text("Hello, world!")),
		),
		h.H("h2", nil, h.Text("List")),
		h.H(List, h.Object{"names": []string{"a", "b", "c"}}),
		h.H("h2", nil, h.Text("Count")),
		h.H("a", h.Object{"href": "https://github.com/syumai/go-hyperscript/"},
			h.Text("Show the code on GitHub"),
		),
	)
	h.Render(node, body)
}
