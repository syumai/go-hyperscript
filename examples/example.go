package main

import (
	"fmt"

	h "github.com/syumai/go-hyperscript"
)

func main() {
	node := h.H("div", nil,
		h.H("h1", nil, h.TextNode("Example App")),
		h.H("strong", nil,
			h.H("font", h.Object{"color": "red"}, h.TextNode("Hello, world!")),
		),
	)
	fmt.Println(node.ToString())
}
