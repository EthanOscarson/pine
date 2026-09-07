package vm

type OpCode uint8

const (
	OP_RETURN OpCode = iota
)

type Chunk struct {
	data []uint8
}
func initChunk(cap int) *Chunk {
	return &Chunk{
		data: make([]uint8, 0, cap),
	}
}
func writeChunk(chunk *Chunk, chunkByte uint8) {
	if chunk.data == nil {
		panic("chunk is not allocated")
	}
	chunk.data = append(chunk.data, chunkByte)
}
func freeChunk(chunk *Chunk) {
	if chunk.data == nil {
		panic("chunk is not allocated")
	}
	chunk.data = nil
}
