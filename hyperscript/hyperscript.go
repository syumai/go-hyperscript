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
	StatelessComponent func(props Object) VNode
)

var (
	BlankElement   = Text("")
	BlankComponent = func(Object) VNode { return BlankElement }
)

func H(tag interface{}, attrs Object, children ...VNode) VNode {
	switch v := tag.(type) {
	case StatelessComponent:
		return v(attrs)
	case func(Object) VNode:
		return v(attrs)
	case string:
		return Element(v, attrs, children...)
	default:
		return BlankElement
	}
}
