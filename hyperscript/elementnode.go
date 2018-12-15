package hyperscript

type (
	ElementNode struct {
		*Node
		Attributes Object
	}
)

func element(nodeName string, attrs Object, children ...VNode) VNode {
	return &ElementNode{
		Node: &Node{
			nodeName: nodeName,
			children: children,
		},
		Attributes: attrs,
	}
}

func (el *ElementNode) NodeType() NodeType {
	return NODE_TYPE_ELEMENT_NODE
}
