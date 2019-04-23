package ast

type Node interface {
	TokenLiteral() string
}

type Stmt interface {
	Node
	stmtNode()
}

type ExprStmt struct {
	Expr Expr
}

func (es *ExprStmt) TokenLiteral() string { return "" }
func (es *ExprStmt) stmtNode()            {}

type Expr interface {
	Node
	exprNode()
}

type BinaryExpr struct {
	Op  string
	Lhs Expr
	Rhs Expr
}

func (be *BinaryExpr) TokenLiteral() string { return "" }
func (be *BinaryExpr) exprNode()            {}

type IntLit struct {
	Val int
}

func (il *IntLit) TokenLiteral() string { return "" }
func (il *IntLit) exprNode()            {}
