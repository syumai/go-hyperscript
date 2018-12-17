package hyperscript

type VNode interface {
	Type() NodeType
	Children() VNodes
	Reference() Value
	SetReference(Value)
}

type VNodes []VNode
