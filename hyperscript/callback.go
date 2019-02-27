package hyperscript

type Func struct {
	Func *func(this Value, args []Value) interface{}
}

func FuncOf(fn func(this Value, args []Value) interface{}) Func {
	return Func{
		Func: &fn,
	}
}

func (c Func) Call(this Value, args []Value) interface{} {
	return (*c.Func)(this, args)
}
