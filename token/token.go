package token

// TokenType
type TokenType int

// tokentype
const (
	TKINT   = iota // 20, 255
	TKPLUS         // '+'
	TKMINUS        // '-'
	TKEOF          // EOF
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
