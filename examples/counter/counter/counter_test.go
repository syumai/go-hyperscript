package counter

import "testing"

func Test_increment(t *testing.T) {
	state = State{}
	increment(nil)
	if state.count != 1 {
		t.Errorf("count must be 1")
	}
}

func Test_decrement(t *testing.T) {
	state = State{}
	decrement(nil)
	if state.count != -1 {
		t.Errorf("count must be -1")
	}
}
