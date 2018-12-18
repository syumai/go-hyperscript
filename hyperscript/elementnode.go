package hyperscript

type elementNode struct {
	name       string
	children   VNodes
	attributes Object
	reference  Value // Reference to real DOM
}

type ElementNode interface {
	Name() string
	Children() VNodes
	Attributes() Object
}

func element(name string, attrs Object, children ...VNode) VNode {
	return &elementNode{
		name:       name,
		children:   children,
		attributes: attrs,
	}
}

func (el *elementNode) Name() string {
	return el.name
}

func (el *elementNode) Children() VNodes {
	return el.children
}

func (el *elementNode) Attributes() Object {
	return el.attributes
}

func (el *elementNode) NodeType() NodeType {
	return NodeTypeElementNode
}

func (el *elementNode) Reference() Value {
	return el.reference
}

func (el *elementNode) SetReference(ref Value) {
	el.reference = ref
}
