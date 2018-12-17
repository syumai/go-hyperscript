package hyperscript

import (
	"unsafe"
)

type pointerGetter struct {
	_   unsafe.Pointer
	ptr unsafe.Pointer
}
