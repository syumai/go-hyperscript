package hyperscript

type (
	Element struct {
		NodeName   string
		Children   VNodes
		Attributes Object
		Reference Value // Reference to real DOM
	}
)

func (el *Element) GetNodeName() string {
	return el.NodeName
}

func (el *Element) GetNodeType() int {
	return NODE_TYPE_ELEMENT_NODE
}

func (el *Element) GetChildren() VNodes {
	return el.Children
}

func (el *Element) SetReference(ref Value) {
	el.Reference = ref
}

func (el *Element) GetReference() Value {
	return el.Reference
}
