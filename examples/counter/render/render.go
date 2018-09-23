package render

import (
	h "github.com/syumai/go-hyperscript/hyperscript"
)

var (
	UpdateHandler func()
)

func Action(f func()) h.Callback {
	return func([]h.Value) { f(); Update() }
}

func Update() {
	UpdateHandler()
}
