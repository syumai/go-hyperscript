package basic

import (
	"syscall/js"

	h "github.com/syumai/go-hyperscript"
)

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
		h.H(List, h.Object{"names": []string{"a", "b", "c"}}),
		h.H(List, h.Object{"names": []string{"d", "e", "f"}}),
	)
	body := js.Global().Get("document").Get("body")
	h.Render(node, body)
}
