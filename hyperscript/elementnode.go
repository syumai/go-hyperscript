package hyperscript

type elementNode struct {
	name       string
	children   VNodes
	reference  Value
	attributes Object
}

func element(name string, attrs Object, children ...VNode) VNode {
	return &elementNode{
		name:       name,
		children:   children,
		attributes: attrs,
	}
}

func (el *elementNode) Type() NodeType {
	return NodeTypeElementNode
}

func (el *elementNode) Children() VNodes {
	return el.children
}

func (el *elementNode) Reference() Value {
	return el.reference
}

func (el *elementNode) SetReference(ref Value) {
	el.reference = ref
}

func (el *elementNode) Name() string {
	return el.name
}

func (el *elementNode) Attributes() Object {
	return el.attributes
}
