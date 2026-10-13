package compiler

type CompileResult = uint8;

const (
	COMPILE_OK             CompileResult = iota
	COMPILE_ERROR
)

func Compile (source string) (CompileResult, []uint8) {
	
}