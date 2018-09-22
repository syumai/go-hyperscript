package hyperscript

type (
	Callback func(args []Object)
	EventCallback struct {
		Flg EventCallbackFlag
		Func func(event Object)
	}
)

type EventCallbackFlag int

const (
	PreventDefault EventCallbackFlag = 1 << iota
	StopPropagation
	StopImmediatePropagation
)

func NewEventCallback(flags EventCallbackFlag, fn func(event Object)) EventCallback {
	return EventCallback{
		Flg: flags,
		Func: fn,
	}
}

