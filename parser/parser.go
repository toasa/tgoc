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
}

func New(t []token.Token) *Parser {
	return &Parser{Tokens: t, Pos: 0}
}

func (p *Parser) parseTerm() ast.Expr {
	if p.curTokenIs(token.EOF) {
		return nil
	}

	utils.Assert(p.curTokenIs(token.INT), "invalid token")
	n, _ := strconv.Atoi(p.Tokens[p.Pos].Literal)
	p.nextToken()
	return &ast.IntLit{Val: n}
}

func (p *Parser) parseMul() ast.Expr {
	lhs := p.parseTerm()
	for p.curTokenIs(token.MUL) {
		p.nextToken()
		rhs := p.parseTerm()
		lhs = &ast.BinaryExpr{Op: "*", Lhs: lhs, Rhs: rhs}
	}
	return lhs
}

func (p *Parser) parseAdd() ast.Expr {
	lhs := p.parseMul()

	for p.curTokenIs(token.ADD) || p.curTokenIs(token.SUB) {
		op := p.curToken().Literal
		p.nextToken()
		rhs := p.parseMul()
		if op == "+" {
			lhs = &ast.BinaryExpr{Op: "+", Lhs: lhs, Rhs: rhs}
		} else {
			lhs = &ast.BinaryExpr{Op: "-", Lhs: lhs, Rhs: rhs}
		}
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
	lhs := p.parseAdd()

	//printTree(lhs, 0)

	return lhs
}

func (p *Parser) parseExprStmt() ast.Stmt {
	expr := p.parseExpr()
	es := &ast.ExprStmt{Expr: expr}
	return es
}

func (p *Parser) parseStmt() ast.Stmt {
	stmt := p.parseExprStmt()
	return stmt
}

func (p *Parser) Parse() ast.Node {
	return p.parseStmt()
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
