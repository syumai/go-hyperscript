package hyperscript

type Object map[string]interface{}

func (o Object) Get(key string) interface{} {
	if v, ok := o[key]; ok {
		return v
	}
	return nil
}

func (o Object) Set(key string, value interface{}) {
	o[key] = value
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

func (o Object) Bool(key string) bool {
	if v, ok := o[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}

func (o Object) Func(key string) Func {
	if v, ok := o[key]; ok {
		if c, ok := v.(Func); ok {
			return c
		}
	}
	return Func{}
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

func (o Object) Bools(key string) []bool {
	if v, ok := o[key]; ok {
		if b, ok := v.([]bool); ok {
			return b
		}
	}
	return []bool{}
}

func (o Object) Funcs(key string) []Func {
	if v, ok := o[key]; ok {
		if cbs, ok := v.([]Func); ok {
			return cbs
		}
	}
	return []Func{}
}

func (o Object) Key() string {
	return o.String("key")
}

func ObjectEqual(a, b Object) bool {
	if len(a) != len(b) {
		return false
	}

	for k, va := range a {
		if vb, ok := b[k]; ok {
			if va != vb {
				return false
			}
		}
	}

	return true
}

func ObjectDiff(old, new Object) Object {
	diff := make(Object)

	for k, v := range old {
		if _, ok := new[k]; !ok {
			diff[k] = v
		}
	}

	return diff
}
