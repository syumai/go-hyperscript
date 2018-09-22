package hyperscript

type Object interface {
	Get(string) interface{}
	String(string) string
	Strings(string) []string
	Int(string) int
	Ints(string) []int
	Bool(string) bool
	Bools(string) []bool
	Key() string
}

type Props map[string]interface{}

func (o Props) Get(key string) interface{} {
	if v, ok := o[key]; ok {
		return v
	}
	return nil
}

func (o Props) String(key string) string {
	if v, ok := o[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func (o Props) Int(key string) int {
	if v, ok := o[key]; ok {
		if i, ok := v.(int); ok {
			return i
		}
	}
	return 0
}

func (o Props) Bool(key string) bool {
	if v, ok := o[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}

func (o Props) Strings(key string) []string {
	if v, ok := o[key]; ok {
		if strs, ok := v.([]string); ok {
			return strs
		}
	}
	return []string{}
}

func (o Props) Ints(key string) []int {
	if v, ok := o[key]; ok {
		if ints, ok := v.([]int); ok {
			return ints
		}
	}
	return []int{}
}

func (o Props) Bools(key string) []bool {
	if v, ok := o[key]; ok {
		if b, ok := v.([]bool); ok {
			return b
		}
	}
	return []bool{}
}

func (o Props) Key() string {
	return o.String("key")
}
