package vm

func (vm *VM) resetStack () {
	vm.stackTop = 0
}

func (vm *VM) push (value Value) {
	vm.stack[vm.stackTop] = value
	vm.stackTop++
}

func (vm *VM) pop () Value {
	vm.stackTop--
	return vm.stack[vm.stackTop]
}



