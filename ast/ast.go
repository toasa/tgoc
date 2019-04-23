package ast

import (
	"bytes"
	"strconv"
)

type Node interface {
	String() string
}

type Stmt interface {
	Node
	stmtNode()
}

type ExprStmt struct {
	Expr Expr
}

func (es *ExprStmt) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(es.Expr.String())
	out.WriteString(")")

	return out.String()
}
func (es *ExprStmt) stmtNode() {}

type Expr interface {
	Node
	exprNode()
}

type BinaryExpr struct {
	Op  string
	Lhs Expr
	Rhs Expr
}

func (be *BinaryExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(be.Lhs.String())
	out.WriteString(be.Op)
	out.WriteString(be.Rhs.String())
	out.WriteString(")")

	return out.String()
}
func (be *BinaryExpr) exprNode() {}

type IntLit struct {
	Val int
}

func (il *IntLit) String() string {
	return strconv.Itoa(il.Val)
}
func (il *IntLit) exprNode() {}
