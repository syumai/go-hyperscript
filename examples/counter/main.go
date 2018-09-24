package main

import (
	"syscall/js"

	"github.com/syumai/go-hyperscript/dom"
	"github.com/syumai/go-hyperscript/examples/counter/counter"
	"github.com/syumai/go-hyperscript/examples/counter/render"
	h "github.com/syumai/go-hyperscript/hyperscript"
)

var app = func(h.Object) h.VNode {
	return h.H("div", nil,
		h.H("h1", nil, h.Text("Counter")),
		h.H(counter.Counter{}, nil),
		h.H("a", h.Object{"href": "https://github.com/syumai/go-hyperscript/"},
			h.Text("Show the code on GitHub"),
		),
	)
}

var body = js.Global().Get("document").Get("body")

func main() {
	render.UpdateHandler = func() {
		body.Set("innerHTML", "")
		dom.Render(h.H(app, nil), body)
	}
	render.Update()
	select {}
}
