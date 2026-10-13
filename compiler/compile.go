package compiler

import "fmt"

type CompileResult = uint8;

const (
	COMPILE_OK             CompileResult = iota
	COMPILE_ERROR
)
// Compile takes the source code and returns the bytecode
func Compile (source string) (CompileResult, []uint8) {
	s := newScanner(source)
	line := -1;
	for {
		token := s.scanToken();
		if (token.line != line) {
			fmt.Printf("%d ", token.line)
			line = token.line
		} else {
			fmt.Printf("   | ")
		}
		fmt.Printf("%2d '%.*s'\n", token.type, token.length, token.start)
		if (token.type == TOKEN_EOF) {
			break;
		}
	} 
}