package hyperscript

type textNode struct {
	reference   Value
	textContent string
}

func Text(t string) VNode {
	return &textNode{
		textContent: t,
	}
}

func (el *textNode) Type() NodeType {
	return NodeTypeTextNode
}

func (el *textNode) Children() VNodes {
	return nil
}

func (el *textNode) Reference() Value {
	return el.reference
}

func (el *textNode) SetReference(ref Value) {
	el.reference = ref
}

func (el *textNode) TextContent() string {
	return el.textContent
}
