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
// Interpret takes the source code and a vm, compiles the source and runs the code
func Interpret (source string, v vm.VM) InterpretResult {
	result, bytecode := compile(source)
	err := v.Run(bytecode)
}