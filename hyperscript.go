package hyperscript

const (
	NODE_TYPE_ELEMENT_NODE = 1
	NODE_TYPE_TEXT_NODE    = 3
)

type (
	Object map[string]interface{}
	Array  []interface{}
)

type (
	Element interface {
		GetNodeName() string
		GetNodeType() int
		GetChildren() Elements
		ToString() string
	}

	Elements []Element
)

func (elements Elements) ToString() string {
	var str string
	for _, el := range elements {
		str += el.ToString()
	}
	return str
}

type (
	Component func() Element
)

func mergeElements(a Elements, b Elements) Elements {
	lenA := len(a)
	ary := make(Elements, lenA+len(b))
	for i, v := range a {
		ary[i] = v
	}
	for i, v := range b {
		ary[i+lenA] = v
	}
	return ary
}

func mergeObjects(a Object, b Object) Object {
	obj := Object{}
	for k, v := range a {
		obj[k] = v
	}
	for k, v := range b {
		obj[k] = v
	}
	return obj
}

func H(tag interface{}, attrs Object, children ...Element) Element {
	if c, ok := tag.(Component); ok {
		node := c()
		if n, ok := node.(*Node); ok {
			n.Attributes = mergeObjects(n.Attributes, attrs)
			n.Children = mergeElements(n.Children, children)
		}
		return node
	}
	s, ok := tag.(string)
	if !ok {
		return TextNode("")
	}
	return &Node{
		NodeName:   s,
		NodeType:   NODE_TYPE_ELEMENT_NODE,
		Children:   children,
		Attributes: attrs,
	}
}
