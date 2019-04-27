package main

import (
	"fmt"
	"os"
	"tgoc/lexer"
	"tgoc/parser"
	"tgoc/x86"
)

func main() {
	var printTokenFlg bool
	var printParseFlg bool
	var input string

	argc := len(os.Args)
	if argc != 2 {
		if argc == 3 {
			input = os.Args[2] + "\000"
			switch os.Args[1] {
			case "-t":
				printTokenFlg = true
			case "-p":
				printParseFlg = true
			}
		} else {
			fmt.Println("[USAGE] go run main.go INPUT")
			os.Exit(1)
		}
	} else {
		input = os.Args[1] + "\000"
	}

	l := lexer.New(input)
	l.Analyze()

	if printTokenFlg {
		printTokens(l)
	}

	p := parser.New(l.Tokens)
	stmts := p.Parse()

	// not yet
	if printParseFlg {
	}

	x86.Gen(stmts, len(p.VarMap))
}

func printTokens(l *lexer.Lexer) {
	for _, tok := range l.Tokens {
		fmt.Printf("%+v\n", tok)
	}
}
