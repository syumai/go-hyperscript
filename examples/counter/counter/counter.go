package counter

import (
	"strconv"

	h "github.com/syumai/go-hyperscript/hyperscript"
)

type State struct {
	count int
}

var state = State{}

var (
	increment = h.FuncOf(func(h.Value, []h.Value) interface{} {
		state.count++
		return nil
	})

	decrement = h.FuncOf(func(h.Value, []h.Value) interface{} {
		state.count--
		return nil
	})
)

func Counter(h.Object) h.VNode {
	return h.H("div", h.Object{"className": "counter"},
		h.H("div", nil, h.Text(strconv.Itoa(state.count))),
		h.H("div", nil,
			h.H("button", h.Object{"onclick": increment}, h.Text("+")),
			h.H("button", h.Object{"onclick": decrement}, h.Text("-")),
		),
	)
}
