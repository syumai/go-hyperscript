package hyperscript

type (
	NodeType int
)

const (
	NODE_TYPE_UNKNOWN      NodeType = 0
	NODE_TYPE_ELEMENT_NODE          = 1
	NODE_TYPE_TEXT_NODE             = 3
)

func (n NodeType) String() string {
	switch n {
	case NODE_TYPE_ELEMENT_NODE:
		return "element"
	case NODE_TYPE_TEXT_NODE:
		return "text"
	}
	return "unknown"
}
