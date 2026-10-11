package vm


const (
	INTERPRET_OK             InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

type VM struct {
	chunk *Chunk
	ip uint8
	stack [256]Value
	stackTop uint8
}

type InterpretResult = uint8;





func (vm *VM) InitVM () VM {
	return VM{}
}

func (vm *VM) FreeVM () {
	
}

func (vm *VM) Interpret (chunk *Chunk) InterpretResult {
	vm.chunk = chunk
	vm.ip = 0
	return vm.Run();
}

func (vm *VM) Run () InterpretResult {
	code := vm.chunk.code
	ip := vm.ip
	var instruction uint8
	for {
		instruction = code[ip]
		if DEBUG {
			dissasembleInstruction(vm.chunk, int(ip))
		}
		ip++
		switch (instruction) {
			case OP_CONSTANT: {
				constant := readByte(code, &ip)
				value := vm.chunk.constants.values[constant]
				printValue(value)
				print("\n")
			}
			case OP_RETURN: {
				return INTERPRET_OK;
			}
			
		}
	}
}

func readByte (code []uint8, ip *uint8) uint8 {
	*ip++
	return code[*ip]
} 