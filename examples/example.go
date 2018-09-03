package main

import (
	"fmt"

	h "github.com/syumai/go-hyperscript"
)

func mapToList(names ...string) h.Elements {
	elements := make(h.Elements, len(names))
	for i, name := range names {
		elements[i] = h.H("li", nil, h.TextNode(name))
	}
	return elements
}

func List(props h.Object) h.Element {
	return h.H("ul", nil, mapToList(props.Strings("names")...)...)
}

func main() {
	node := h.H("div", nil,
		h.H("h1", nil, h.TextNode("Example App")),
		h.H("strong", nil,
			h.H("font", h.Object{"color": "red"}, h.TextNode("Hello, world!")),
		),
		h.H(List, h.Object{"names": []string{"a", "b", "c"}}),
		h.H(List, h.Object{"names": []string{"d", "e", "f"}}),
	)
	fmt.Println(node.ToString())
}
