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
	case h.TextNode:
		el = document.Call("createTextNode", n.Content())
		node.SetReference(jsValue(el))
	case h.ElementNode:
		el = document.Call("createElement", n.Name())
		setAttributes(el, n.Attributes())
		node.SetReference(jsValue(el))
		for _, child := range n.Children() {
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
				c.Call(s)
			}))
		case h.EventCallback:
			el.Set(k, js.NewEventCallback(js.EventCallbackFlag(c.Flg), func(event js.Value) {
				c.Call(jsValue(event))
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

	ref := node.Reference()
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
	newEl := createElement(newNode)
	oldEl := js.Value(oldNode.Reference().(jsValue))
	parent.Call("insertBefore", newEl, oldEl)
	removeElement(oldNode)
	newNode.SetReference(jsValue(newEl))
}

func updateElement(oldNode, newNode h.VNode) {
	parent := getParentElement(oldNode)
	if parent == js.Null() {
		return
	}

	elRef := js.Value(oldNode.Reference().(jsValue))

	if oldNode.NodeType() != newNode.NodeType() {
		replaceElement(oldNode, newNode, parent)
		return
	}

	if newNode.NodeType() == h.NodeTypeTextNode {
		oldText := oldNode.(h.TextNode)
		newText := newNode.(h.TextNode)
		newNode.SetReference(oldNode.Reference())
		if oldText.Content() == newText.Content() {
			return
		}
		oldNode.Reference().Set("textContent", newText.Content())
		return
	}

	if newNode.NodeType() != h.NodeTypeElementNode {
		// Not supported node type
		return
	}

	oldEl, ok := oldNode.(h.ElementNode)
	if !ok {
		return
	}

	newEl, ok := newNode.(h.ElementNode)
	if !ok {
		return
	}

	// Update properties
	if !h.ObjectEqual(oldEl.Attributes(), newEl.Attributes()) {
		for k, _ := range h.ObjectDiff(oldEl.Attributes(), newEl.Attributes()) {
			elRef.Set(k, js.Undefined())
		}
		setAttributes(elRef, newEl.Attributes())
	}

	// Remove unused children
	oldChilrenLen, newChildrenLen := len(oldEl.Children()), len(newEl.Children())
	if oldChilrenLen-newChildrenLen > 0 {
		for i := newChildrenLen; i < oldChilrenLen; i++ {
			removeElement(oldEl.Children()[i])
		}
	}

	for i, newChild := range newEl.Children() {
		// Create new node if old one does not exist
		if i >= len(oldEl.Children()) {
			node := createElement(newChild)
			elRef.Call("appendChild", node)
			continue
		}
		oldChild := oldEl.Children()[i]
		updateElement(oldChild, newChild)
	}

	newNode.SetReference(oldNode.Reference())
}

func removeElement(node h.VNode) {
	if node == nil {
		return
	}

	if el, ok := node.(h.ElementNode); ok {
		for _, childNode := range el.Children() {
			removeElement(childNode)
		}
	}

	parent := getParentElement(node)
	if parent == js.Null() {
		return
	}

	el := js.Value(node.Reference().(jsValue))
	parent.Call("removeChild", el)
	node.SetReference(nil)
}
