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
	EventCallbackFlgPreventDefault EventCallbackFlag = 1 << iota
	EventCallbackFlgStopPropagation
	EventCallbackFlgStopImmediatePropagation
)

func NewEventCallback(flags EventCallbackFlag, fn func(event Value)) EventCallback {
	return EventCallback{
		Flg:  flags,
		Func: fn,
	}
}

func IsCallback(v interface{}) bool {
	switch v.(type) {
	case Callback:
		return true
	case EventCallback:
		return true
	}
	return false
}
