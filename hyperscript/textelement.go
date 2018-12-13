package hyperscript

type (
	TextElement struct {
		*Element
		TextContent string
	}
)

func Text(t string) VNode {
	return &TextElement{
		Element: &Element{
			NodeName: t,
			NodeType: NODE_TYPE_TEXT_NODE,
		},
		TextContent: t,
	}
}
