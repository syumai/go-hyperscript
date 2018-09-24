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
	if v, ok := tag.(func(Object) VNode); ok {
		v = StatelessComponent(v)
	}
	switch v := tag.(type) {
	case Component:
		v.Initialize(v.InitialState(), attrs, func() {})
		return v.Render()
	case StatelessComponent:
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
