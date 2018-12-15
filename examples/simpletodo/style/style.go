package style

type Style map[string]string

func (s Style) String() string {
	var styleStr string
	for k, v := range s {
		styleStr += k + ": " + v + ";"
	}
	return styleStr
}
