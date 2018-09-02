package hyperscript

import "fmt"

const (
	NODE_TYPE_ELEMENT_NODE = 1
	NODE_TYPE_TEXT_NODE    = 3
)

type (
	Object map[string]interface{}
	Array  []interface{}
)

type (
	VNode struct {
		NodeName   string
		NodeType   int
		Children   Elements
		Attributes Object
	}

	TextNode string
)

type (
	Element interface {
		GetNodeName() string
		GetNodeType() int
		GetChildren() Elements
		ToString() string
	}

	Elements []Element
)

func (vn *VNode) GetNodeName() string {
	return vn.NodeName
}

func (vn *VNode) GetNodeType() int {
	return vn.NodeType
}

func (vn *VNode) GetChildren() Elements {
	return vn.Children
}

func (vn *VNode) ToString() string {
	var attrs string
	for k, v := range vn.Attributes {
		attrs += fmt.Sprintf(` %s="%s"`, k, v)
	}
	return fmt.Sprintf(`<%s%s>%s</%s>`, vn.GetNodeName(), attrs, vn.GetChildren().ToString(), vn.GetNodeName())
}

func (tn TextNode) GetNodeName() string {
	return string(tn)
}

func (tn TextNode) GetNodeType() int {
	return NODE_TYPE_TEXT_NODE
}

func (tn TextNode) GetChildren() Elements {
	return []Element{}
}

func (tn TextNode) ToString() string {
	return string(tn)
}

func (elements Elements) ToString() string {
	var str string
	for _, el := range elements {
		str += el.ToString()
	}
	return str
}

func H(tag string, attrs Object, elements ...Element) Element {
	return &VNode{
		NodeName:   tag,
		NodeType:   NODE_TYPE_ELEMENT_NODE,
		Children:   elements,
		Attributes: attrs,
	}
}
