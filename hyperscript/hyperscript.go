package hyperscript

type (
	VNode interface {
		NodeName() string
		NodeType() NodeType
		Children() VNodes
		Reference() Value
		SetReference(Value)
	}

	VNodes []VNode
)

type (
	Component = func(props Object) VNode
)

var (
	BlankElement   = Text("")
	BlankComponent = func(Object) VNode { return BlankElement }
)

func H(tag interface{}, attrs Object, children ...VNode) VNode {
	switch v := tag.(type) {
	case Component:
		return v(attrs)
	case string:
		return element(v, attrs, children...)
	default:
		return BlankElement
	}
}
