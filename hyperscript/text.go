package hyperscript

type (
	Text string
)

func (tn Text) GetNodeName() string {
	return string(tn)
}

func (tn Text) GetNodeType() int {
	return NODE_TYPE_TEXT_NODE
}

func (tn Text) GetChildren() VNodes {
	return []VNode{}
}
