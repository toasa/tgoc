package ast

import (
	"bytes"
	"fmt"
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

type DeclStmt struct {
	Decl Decl
}

func (ds *DeclStmt) String() string {
	return ds.Decl.String()
}
func (ds *DeclStmt) stmtNode() {}

type AssignStmt struct {
	Name string
	Val  Expr
}

func (as *AssignStmt) String() string {
	return ""
}

func (as *AssignStmt) stmtNode() {}

type ReturnStmt struct {
	Expr Expr
}

func (rs *ReturnStmt) String() string {
	return "return " + rs.Expr.String()
}
func (rs *ReturnStmt) stmtNode() {}

type IfStmt struct {
	Cond Expr
	Cons []Stmt
	Alt  []Stmt
}

func (is *IfStmt) String() string {
	var out bytes.Buffer

	out.WriteString("if ")
	out.WriteString(is.Cond.String())
	out.WriteString("{")
	for _, stmt := range is.Cons {
		out.WriteString(stmt.String())
	}
	out.WriteString("}")

	if is.Alt != nil {
		out.WriteString("{")
		for _, stmt := range is.Alt {
			out.WriteString(stmt.String())
		}
		out.WriteString("}")
	}

	return out.String()
}
func (is *IfStmt) stmtNode() {}

type ForSingleStmt struct {
	Cond  Expr
	Stmts []Stmt
}

func (fss *ForSingleStmt) String() string {
	var out bytes.Buffer

	out.WriteString("for ")
	out.WriteString(fss.Cond.String())
	out.WriteString("{")
	for _, stmt := range fss.Stmts {
		out.WriteString(stmt.String())
	}
	out.WriteString("}")

	return out.String()
}
func (fss *ForSingleStmt) stmtNode() {}

type ForClauseStmt struct {
	Init  Stmt
	Cond  Expr
	Post  Stmt
	Stmts []Stmt
}

func (fc *ForClauseStmt) String() string {
	var out bytes.Buffer

	out.WriteString("for ")
	out.WriteString(fc.Init.String())
	out.WriteString("; ")
	out.WriteString(fc.Cond.String())
	out.WriteString("; ")
	out.WriteString(fc.Post.String())
	out.WriteString("{")
	for _, stmt := range fc.Stmts {
		out.WriteString(stmt.String())
	}
	out.WriteString("}")

	return out.String()
}
func (fc *ForClauseStmt) stmtNode() {}

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

type UnaryExpr struct {
	Op   string
	Expr Expr
}

func (ue *UnaryExpr) String() string {
	return ue.Op + ue.Expr.String()
}
func (ue *UnaryExpr) exprNode() {}

type PtrExpr struct {
	Of   *PtrExpr
	Expr Expr
}

func (pe *PtrExpr) String() string {
	return ""
}
func (pe *PtrExpr) exprNode() {}

type LogicalExpr struct {
	Op  string
	Lhs Expr
	Rhs Expr
}

func (le *LogicalExpr) String() string {
	return fmt.Sprintln(le.Lhs.String(), le.Op, le.Rhs.String())
}
func (le *LogicalExpr) exprNode() {}

type IntLit struct {
	Val int
}

func (il *IntLit) String() string {
	return strconv.Itoa(il.Val)
}
func (il *IntLit) exprNode() {}

type Boolean struct {
	Val bool
}

func (b *Boolean) String() string {
	if b.Val {
		return "true"
	}
	return "false"
}
func (b *Boolean) exprNode() {}

type Ident struct {
	Name string
	Val  Expr
	Type Type
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

type VarDecl struct {
	Name string
}

func (vd *VarDecl) String() string {
	var out bytes.Buffer

	out.WriteString("var ")
	out.WriteString(vd.Name)

	return out.String()
}
func (vd *VarDecl) declNode() {}

// --------------------------------------------------------
// - Type
// --------------------------------------------------------

type Type struct {
	Val   string // TINT or TPTR
	PtrOf *Type
}
