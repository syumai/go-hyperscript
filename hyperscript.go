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

func H(tag string, attrs Object, elements ...Element) Element {
	return &Node{
		NodeName:   tag,
		NodeType:   NODE_TYPE_ELEMENT_NODE,
		Children:   elements,
		Attributes: attrs,
	}
}
