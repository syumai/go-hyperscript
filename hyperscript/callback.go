package hyperscript

type (
	Callback struct {
		Func *func(args []Value)
	}

	EventCallback struct {
		Flg  EventCallbackFlag
		Func *func(event Value)
	}
)

type EventCallbackFlag int

const (
	EventCallbackFlgPreventDefault EventCallbackFlag = 1 << iota
	EventCallbackFlgStopPropagation
	EventCallbackFlgStopImmediatePropagation
)

func NewCallback(fn func(args []Value)) Callback {
	return Callback{
		Func: &fn,
	}
}

func NewEventCallback(flags EventCallbackFlag, fn func(event Value)) EventCallback {
	return EventCallback{
		Flg:  flags,
		Func: &fn,
	}
}

func (c Callback) Call(args []Value) {
	(*c.Func)(args)
}

func (c EventCallback) Call(event Value) {
	(*c.Func)(event)
}
