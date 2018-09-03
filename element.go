package hyperscript

import "fmt"

type (
	Element struct {
		NodeName   string
		NodeType   int
		Children   VNodes
		Attributes Object
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

func (vn *Element) ToString() string {
	var attrs string
	for k, v := range vn.Attributes {
		attrs += fmt.Sprintf(` %s="%s"`, k, v)
	}
	return fmt.Sprintf(`<%s%s>%s</%s>`, vn.GetNodeName(), attrs, vn.GetChildren().ToString(), vn.GetNodeName())
}
