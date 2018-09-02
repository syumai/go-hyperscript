package main

import (
	"fmt"

	h "github.com/syumai/go-hyperscript"
)

func List(names ...string) h.Component {
	elements := make(h.Elements, len(names))
	for i, name := range names {
		elements[i] = h.H("li", nil, h.TextNode(name))
	}
	return func() h.Element { return h.H("ul", nil, elements...) }
}

func main() {
	node := h.H("div", nil,
		h.H("h1", nil, h.TextNode("Example App")),
		h.H("strong", nil,
			h.H("font", h.Object{"color": "red"}, h.TextNode("Hello, world!")),
		),
		h.H(List("a", "b", "c"), nil),
		h.H(List("d", "e", "f"), nil),
	)
	fmt.Println(node.ToString())
}
