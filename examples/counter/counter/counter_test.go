package counter

import (
	"testing"
)

func newCounter() *Counter {
	c := &Counter{}
	c.Initialize(c.InitialState(), nil, nil)
	return c
}

func Test_increment(t *testing.T) {
	c := newCounter()
	c.increment()
	if c.State.Int("count") != 1 {
		t.Errorf("count must be 1")
	}
}

func Test_decrement(t *testing.T) {
	c := newCounter()
	c.decrement()
	if c.State.Int("count") != -1 {
		t.Errorf("count must be -1")
	}
}
