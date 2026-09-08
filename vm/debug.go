package vm

import (
	"fmt"
)

func DissasembleChunk (chunk *Chunk, name string) {
	fmt.Printf("== %s ==\n", name)

	for offset := 0; offset < len(chunk.data); {
		offset = dissasembleInstruction(chunk, offset)
	}
	
}

func dissasembleInstruction (chunk *Chunk, offset int) int {
	fmt.Printf("%04d", offset)
	instruction := chunk.data[offset]
	switch instruction {
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