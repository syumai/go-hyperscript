package counter

import (
	"strconv"

	"github.com/syumai/go-hyperscript/examples/counter/render"
	h "github.com/syumai/go-hyperscript/hyperscript"
)

type State struct {
	count int
}

var state = State{}

func increment() {
	state.count++
}

func decrement() {
	state.count--
}

func Counter(h.Object) h.VNode {
	return h.H("div", h.Object{"className": "counter"},
		h.H("div", nil, h.Text(strconv.Itoa(state.count))),
		h.H("div", nil,
			h.H("button", h.Object{"onclick": render.Action(increment)}, h.Text("+")),
			h.H("button", h.Object{"onclick": render.Action(decrement)}, h.Text("-")),
		),
	)
}
