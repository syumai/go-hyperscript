package components

import (
	"github.com/syumai/go-hyperscript/examples/simpletodo/state"
	"github.com/syumai/go-hyperscript/examples/simpletodo/style"
	h "github.com/syumai/go-hyperscript/hyperscript"
)

func ToDo(props h.Object) h.VNode {
	setToDoDone := props.Get("setToDoDone").(func(int, bool))
	removeToDo := props.Get("removeToDo").(func(int))
	toDos := props.Get("todos").(state.ToDos)

	var toDoNodes h.VNodes
	for i, toDo := range toDos {
		i, toDo := i, toDo
		onCheck := h.Callback(func([]h.Value) { setToDoDone(i, !toDo.Done) })
		onClickRemove := h.Callback(func([]h.Value) { removeToDo(i) })
		toDoNodes = append(toDoNodes,
			h.H("div", h.Object{"style": style.Style{"display": "flex", "align-items": "center"}.String()},
				h.H("input", h.Object{"type": "checkbox", "checked": toDo.Done, "onchange": onCheck}),
				h.H("div", nil, h.Text(toDo.Title)),
				h.H("button", h.Object{"onclick": onClickRemove}, h.Text("Remove")),
			),
		)
	}

	return h.H("div", nil,
		h.H("h3", nil, h.Text("ToDos")),
		h.H("div", nil, toDoNodes...),
	)
}
