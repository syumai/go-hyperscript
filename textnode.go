package hyperscript

type (
	TextNode string
)

func (tn TextNode) GetNodeName() string {
	return string(tn)
}

func (tn TextNode) GetNodeType() int {
	return NODE_TYPE_TEXT_NODE
}

func (tn TextNode) GetChildren() Elements {
	return []Element{}
}

func (tn TextNode) ToString() string {
	return string(tn)
}
