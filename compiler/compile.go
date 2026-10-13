package compiler

type CompileResult = uint8;

const (
	COMPILE_OK             CompileResult = iota
	COMPILE_ERROR
)
// Compile takes the source code and returns the bytecode
func Compile (source string) (CompileResult, []uint8) {
	
}