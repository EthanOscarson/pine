package main

import (
	//"bytes"
	"bufio"
	"fmt"
	"os"
)

func main() {
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
		if scanner.Scan() {
			input := scanner.Text()
			if input == "" {
				break
			}
			runCode(input)
			fmt.Println(input)
		}

	}

}

func runCode(code string) {

}
