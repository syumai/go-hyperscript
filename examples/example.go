package main

import (
	"fmt"

	h "github.com/syumai/go-hyperscript"
)

func List(props h.Object) h.Element {
	names := props["names"].([]string)
	elements := make(h.Elements, len(names))
	for i, name := range names {
		elements[i] = h.H("li", nil, h.TextNode(name))
	}
	return h.H("ul", nil, elements...)
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
