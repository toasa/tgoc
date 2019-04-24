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

type Expr interface {
	Node
	exprNode()
}

type Decl interface {
	Node
	declNode()
}

// --------------------------------------------------------
// - Statement
// --------------------------------------------------------
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

type AssignStmt struct {
	Decl Decl
}

func (as *AssignStmt) String() string {
	return as.Decl.String()
}
func (as *AssignStmt) stmtNode() {}

type ReturnStmt struct {
	Expr Expr
}

func (rs *ReturnStmt) String() string {
	return "return " + rs.Expr.String()
}
func (rs *ReturnStmt) stmtNode() {}

// --------------------------------------------------------
// - Expression
// --------------------------------------------------------

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

type Ident struct {
	Name string
	Val  Expr
}

func (id *Ident) String() string {
	return id.Name
}
func (id *Ident) exprNode() {}

// --------------------------------------------------------
// - Declaration
// --------------------------------------------------------
type SVDecl struct {
	Name string
	Val  Expr
}

func (svd *SVDecl) String() string {
	var out bytes.Buffer

	out.WriteString(svd.Name + " := ")
	out.WriteString(svd.Val.String())

	return out.String()
}
func (svd *SVDecl) declNode() {}
