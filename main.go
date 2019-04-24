package main

import (
	"fmt"
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

	// printTokens(l)

	p := parser.New(l.Tokens)
	stmts := p.Parse()

	x86.Gen(stmts, len(p.VarMap))
}

func printTokens(l *lexer.Lexer) {
	for _, tok := range l.Tokens {
		fmt.Printf("%+v\n", tok)
	}
}
