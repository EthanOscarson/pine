package vm


import (
	"fmt"
)

const DEBUG = true;


func DissasembleChunk (chunk *Chunk, name string) {
	fmt.Printf("== %s ==\n", name)

	for offset := 0; offset < len(chunk.code); {
		offset = dissasembleInstruction(chunk, offset)
	}
	
}

func dissasembleInstruction (chunk *Chunk, offset int) int {
	fmt.Printf("%04d", offset)
	instruction := chunk.code[offset]
	switch instruction {
		case OP_CONSTANT:
			return constantInstruction("OP_CONSTANT", chunk, offset)
		case OP_NEGATE:
			return simpleInstruction("OP_NEGATE", offset)
		case OP_RETURN:
			return simpleInstruction("OP_RETURN", offset)
		default:
			print("Unknown opcode %d", instruction)
			return offset + 1
	}
}

func simpleInstruction(name string, offset int) int {
	fmt.Printf(" %s\n", name)
	return offset + 1
}

func printValue(value Value) {
	fmt.Printf("%d", value)
}

func constantInstruction(name string, chunk *Chunk, offset int) int {
	constant := chunk.code[offset +1]
	fmt.Printf(" %-16s %4d=", name, constant)
	printValue(chunk.constants.values[constant])
	fmt.Printf("\n")
	return offset + 2
}