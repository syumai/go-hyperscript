package dom

import (
	"syscall/js"
	h "github.com/syumai/go-hyperscript/hyperscript"
)

var (
	window   = js.Global().Get("window")
	document = js.Global().Get("document")
)

func Render(node h.VNode, container js.Value) {
	container.Call("appendChild", createElement(node))
}

func createElement(node h.VNode) js.Value {
	var el js.Value
	switch n := node.(type) {
	case h.Text:
		el = document.Call("createTextNode", string(n))
	case *h.Element:
		el = document.Call("createElement", n.GetNodeName())
		for k, v := range n.Attributes {
			el.Call("setAttribute", k, v)
			el.Set(k, v)
		}
		for _, child := range n.Children {
			el.Call("appendChild", createElement(child))
		}
		n.Base = el
	default:
		el = document.Call("createTextNode", "")
	}
	return el
}
