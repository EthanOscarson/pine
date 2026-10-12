package vm

import "fmt"

type OpCode = uint8

const (
	OP_RETURN OpCode = iota
	OP_CONSTANT
	OP_NEGATE
)

// Chunk.code is dynamic currently, should be refactored afterwards as we should know how much code & constants are in a chunk after the preprocess step

type Chunk struct {
	code []uint8
 	constants ValueArray
}
func InitChunk(cap int) *Chunk {
	c := &Chunk{
		code: make([]uint8, 0, cap),
	}
	c.constants.init(make([]int64, 0, 255))
	return c
}
func WriteChunk(chunk *Chunk, chunkByte uint8) {
	if chunk.code == nil {
		panic("chunk is not allocated")
	}
	chunk.code = append(chunk.code, chunkByte)
}
func AddConstant(chunk *Chunk, value Value) uint8 {
	err := chunk.constants.write(value)
	if err != nil {
		panic("Constant buffer exausted")
	}
	fmt.Printf("DEBUG%d", len(chunk.constants.values) -1 )
	return uint8(len(chunk.constants.values)-1);
	// temporary panic, we should know constant buffer limit before compilation
}
func FreeChunk(chunk *Chunk) {
	if chunk.code == nil {
		panic("chunk is not allocated")
	}
	chunk.code = nil
	chunk.constants.free()
}
