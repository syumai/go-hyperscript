package hyperscript

type (
	VNode interface {
		NodeType() NodeType
		Reference() Value
		SetReference(Value)
	}

	VNodes []VNode
)
