package parser

import (
	"fmt"
	"strconv"
	"strings"
	"tgoc/ast"
	"tgoc/token"
	"tgoc/utils"
)

type Parser struct {
	Tokens []token.Token
	Pos    int
	VarMap map[string]*ast.Ident
	Stmts  []ast.Stmt
}

func New(t []token.Token) *Parser {
	return &Parser{Tokens: t, Pos: 0, VarMap: map[string]*ast.Ident{}, Stmts: []ast.Stmt{}}
}

func (p *Parser) parseTerm() ast.Expr {
	utils.Assert(p.curTokenIs(token.INT) || p.curTokenIs(token.LPAREN), fmt.Sprintf("invalid token: %s", p.curToken().Literal))

	if p.curTokenIs(token.INT) {
		n, _ := strconv.Atoi(p.Tokens[p.Pos].Literal)
		p.nextToken()
		return &ast.IntLit{Val: n}
	}
	if p.curTokenIs(token.LPAREN) {
		p.nextToken()
		node := p.parseAdd()
		utils.Assert(p.curTokenIs(token.RPAREN), fmt.Sprintf("expected RPAREN, but got %s", p.curToken().Literal))
		p.nextToken()
		return node
	}

	return nil
}

func (p *Parser) parseIdent() ast.Expr {
	if p.curTokenIs(token.IDENT) {
		ident, ok := p.VarMap[p.curToken().Literal]
		utils.Assert(ok, fmt.Sprintf("undeclared identifier: %s", p.curToken().Literal))
		p.nextToken()
		return ident
	} else if p.curTokenIs(token.TRUE) {
		p.nextToken()
		return &ast.IntLit{Val: 1}
	} else if p.curTokenIs(token.FALSE) {
		p.nextToken()
		return &ast.IntLit{Val: 0}
	} else {
		return p.parseTerm()
	}
}

func (p *Parser) parseUnary() ast.Expr {
	var lhs ast.Expr
	if p.curTokenIs(token.SUB) {
		p.nextToken()
		lhs = &ast.UnaryExpr{Op: "-", Expr: p.parseIdent()}
	} else {
		if p.curTokenIs(token.ADD) {
			p.nextToken()
		}
		lhs = p.parseIdent()
	}
	return lhs
}

func (p *Parser) parseMul() ast.Expr {
	lhs := p.parseUnary()

	for p.curTokenIs(token.MUL) || p.curTokenIs(token.DIV) || p.curTokenIs(token.REM) ||
		p.curTokenIs(token.LSHIFT) || p.curTokenIs(token.RSHIFT) {

		op := p.curToken().Literal
		p.nextToken()
		rhs := p.parseUnary()
		lhs = &ast.BinaryExpr{Op: op, Lhs: lhs, Rhs: rhs}
	}

	return lhs
}

func (p *Parser) parseAdd() ast.Expr {
	lhs := p.parseMul()

	for p.curTokenIs(token.ADD) || p.curTokenIs(token.SUB) {
		op := p.curToken().Literal
		p.nextToken()
		rhs := p.parseMul()
		lhs = &ast.BinaryExpr{Op: op, Lhs: lhs, Rhs: rhs}
	}
	return lhs
}

func (p *Parser) parseComparison() ast.Expr {
	lhs := p.parseAdd()

	for p.curTokenIs(token.EQ) || p.curTokenIs(token.NQ) ||
		p.curTokenIs(token.LT) || p.curTokenIs(token.GT) {

		op := p.curToken().Literal
		p.nextToken()
		rhs := p.parseAdd()
		lhs = &ast.BinaryExpr{Op: op, Lhs: lhs, Rhs: rhs}
	}
	return lhs
}

func printTree(node ast.Expr, tab int) {
	be, ok := node.(*ast.BinaryExpr)
	if ok {
		printTree(be.Lhs, tab+4)
		fmt.Println(strings.Repeat(" ", tab), be.Op)
		printTree(be.Rhs, tab+4)
		return
	}

	il, ok := node.(*ast.IntLit)
	if ok {
		fmt.Println(strings.Repeat(" ", tab), il.Val)
	}
	return
}

func (p *Parser) parseExpr() ast.Expr {
	lhs := p.parseComparison()
	//printTree(lhs, 0)
	return lhs
}

func (p *Parser) parseExprStmt() ast.Stmt {
	expr := p.parseExpr()
	es := &ast.ExprStmt{Expr: expr}
	return es
}

func (p *Parser) parseDecl() ast.Decl {
	utils.Assert(p.curTokenIs(token.IDENT), "identifier needed")
	name := p.Tokens[p.Pos].Literal
	p.nextToken()
	p.nextToken()
	val := p.parseExpr()
	return &ast.SVDecl{Name: name, Val: val}
}

func (p *Parser) parseAssignStmt() ast.Stmt {
	decl := p.parseDecl()
	svd, ok := decl.(*ast.SVDecl)
	if ok {
		p.VarMap[svd.Name] = &ast.Ident{Name: svd.Name, Val: svd.Val}
	}

	as := &ast.AssignStmt{Decl: decl}
	return as
}

func (p *Parser) parseReturnStmt() ast.Stmt {
	p.nextToken()
	return &ast.ReturnStmt{Expr: p.parseExpr()}
}

func (p *Parser) parseStmt() ast.Stmt {
	var stmt ast.Stmt

	if p.curTokenIs(token.IDENT) && p.peepTokenIs(token.SVDECL) {
		stmt = p.parseAssignStmt()
	} else if p.curTokenIs(token.RETURN) {
		stmt = p.parseReturnStmt()
	} else {
		stmt = p.parseExprStmt()
	}

	if p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) Parse() []ast.Stmt {
	for !p.curTokenIs(token.EOF) {
		p.Stmts = append(p.Stmts, p.parseStmt())
	}
	return p.Stmts
}

func (p *Parser) curTokenIs(tt token.TokenType) bool {
	return tt == p.curToken().Type
}

func (p *Parser) peepTokenIs(tt token.TokenType) bool {
	return tt == p.peepToken().Type
}

func (p *Parser) curToken() token.Token {
	return p.Tokens[p.Pos]
}

func (p *Parser) peepToken() token.Token {
	return p.Tokens[p.Pos+1]
}

func (p *Parser) nextToken() {
	if p.curTokenIs(token.EOF) {
		return
	}
	p.Pos++
}
