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
			switch c := v.(type) {
			case h.Callback:
				el.Set(k, js.NewCallback(func(v []js.Value) {
					s := make([]h.Value, len(v))
					for i := 0; i < len(v); i++ {
						s[i] = jsValue(v[i])
					}
					c(s)
				}))
			case h.EventCallback:
				el.Set(k, js.NewEventCallback(js.EventCallbackFlag(c.Flg), func(event js.Value) {
					c.Func(jsValue(event))
				}))
			default:
				el.Set(k, v)
			}
		}
		for _, child := range n.Children {
			el.Call("appendChild", createElement(child))
		}
	default:
		el = document.Call("createTextNode", "")
	}
	return el
}
