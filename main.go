package main

import (
	//"bytes"
	"bufio"
	"fmt"
	"os"
	"pine/compiler"
	. "pine/vm"
)

func main() {
	var v VM;
	v = v.InitVM()
	args := os.Args[1:]
	fmt.Println("args:", args)
	if len(args) > 1 {
		fmt.Println("Unexpected number of args, usage pine [script]")
	} else if len(args) == 1 {
		runFile(args[0], v)
	} else {
		runPrompt(v)
	}
	v.FreeVM()
}

func runFile(filename string, v VM) {
	fmt.Println(filename)
	file, err := os.ReadFile(filename)
	code := string(file)
	if err != nil {
		fmt.Errorf("Error reading file")
	}
	runCode(code, v)
}

func runPrompt(v VM) {
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
		runCode(input, v)
		fmt.Println(input)
	}

}

func runCode(code string, v VM) {
	compiler.Interpret(code, v)
}
