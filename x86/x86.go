package x86

import (
	"fmt"
	"tgoc/ast"
)

func gen(node ast.Node) {
	il, ok := node.(*ast.IntLit)
	if ok {
		fmt.Printf("	push %d\n", il.Val)
		return
	}

	be, ok := node.(*ast.BinaryExpr)
	if ok {
		gen(be.Lhs)
		gen(be.Rhs)

		switch be.Op {
		case "*":
			fmt.Printf("	pop rdi\n")
			fmt.Printf("	pop rax\n")
			fmt.Printf("	mul rdi\n")
			fmt.Printf("	push rax\n")
		case "+":
			fmt.Printf("	pop rdi\n")
			fmt.Printf("	pop rax\n")
			fmt.Printf("	add rax, rdi\n")
			fmt.Printf("	push rax\n")
		case "-":
			fmt.Printf("	pop rdi\n")
			fmt.Printf("	pop rax\n")
			fmt.Printf("	sub rax, rdi\n")
			fmt.Printf("	push rax\n")
		}
		return
	}

	es, ok := node.(*ast.ExprStmt)
	if ok {
		gen(es.Expr)
	}

	//fmt.Printf("	pop rdi\n")
	fmt.Printf("	pop rax\n")
}

func Gen(node ast.Node) {
	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".globl _main\n")
	fmt.Printf("_main:\n")
	fmt.Printf("	push rbp\n")
	fmt.Printf("	mov rbp, rsp\n")

	gen(node)

	fmt.Printf("	pop rbp\n")
	fmt.Printf("	ret\n")
}
