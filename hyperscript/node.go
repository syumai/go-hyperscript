package hyperscript

import "unsafe"

type ElementNode interface {
	Name() string
	Attributes() Object
}

type TextNode interface {
	TextContent() string
}

type ComponentNode interface {
	Component() Component
	ComponentPointer() unsafe.Pointer
	Attributes() Object
	NodeTree() VNode
	SetNodeTree(nt VNode)
}
