package fakevalue

import (
	"reflect"

	h "github.com/syumai/go-hyperscript/hyperscript"
)

type fakeValue struct {
	value interface{}
}

func New(v interface{}) h.Value {
	return fakeValue{v}
}

func (f fakeValue) Get(p string) h.Value {
	v, ok := f.value.(h.Object)
	if !ok {
		return fakeValue{nil}
	}
	return fakeValue{v.Get(p)}
}

func (f fakeValue) Set(p string, x interface{}) {
	v, ok := f.value.(h.Object)
	if !ok {
		return
	}
	v.Set(p, x)
}

func (f fakeValue) Index(i int) h.Value {
	v, ok := f.value.([]interface{})
	if !ok {
		return fakeValue{nil}
	}
	if len(v) <= i {
		return fakeValue{nil}
	}
	return fakeValue{v[i]}
}

func (f fakeValue) SetIndex(i int, x interface{}) {
	v, ok := f.value.([]interface{})
	if !ok {
		return
	}
	if len(v) <= i {
		return
	}
	v[i] = x
}

func (f fakeValue) Length() int {
	v, ok := f.value.([]interface{})
	if !ok {
		return 0
	}
	return len(v)
}

func (f fakeValue) Call(m string, args ...interface{}) h.Value {
	v := f.Get(m)
	return v.Invoke(args...)
}

func (f fakeValue) Invoke(args ...interface{}) h.Value {
	if reflect.TypeOf(f.value).Kind() != reflect.Func {
		return fakeValue{nil}
	}
	v := reflect.ValueOf(f.value)
	s := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		s[i] = reflect.ValueOf(args[i])
	}
	return fakeValue{v.Call(s)[0].Interface()}
}

func (f fakeValue) New(args ...interface{}) h.Value {
	return New(h.Object{})
}

func (f fakeValue) Float() float64 {
	v, ok := f.value.(float64)
	if !ok {
		return 0
	}
	return v
}

func (f fakeValue) Int() int {
	v, ok := f.value.(int)
	if !ok {
		return 0
	}
	return v
}

func (f fakeValue) Bool() bool {
	v, ok := f.value.(bool)
	if !ok {
		return false
	}
	return v
}

func (f fakeValue) String() string {
	v, ok := f.value.(string)
	if !ok {
		return ""
	}
	return v
}

func (f fakeValue) InstanceOf(t h.Value) bool {
	// TODO: implement
	return false
}
