package hyperscript

const (
	NODE_TYPE_ELEMENT_NODE = 1
	NODE_TYPE_TEXT_NODE    = 3
)

type Object map[string]interface{}

func (o Object) Get(key string) interface{} {
	if v, ok := o[key]; ok {
		return v
	}
	return nil
}

func (o Object) String(key string) string {
	if v, ok := o[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func (o Object) Int(key string) int {
	if v, ok := o[key]; ok {
		if i, ok := v.(int); ok {
			return i
		}
	}
	return 0
}

func (o Object) Strings(key string) []string {
	if v, ok := o[key]; ok {
		if strs, ok := v.([]string); ok {
			return strs
		}
	}
	return []string{}
}

func (o Object) Ints(key string) []int {
	if v, ok := o[key]; ok {
		if ints, ok := v.([]int); ok {
			return ints
		}
	}
	return []int{}
}

type (
	VNode interface {
		GetNodeName() string
		GetNodeType() int
		GetChildren() VNodes
	}

	VNodes []VNode
)

type (
	Component func(props Object) VNode
)

var (
	BlankElement   = Text("")
	BlankComponent = func(Object) VNode { return BlankElement }
)

func H(tag interface{}, attrs Object, children ...VNode) VNode {
	switch v := tag.(type) {
	case Component:
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
