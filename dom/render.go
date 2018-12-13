package dom

import (
	"strconv"
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
	case *h.TextElement:
		el = document.Call("createTextNode", n.TextContent)
		node.SetReference(jsValue(el))
	case *h.Element:
		el = document.Call("createElement", node.GetNodeName())
		setAttributes(el, n.Attributes)
		node.SetReference(jsValue(el))
		for _, child := range node.GetChildren() {
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
	println("replace element")
	newEl := createElement(newNode)
	oldEl := js.Value(oldNode.GetReference().(jsValue))
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

	elRef := js.Value(oldNode.GetReference().(jsValue))

	if oldNode.GetNodeType() != newNode.GetNodeType() {
		println("node type not equal")
		replaceElement(oldNode, newNode, parent)
		return
	}

	if newNode.GetNodeType() == h.NODE_TYPE_TEXT_NODE {
		println("node type is text node")
		oldText := oldNode.(*h.TextElement)
		newText := newNode.(*h.TextElement)
		newNode.SetReference(oldNode.GetReference())
		if oldText.TextContent == newText.TextContent {
			return
		}
		oldNode.GetReference().Set("textContent", newText.TextContent)
		return
	}

	println("node type" + strconv.Itoa(int(newNode.GetNodeType())))

	if newNode.GetNodeType() != h.NODE_TYPE_ELEMENT_NODE {
		println("node type is not element node")
		println(newNode.GetNodeName())
		// Not supported node type
		return
	}

	oldEl, ok := oldNode.(*h.Element)
	if !ok {
		return
	}

	newEl, ok := newNode.(*h.Element)
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

	newNode.SetReference(oldNode.GetReference())
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

	el := js.Value(node.GetReference().(jsValue))
	parent.Call("removeChild", el)
	node.SetReference(nil)
}
