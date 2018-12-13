package hyperscript

const (
	NODE_TYPE_ELEMENT_NODE = 1
	NODE_TYPE_TEXT_NODE    = 3
)

type (
	VNode interface {
		GetNodeName() string
		GetNodeType() int
		GetChildren() VNodes
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
		return &Element{
			NodeName:   v,
			Children:   children,
			Attributes: attrs,
		}
	default:
		return BlankElement
	}
}
