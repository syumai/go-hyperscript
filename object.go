package hyperscript

type Object map[string]interface{}

func (o Object) Get(key string) interface{} {
	if v, ok := o[key]; ok {
		return v
	}
	return nil
}

func (o Object) String(key string) string {
	if v, ok := o[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func (o Object) Int(key string) int {
	if v, ok := o[key]; ok {
		if i, ok := v.(int); ok {
			return i
		}
	}
	return 0
}

func (o Object) Strings(key string) []string {
	if v, ok := o[key]; ok {
		if strs, ok := v.([]string); ok {
			return strs
		}
	}
	return []string{}
}

func (o Object) Ints(key string) []int {
	if v, ok := o[key]; ok {
		if ints, ok := v.([]int); ok {
			return ints
		}
	}
	return []int{}
}
