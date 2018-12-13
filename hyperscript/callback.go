package hyperscript

type (
	Callback      func(args []Value)
	EventCallback struct {
		Flg  EventCallbackFlag
		Func func(event Value)
	}
)

type EventCallbackFlag int

const (
	PreventDefault EventCallbackFlag = 1 << iota
	StopPropagation
	StopImmediatePropagation
)

func NewEventCallback(flags EventCallbackFlag, fn func(event Value)) EventCallback {
	return EventCallback{
		Flg:  flags,
		Func: fn,
	}
}

func IsCallback(v interface{}) bool {
	_, isCallback := v.(Callback)
	return isCallback
}
