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
	appendToDo = h.FuncOf(func(this h.Value, args []h.Value) interface{} {
		event := args[0]
		event.Call("preventDefault")
		s.AppendToDo(state.ToDo{
			Title: s.Title,
			Done:  false,
		})
		s.SetTitle("")
		return nil
	})

	setTitle = h.FuncOf(func(this h.Value, args []h.Value) interface{} {
		event := args[0]
		event.Call("preventDefault")
		s.SetTitle(event.Get("target").Get("value").String())
		return nil
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
			h.H("a", h.Object{"href": "https://github.com/syumai/go-hyperscript/"},
				h.Text("Show the code on GitHub"),
			),
		),
		body,
	)
}

func main() {
	s.Subscribe(render)
	render()
	select {}
}
