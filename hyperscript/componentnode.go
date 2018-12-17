package hyperscript

import "unsafe"

type componentNode struct {
	component  Component
	nodeTree   VNode
	attributes Object
}

func component(comp Component, attrs Object) VNode {
	return &componentNode{
		component:  comp,
		attributes: attrs,
	}
}

func (el *componentNode) Type() NodeType {
	return NodeTypeComponentNode
}

func (el *componentNode) Children() VNodes {
	return nil
}

func (el *componentNode) Reference() Value {
	if el.nodeTree != nil {
		return el.nodeTree.Reference()
	}
	return nil
}

func (el *componentNode) SetReference(ref Value) {
	// Do nothing
}

func (el *componentNode) NodeTree() VNode {
	return el.nodeTree
}

func (el *componentNode) SetNodeTree(nt VNode) {
	el.nodeTree = nt
}

func (el *componentNode) Attributes() Object {
	return el.attributes
}

func (el *componentNode) Component() Component {
	return el.component
}

func (el *componentNode) ComponentPointer() unsafe.Pointer {
	i := (interface{})(el.component)
	return (*pointerGetter)(unsafe.Pointer(&i)).ptr
}
