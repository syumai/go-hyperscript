package hyperscript

type (
	Element struct {
		NodeName   string
		NodeType   int
		Children   VNodes
		Attributes Props
	}
)

func (vn *Element) GetNodeName() string {
	return vn.NodeName
}

func (vn *Element) GetNodeType() int {
	return vn.NodeType
}

func (vn *Element) GetChildren() VNodes {
	return vn.Children
}
