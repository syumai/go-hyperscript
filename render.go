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
	switch v := node.(type) {
	case Text:
		return createElementFromTextVNode(v)
	case *Element:
		return createElementFromElementVNode(v)
	default:
		return js.Null()
	}
}

func createElementFromTextVNode(node Text) js.Value {
	return document.Call("createTextNode", string(node))
}

func createElementFromElementVNode(node *Element) js.Value {
	el := document.Call("createElement", node.GetNodeName())
	for k, v := range node.Attributes {
		el.Set(k, v)
	}
	for _, child := range node.Children {
		el.Call("appendChild", createElement(child))
	}
	return el
}
