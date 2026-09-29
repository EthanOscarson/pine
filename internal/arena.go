package internal

import (
	"unsafe"
)
// provides a small arena implementation for memory allocation in pine

type Arena struct {
	buf []byte
	offset uintptr
}

func New (size int) *Arena {
	a := &Arena{}
	a.Init(make([]byte, (size+7)/8))
	return a;
}

func (a *Arena) Init(buf []byte) {
	a.offset = 0
	if len(buf) == 0 {
		a.buf = nil
		return
	}
	a.buf = buf
}

func (a *Arena) Alloc(buf []byte) {
	
}

