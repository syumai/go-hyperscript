package state

type State struct {
	ToDos              ToDos
	Title              string
	afterSetStateHooks []func()
}

type (
	ToDos []ToDo
)

func NewState() *State {
	return &State{}
}

func (s *State) Subscribe(f func()) {
	s.afterSetStateHooks = append(s.afterSetStateHooks, f)
}

func (s *State) afterSetState() {
	for _, f := range s.afterSetStateHooks {
		f()
	}
}

func (s *State) SetTitle(title string) {
	s.Title = title
	s.afterSetState()
}

func (s *State) AppendToDo(toDo ToDo) {
	s.ToDos = append(s.ToDos, toDo)
	s.afterSetState()
}

func (s *State) RemoveToDo(index int) {
	if index >= len(s.ToDos) {
		return
	}
	var toDos ToDos
	for i, toDo := range s.ToDos {
		if i == index {
			continue
		}
		toDos = append(toDos, toDo)
	}
	s.ToDos = toDos
	s.afterSetState()
}

func (s *State) SetToDoDone(index int, done bool) {
	if index >= len(s.ToDos) {
		return
	}
	s.ToDos[index].Done = done
	s.afterSetState()
}
