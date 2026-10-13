package vm

import "fmt"


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
			for i := range vm.stackTop {
				print("[ ")
				printValue(vm.stack[i])
				print(" ]")
				print("\n")
			}
			
			dissasembleInstruction(vm.chunk, int(ip))
		}
		ip++

		switch (instruction) {
			case OP_CONSTANT: {
				fmt.Printf("A%d", ip)
				constant := readByte(code, &ip)
				fmt.Printf("CONST%d", constant)
				value := vm.chunk.constants.values[constant]
				vm.push(value)
				printValue(value)
				print("\n")
				break
			}
			case OP_ADD: {
				vm.push(vm.pop() + vm.pop())
				break
			}
			case OP_SUBTRACT: {
				vm.push(vm.pop() - vm.pop())
				break
			}
			case OP_MULTIPLY: {
				vm.push(vm.pop() * vm.pop())
				break
			}
			case OP_DIVIDE: {
				vm.push(vm.pop() / vm.pop())
				break
			}
			
			case OP_NEGATE: {
				vm.push(-vm.pop());
				break
			}
			case OP_RETURN: {
				printValue(vm.pop());
				print("\n")
				return INTERPRET_OK;
			}
			
		}
	}
}

func readByte (code []uint8, ip *uint8) uint8 {
	a := code[*ip]
	*ip++
	return a
} 