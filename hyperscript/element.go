package hyperscript

type (
	Element struct {
		NodeName   string
		NodeType   NodeType
		Children   VNodes
		Attributes Object
		Reference  Value // Reference to real DOM
	}
)

func (el *Element) GetNodeName() string {
	return el.NodeName
}

func (el *Element) GetNodeType() NodeType {
	return el.NodeType
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
