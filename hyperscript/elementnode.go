package hyperscript

type elementNode struct {
	*Node
	attributes Object
}

type Attributer interface {
	Attributes() Object
}

func element(nodeName string, attrs Object, children ...VNode) VNode {
	return &elementNode{
		Node: &Node{
			nodeName: nodeName,
			children: children,
		},
		attributes: attrs,
	}
}

func (el *elementNode) Attributes() Object {
	return el.attributes
}

func (el *elementNode) NodeType() NodeType {
	return NodeTypeElementNode
}
