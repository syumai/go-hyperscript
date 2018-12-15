package main

import (
	"syscall/js"
	"time"

	"github.com/syumai/go-hyperscript/dom"
	"github.com/syumai/go-hyperscript/examples/counter/counter"
	h "github.com/syumai/go-hyperscript/hyperscript"
)

var (
	body     = js.Global().Get("document").Get("body")
	renderer = dom.NewRenderer()
)

func app(h.Object) h.VNode {
	return h.H("div", nil,
		h.H("h1", nil, h.Text("Counter")),
		h.H(counter.Counter, nil),
		h.H("a", h.Object{"href": "https://github.com/syumai/go-hyperscript/"},
			h.Text("Show the code on GitHub"),
		),
	)
}

func main() {
	t := time.NewTicker(time.Millisecond)
	for {
		select {
		case <-t.C:
			renderer.Render(
				h.H(app, nil),
				body,
			)
		}
	}
}
