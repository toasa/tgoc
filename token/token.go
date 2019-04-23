package token

// TokenType is assigned unique number
type TokenType string

// tokentype
const (
	INT   = "INT"   // 20, 255
	PLUS  = "PLUS"  // '+'
	MINUS = "MINUS" // '-'
	MUL   = "MUL"   // '*'
	EOF   = "EOF"   // End of file
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
