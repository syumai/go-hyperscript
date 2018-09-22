package dom

import (
	"syscall/js"

	h "github.com/syumai/go-hyperscript/hyperscript"
)

type jsValue js.Value

func (v jsValue) Get(p string) h.Value {
	jv := js.Value(v)
	return jsValue(jv.Get(p))
}

func (v jsValue) Set(p string, x interface{}) {
	jv := js.Value(v)
	jv.Set(p, x)
}

func (v jsValue) Index(i int) h.Value {
	jv := js.Value(v)
	return jsValue(jv.Index(i))
}

func (v jsValue) SetIndex(i int, x interface{}) {
	jv := js.Value(v)
	jv.SetIndex(i, x)
}

func (v jsValue) Length() int {
	jv := js.Value(v)
	return jv.Length()
}

func (v jsValue) Call(m string, args ...interface{}) h.Value {
	jv := js.Value(v)
	return jsValue(jv.Call(m, args...))
}

func (v jsValue) Invoke(args ...interface{}) h.Value {
	jv := js.Value(v)
	return jsValue(jv.Invoke(args...))
}

func (v jsValue) New(args ...interface{}) h.Value {
	jv := js.Value(v)
	return jsValue(jv.New(args...))
}

func (v jsValue) Float() float64 {
	jv := js.Value(v)
	return jv.Float()
}

func (v jsValue) Int() int {
	jv := js.Value(v)
	return jv.Int()
}

func (v jsValue) Bool() bool {
	jv := js.Value(v)
	return jv.Bool()
}

func (v jsValue) String() string {
	jv := js.Value(v)
	return jv.String()
}

func (v jsValue) InstanceOf(t h.Value) bool {
	jv := js.Value(v)
	if jt, ok := t.(jsValue); ok {
		return jv.InstanceOf(js.Value(jt))
	}
	return false
}
