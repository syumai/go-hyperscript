package hyperscript

type textNode struct {
	*Node
	textContent string
}

type TextContenter interface {
	TextContent() string
}

func Text(t string) VNode {
	return &textNode{
		Node: &Node{
			nodeName: t,
		},
		textContent: t,
	}
}

func (el *textNode) TextContent() string {
	return el.textContent
}

func (el *textNode) NodeType() NodeType {
	return NodeTypeTextNode
}
