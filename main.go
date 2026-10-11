package main

import (
	//"bytes"
	"bufio"
	"fmt"
	"os"
	. "pine/vm"
)

func main() {
	var v VM;
	v = v.InitVM();
	chunk := InitChunk(5)
	index := AddConstant(chunk, 12)
	WriteChunk(chunk, OP_CONSTANT)
	WriteChunk(chunk, index);
	WriteChunk(chunk, OP_RETURN)
	v.Interpret(chunk)
	DissasembleChunk(chunk, "test chunk")
	FreeChunk(chunk)
	v.FreeVM()
	
	args := os.Args[1:]
	fmt.Println("args:", args)
	if len(args) > 1 {
		fmt.Println("Unexpected number of args, usage pine [script]")
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(filename string) {
	fmt.Println(filename)
	file, err := os.ReadFile(filename)
	code := string(file)
	if err != nil {
		fmt.Errorf("Error reading file")
	}
	runCode(code)
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if input == "" {
			break
		}
		runCode(input)
		fmt.Println(input)
	}

}

func runCode(code string) {

}
