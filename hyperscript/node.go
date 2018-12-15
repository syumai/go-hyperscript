package hyperscript

type (
	Node struct {
		nodeName  string
		children  VNodes
		reference Value // Reference to real DOM
	}
)

func (el *Node) NodeName() string {
	return el.nodeName
}

func (el *Node) Children() VNodes {
	return el.children
}

func (el *Node) Reference() Value {
	return el.reference
}

func (el *Node) SetReference(ref Value) {
	el.reference = ref
}
