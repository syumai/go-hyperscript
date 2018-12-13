package hyperscript

type (
	Element struct {
		NodeName   string
		Children   VNodes
		Attributes Object
	}
)

func (vn *Element) GetNodeName() string {
	return vn.NodeName
}

func (vn *Element) GetNodeType() int {
	return NODE_TYPE_ELEMENT_NODE
}

func (vn *Element) GetChildren() VNodes {
	return vn.Children
}
