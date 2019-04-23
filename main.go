package main

import (
	"os"
	"tgoc/lexer"
	"tgoc/parser"
	"tgoc/x86"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	input := os.Args[1] + "\000"

	l := lexer.New(input)
	l.Analyze()

	// for _, tok := range l.Tokens {
	// 	fmt.Printf("%+v\n", tok)
	// }

	p := parser.New(l.Tokens)
	node := p.Parse()

	x86.Gen(node)
}
