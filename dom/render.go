package dom

import (
	"syscall/js"

	h "github.com/syumai/go-hyperscript/hyperscript"
)

var (
	window   = js.Global().Get("window")
	document = js.Global().Get("document")
)

type Renderer struct {
	oldNode h.VNode
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Render(node h.VNode, container js.Value) {
	if r.oldNode == nil {
		container.Call("appendChild", createElement(node))
	} else {
		updateElement(r.oldNode, node)
	}
	r.oldNode = node
}

func createElement(node h.VNode) js.Value {
	var el js.Value
	switch n := node.(type) {
	case h.Text:
		el = document.Call("createTextNode", string(n))
	case *h.Element:
		el = document.Call("createElement", n.GetNodeName())
		setAttributes(el, n.Attributes)
		node.SetReference(jsValue(el))
		for _, child := range n.Children {
			el.Call("appendChild", createElement(child))
		}
	default:
		el = document.Call("createTextNode", "")
	}
	return el
}

func setAttributes(el js.Value, attrs h.Object) {
	for k, v := range attrs {
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
}

func getParentElement(node h.VNode) js.Value {
	if node == nil {
		return js.Null()
	}

	ref := node.GetReference()
	if ref == nil {
		return js.Null()
	}

	parent, ok := ref.Get("parentNode").(jsValue)
	if !ok {
		return js.Null()
	}
	return js.Value(parent)
}

func replaceElement(oldNode, newNode h.VNode, parent js.Value) {
	node := createElement(newNode)
	parent.Call("insertBefore", node, oldNode.GetReference())
	removeElement(oldNode)
}

func updateElement(oldNode, newNode h.VNode) {
	parent := getParentElement(oldNode)
	if parent == js.Null() {
		return
	}

	elRef := js.Value(oldNode.GetReference().(jsValue))

	if oldNode.GetNodeType() != newNode.GetNodeType() {
		replaceElement(oldNode, newNode, parent)
		return
	}

	if newNode.GetNodeType() == h.NODE_TYPE_TEXT_NODE {
		replaceElement(oldNode, newNode, parent)
		return
	}

	if newNode.GetNodeType() != h.NODE_TYPE_ELEMENT_NODE {
		// Not supported node type
		return
	}

	oldEl, ok := oldNode.(*h.Element)
	if !ok {
		return
	}

	newEl, ok := oldNode.(*h.Element)
	if !ok {
		return
	}

	// Update properties
	if !h.ObjectEqual(oldEl.Attributes, newEl.Attributes) {
		for k, _ := range h.ObjectDiff(oldEl.Attributes, newEl.Attributes) {
			elRef.Set(k, js.Undefined())
		}
		setAttributes(elRef, newEl.Attributes)
	}

	// Remove unused children
	childCountDiff := len(oldNode.GetChildren()) - len(newNode.GetChildren())
	if childCountDiff > 0 {
		for i := len(newNode.GetChildren()) - 1; i < childCountDiff; i++ {
			removeElement(oldNode.GetChildren()[i])
		}
	}

	for i, newChild := range newNode.GetChildren() {
		// Create new node if old one does not exist
		if i >= len(oldNode.GetChildren()) {
			node := createElement(newChild)
			elRef.Call("appendChild", node)
			continue
		}
		oldChild := oldNode.GetChildren()[i]
		updateElement(oldChild, newChild)
	}
}

func removeElement(node h.VNode) {
	if node == nil {
		return
	}

	for _, childNode := range node.GetChildren() {
		removeElement(childNode)
	}

	parent := getParentElement(node)
	if parent == js.Null() {
		return
	}

	parent.Call("removeChild", node.GetReference())
	node.SetReference(nil)
}
