package counter

import (
	"strconv"

	h "github.com/syumai/go-hyperscript/hyperscript"
)

type Counter struct {
	h.ComponentBase
}

func (c *Counter) InitialState() h.Object {
	return h.Object{
		"count": 0,
	}
}

func (c *Counter) Render() h.VNode {
	return h.H("div", h.Object{"className": "counter"},
		h.H("div", nil, h.Text(strconv.Itoa(c.State.Int("count")))),
		h.H("div", nil,
			h.H("button", h.Object{"onclick": h.Callback(func([]h.Value) { c.increment() })}, h.Text("+")),
			h.H("button", h.Object{"onclick": h.Callback(func([]h.Value) { c.decrement() })}, h.Text("-")),
		),
	)
}

func (c *Counter) increment() {
	c.SetState(h.Object{
		"count": c.State.Int("count") + 1,
	})
}

func (c *Counter) decrement() {
	c.SetState(h.Object{
		"count": c.State.Int("count") - 1,
	})
}
