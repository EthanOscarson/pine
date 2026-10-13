package compiler

import (
	"pine/vm"
)

type InterpretResult = uint8;

const (
	INTERPRET_OK             InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

func Interpret (source string, v vm.VM) InterpretResult {
	result, bytecode := compile(source)
	v.Run(bytecode)
}