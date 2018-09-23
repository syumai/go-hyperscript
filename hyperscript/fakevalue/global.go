package fakevalue

import (
	h "github.com/syumai/go-hyperscript/hyperscript"
)

var global = New(h.Object{
	"document": New(h.Object{
		"body": New(h.Object{}),
	}),
})

func Global() h.Value {
	return global
}
