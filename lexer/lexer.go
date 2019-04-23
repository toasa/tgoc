package lexer

import (
	"tgoc/token"
)

// Lexer type
type Lexer struct {
	Tokens []token.Token
	Pos    int
	Input  string
}

var keywords []string = []string{
	"break", "default", "func",
	"interface", "select",
	"case", "defer", "go", "map",
	"struct", "chan", "else",
	"goto", "package", "switch",
	"const", "fallthrough", "if",
	"range", "type", "continue",
	"for", "import", "return", "var",
}

// New lexer create
func New(input string) *Lexer {
	t := []token.Token{}
	return &Lexer{Tokens: t, Input: input, Pos: 0}
}

// Analyze the input string ans split the token sequences
func (l *Lexer) Analyze() {

	var tok token.Token

	for ; l.Pos < len(l.Input); l.Pos++ {
		l.skip()

		switch l.Input[l.Pos] {
		case '+':
			tok = token.New(token.ADD, "+")
		case '-':
			tok = token.New(token.SUB, "-")
		case '*':
			tok = token.New(token.MUL, "*")
		case '/':
			tok = token.New(token.DIV, "/")
		case '%':
			tok = token.New(token.REM, "%")
		case '(':
			tok = token.New(token.LPAREN, "(")
		case ')':
			tok = token.New(token.RPAREN, ")")
		case ':':
			if l.Input[l.Pos+1] == '=' {
				l.Pos++
				tok = token.New(token.SVDECL, ":=")
			}
		case '\000':
			tok = token.New(token.EOF, "")
			l.Tokens = append(l.Tokens, tok)
			return
		default:
			if isDigit(l.Input[l.Pos]) {
				tok = token.New(token.INT, l.readDigit())
			} else if isChar(l.Input[l.Pos]) {
				str := l.readIdent()
				tok = token.New(token.IDENT, str)
			}
		}

		l.Tokens = append(l.Tokens, tok)
	}
}

func isDigit(c byte) bool {
	return ('0' <= c) && (c <= '9')
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func isChar(c byte) bool {
	return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') || (c == '_')
}

func (l *Lexer) skip() {
	for isSpace(l.Input[l.Pos]) {
		l.Pos++
	}
}

func (l *Lexer) readDigit() string {
	head := l.Pos
	tail := l.Pos + 1
	for ; tail < len(l.Input); tail++ {
		if !isDigit(l.Input[tail]) {
			break
		}
	}

	l.Pos = tail - 1
	return l.Input[head:tail]
}

func (l *Lexer) readIdent() string {
	head := l.Pos
	tail := l.Pos + 1
	for ; tail < len(l.Input); tail++ {
		if !isChar(l.Input[tail]) {
			break
		}
	}
	l.Pos = tail - 1
	return l.Input[head:tail]
}
