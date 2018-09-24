package hyperscript

type Component interface {
	SetState(newState Object)
	Initialize(initialState Object, props Object, updatedHandler func())
	// InitialState must be defined by user.
	InitialState() Object
	// Render must be defined by user.
	Render() VNode
}

type ComponentBase struct {
	State          Object
	Props          Object
	updatedHandler func()
}

func (cb *ComponentBase) setUpdatedHandler(handler func()) {
	cb.updatedHandler = handler
}

func (cb *ComponentBase) Initialize(initialState Object, props Object, updatedHandler func()) {
	cb.Props = Object{}
	for k, v := range props {
		cb.Props[k] = v
	}
	cb.updatedHandler = updatedHandler
}

func (cb *ComponentBase) SetState(newState Object) {
	if cb.State == nil {
		cb.State = Object{}
	}
	for k, v := range newState {
		cb.State[k] = v
	}
	if cb.updatedHandler != nil {
		cb.updatedHandler()
	}
}
