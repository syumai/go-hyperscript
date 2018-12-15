package main

import (
	"syscall/js"

	"github.com/syumai/go-hyperscript/dom"
	"github.com/syumai/go-hyperscript/examples/simpletodo/components"
	"github.com/syumai/go-hyperscript/examples/simpletodo/state"
	h "github.com/syumai/go-hyperscript/hyperscript"
)

var (
	s    = state.NewState()
	r    = dom.NewRenderer()
	body = js.Global().Get("document").Get("body")
)

var (
	appendToDo = h.NewEventCallback(h.EventCallbackFlgPreventDefault, func(h.Value) {
		s.AppendToDo(state.ToDo{
			Title: s.Title,
			Done:  false,
		})
		s.SetTitle("")
	})

	setTitle = h.NewEventCallback(h.EventCallbackFlgPreventDefault, func(event h.Value) {
		s.SetTitle(event.Get("target").Get("value").String())
	})
)

func render() {
	r.Render(
		h.H("div", nil,
			h.H("h2", nil, h.Text("Simple ToDo list example")),
			h.H("form", h.Object{"onsubmit": appendToDo},
				h.H("input", h.Object{"type": "text", "value": s.Title, "oninput": setTitle}),
				h.H("button", nil, h.Text("Add")),
			),
			h.H(components.ToDo, h.Object{
				"todos":       s.ToDos,
				"setToDoDone": s.SetToDoDone,
				"removeToDo":  s.RemoveToDo,
			}),
		),
		body,
	)
}

func main() {
	s.Subscribe(render)
	render()
	select {}
}
