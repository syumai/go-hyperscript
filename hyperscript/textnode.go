package hyperscript

type (
	TextNode struct {
		*Node
		TextContent string
	}
)

func Text(t string) VNode {
	return &TextNode{
		Node: &Node{
			nodeName: t,
		},
		TextContent: t,
	}
}

func (el *TextNode) NodeType() NodeType {
	return NodeTypeTextNode
}
