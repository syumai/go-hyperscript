package hyperscript

type (
	VNode interface {
		GetNodeName() string
		GetNodeType() NodeType
		GetChildren() VNodes
		SetReference(Value)
		GetReference() Value
	}

	VNodes []VNode

	NodeType int
)

type (
	StatelessComponent func(props Object) VNode
)

const (
	NODE_TYPE_UNKNOWN      NodeType = 0
	NODE_TYPE_ELEMENT_NODE          = 1
	NODE_TYPE_TEXT_NODE             = 3
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
			NodeType:   NODE_TYPE_ELEMENT_NODE,
			Children:   children,
			Attributes: attrs,
		}
	default:
		return BlankElement
	}
}
