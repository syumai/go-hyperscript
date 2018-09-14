package hyperscript

import (
	"syscall/js"
)

var (
	window   = js.Global().Get("window")
	document = js.Global().Get("document")
)

func Render(node VNode, container js.Value) {
	container.Call("appendChild", createElement(node))
}

func createElement(node VNode) js.Value {
	var el js.Value
	switch n := node.(type) {
	case Text:
		el = document.Call("createTextNode", string(n))
	case *Element:
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
