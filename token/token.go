package token

// TokenType is assigned unique number
type TokenType string

// tokentype
const (
	INT   = "INT"   // 20, 255
	IDENT = "IDENT" // abc, toasa
	TRUE  = "TRUE"  // 'true'
	FALSE = "FALSE" // 'false'

	// Arithmetic operator
	ADD    = "ADD"    // '+'
	SUB    = "SUB"    // '-'
	MUL    = "MUL"    // '*'
	DIV    = "DIV"    // '/'
	REM    = "REM"    // '%'
	LSHIFT = "LSHIFT" // '<<'
	RSHIFT = "RSHIFT" // '>>'
	BAND   = "BAND"   // '&': bitwise and
	BOR    = "BOR"    // '|': bitwise or
	BXOR   = "BXOR"   // '^': bitwise xor
	BCLR   = "BCLR"   // '&^': bit clear

	// Yield boolean
	LT   = "LT"   // '<'
	GT   = "GT"   // '>'
	LTE  = "LTE"  // '<='
	GTE  = "GTE"  // '>='
	EQ   = "EQ"   // '=='
	NQ   = "NQ"   // '!='
	CAND = "CAND" // '&&': conditional and
	COR  = "COR"  // '||': conditional or
	NOT  = "NOT"  // '!'

	// Statement
	IF     = "IF"     // 'if'
	ELSE   = "ELSE"   // 'ELSE'
	RETURN = "RETURN" // 'return'
	FOR    = "FOR"    // 'for'

	// Type
	TINT = "TINT" // 'int'

	LPAREN = "LPAREN" // '('
	RPAREN = "RPAREN" // ')'
	LBRACE = "LBRACE" // '{'
	RBRACE = "RBRACE" // '}'

	ASSIGN = "ASSIGN" // '='

	// Declaration
	VAR    = "VAR"    // 'var'
	SVDECL = "SVDECL" // ':=' Short Var Declaration

	SEMICOLON = "SEMICOLON" // ';'
	EOF       = "EOF"       // End of file
)

// Token (minimum unit of Go code)
type Token struct {
	Type    TokenType
	Literal string
}

// New token
func New(t TokenType, lit string) Token {
	return Token{Type: t, Literal: lit}
}
