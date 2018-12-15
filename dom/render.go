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
	println("create element")
	var el js.Value
	switch n := node.(type) {
	case *h.TextNode:
		el = document.Call("createTextNode", n.TextContent)
		node.SetReference(jsValue(el))
	case *h.ElementNode:
		el = document.Call("createElement", node.NodeName())
		setAttributes(el, n.Attributes)
		node.SetReference(jsValue(el))
		for _, child := range node.Children() {
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
	println("replace element")
	newEl := createElement(newNode)
	oldEl := js.Value(oldNode.Reference().(jsValue))
	parent.Call("insertBefore", newEl, oldEl)
	removeElement(oldNode)
	newNode.SetReference(jsValue(newEl))
}

func updateElement(oldNode, newNode h.VNode) {
	println("update element")
	parent := getParentElement(oldNode)
	if parent == js.Null() {
		println("parent is null")
		return
	}

	elRef := js.Value(oldNode.Reference().(jsValue))

	if oldNode.NodeType() != newNode.NodeType() {
		println("node type not equal")
		replaceElement(oldNode, newNode, parent)
		return
	}

	if newNode.NodeType() == h.NodeTypeTextNode {
		println("node type is text node")
		oldText := oldNode.(*h.TextNode)
		newText := newNode.(*h.TextNode)
		newNode.SetReference(oldNode.Reference())
		if oldText.TextContent == newText.TextContent {
			return
		}
		oldNode.Reference().Set("textContent", newText.TextContent)
		return
	}

	println("node type " + newNode.NodeType().String())

	if newNode.NodeType() != h.NodeTypeElementNode {
		println("node type is not element node")
		println(newNode.NodeName())
		// Not supported node type
		return
	}

	oldEl, ok := oldNode.(*h.ElementNode)
	if !ok {
		return
	}

	newEl, ok := newNode.(*h.ElementNode)
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
	oldChilrenLen, newChildrenLen := len(oldNode.Children()), len(newNode.Children())
	if oldChilrenLen-newChildrenLen > 0 {
		for i := newChildrenLen; i < oldChilrenLen; i++ {
			removeElement(oldNode.Children()[i])
		}
	}

	for i, newChild := range newNode.Children() {
		// Create new node if old one does not exist
		if i >= len(oldNode.Children()) {
			node := createElement(newChild)
			elRef.Call("appendChild", node)
			continue
		}
		oldChild := oldNode.Children()[i]
		updateElement(oldChild, newChild)
	}

	newNode.SetReference(oldNode.Reference())
}

func removeElement(node h.VNode) {
	if node == nil {
		return
	}

	for _, childNode := range node.Children() {
		removeElement(childNode)
	}

	parent := getParentElement(node)
	if parent == js.Null() {
		return
	}

	el := js.Value(node.Reference().(jsValue))
	parent.Call("removeChild", el)
	node.SetReference(nil)
}
