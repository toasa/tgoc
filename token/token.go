package token

// TokenType is assigned unique number
type TokenType string

// tokentype
const (
	INT       = "INT"       // 20, 255
	ADD       = "ADD"       // '+'
	SUB       = "SUB"       // '-'
	MUL       = "MUL"       // '*'
	DIV       = "DIV"       // '/'
	REM       = "REM"       // '%'
	LSHIFT    = "LSHIFT"    // '<<'
	RSHIFT    = "RSHIFT"    // '>>'
	LT        = "LT"        // '<'
	GT        = "GT"        // '>'
	LTE       = "LTE"       // '<='
	GTE       = "GTE"       // '>='
	EQ        = "EQ"        // '=='
	NQ        = "NQ"        // '!='
	LPAREN    = "LPAREN"    // '('
	RPAREN    = "RPAREN"    // ')'
	IDENT     = "IDENT"     // abc, toasa
	TRUE      = "TRUE"      // 'true'
	FALSE     = "FALSE"     // 'false'
	SVDECL    = "SVDECL"    // ':=' Short Var Declaration
	SEMICOLON = "SEMICOLON" // ';'
	RETURN    = "RETURN"    // 'return'
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
