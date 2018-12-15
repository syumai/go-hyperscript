package hyperscript

type (
	NodeType int
)

const (
	NodeTypeUnknown     NodeType = 0
	NodeTypeElementNode          = 1
	NodeTypeTextNode             = 3
)

func (n NodeType) String() string {
	switch n {
	case NodeTypeElementNode:
		return "element"
	case NodeTypeTextNode:
		return "text"
	}
	return "unknown"
}
