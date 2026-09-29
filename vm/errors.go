package vm

import (
	"strconv"
)

type vmError uint8

const (
	_                    vmError = iota
	ErrorExaustedBuf
)

func (err vmError) Error() string {
	return "VM ERROR " + strconv.Itoa(int(err))
}