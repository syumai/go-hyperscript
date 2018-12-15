package style

type Prop [2]string

func Style(styles ...Prop) string {
	if len(styles) < 2 {
		return ""
	}
	var styleStr string
	for _, style := range styles {
		styleStr += style[0] + ": " + style[1] + ";"
	}
	return styleStr
}
