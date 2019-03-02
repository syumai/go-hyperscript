package hyperscript

type Value interface {
	Get(p string) Value
	Set(p string, x interface{})
	Index(i int) Value
	SetIndex(i int, x interface{})
	Length() int
	Call(m string, args ...interface{}) Value
	Invoke(args ...interface{}) Value
	New(args ...interface{}) Value
	Float() float64
	Int() int
	Bool() bool
	Truthy() bool
	String() string
	InstanceOf(t Value) bool
}
