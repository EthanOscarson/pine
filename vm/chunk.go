package vm

type OpCode = uint8

const (
	OP_RETURN OpCode = iota
)

type Chunk struct {
	data []uint8
}
func InitChunk(cap int) *Chunk {
	return &Chunk{
		data: make([]uint8, 0, cap),
	}
}
func WriteChunk(chunk *Chunk, chunkByte uint8) {
	if chunk.data == nil {
		panic("chunk is not allocated")
	}
	chunk.data = append(chunk.data, chunkByte)
}
func FreeChunk(chunk *Chunk) {
	if chunk.data == nil {
		panic("chunk is not allocated")
	}
	chunk.data = nil
}
