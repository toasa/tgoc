package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	input := os.Args[1]
	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".globl _main\n")
	fmt.Printf("_main:\n")
	fmt.Printf("	push rbp\n")
	fmt.Printf("	mov rbp, rsp\n")
	fmt.Printf("	mov rax, %s\n", input)
	fmt.Printf("	pop rbp\n")
	fmt.Printf("	ret\n")
}
