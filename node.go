package hyperscript

import "fmt"

type (
	Node struct {
		NodeName   string
		NodeType   int
		Children   Elements
		Attributes Object
	}
)

func (vn *Node) GetNodeName() string {
	return vn.NodeName
}

func (vn *Node) GetNodeType() int {
	return vn.NodeType
}

func (vn *Node) GetChildren() Elements {
	return vn.Children
}

func (vn *Node) ToString() string {
	var attrs string
	for k, v := range vn.Attributes {
		attrs += fmt.Sprintf(` %s="%s"`, k, v)
	}
	return fmt.Sprintf(`<%s%s>%s</%s>`, vn.GetNodeName(), attrs, vn.GetChildren().ToString(), vn.GetNodeName())
}
