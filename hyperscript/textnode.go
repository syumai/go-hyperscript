package hyperscript

type textNode struct {
	content   string
	reference Value // Reference to real DOM
}

type TextNode interface {
	VNode
	Content() string
}

func Text(t string) TextNode {
	return &textNode{
		content: t,
	}
}

func (el *textNode) Content() string {
	return el.content
}

func (el *textNode) NodeType() NodeType {
	return NodeTypeTextNode
}

func (el *textNode) Reference() Value {
	return el.reference
}

func (el *textNode) SetReference(ref Value) {
	el.reference = ref
}
