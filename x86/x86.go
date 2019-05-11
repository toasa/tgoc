package x86

import (
	"bytes"
	"fmt"
	"os"
	"tgoc/ast"
	"tgoc/utils"
)

// Identifier name: offset from bsp
var offsets map[string]int

// Identifier map
var varMap map[string]*ast.Ident

// The number of stored identifier to stack
var varCount int

// The number of total identifier
var varNum int

// To assign a unique number to a label.
var labelCount int

// Assembly string to output
var out bytes.Buffer

func initi(identMap map[string]*ast.Ident) {
	offsets = map[string]int{}
	varCount = 1
	varNum = len(identMap)
	varMap = identMap
	labelCount = 0
}

func genExpr(expr ast.Node) {
	switch expr := expr.(type) {
	case *ast.IntLit:
		writeBuf("	push %d\n", expr.Val)
	case *ast.Boolean:
		if expr.Val {
			writeBuf("	push 1\n")
		} else {
			writeBuf("	push 0\n")
		}
	case *ast.LogicalExpr:
		genExpr(expr.Lhs)
		genExpr(expr.Rhs)

		writeBuf("	pop rdi\n")
		writeBuf("	pop rax\n")

		switch expr.Op {
		case "==":
			writeBuf("	cmp rax, rdi\n")
			writeBuf("	sete al\n")
			writeBuf("	movzx rax, al\n")
		case "!=":
			writeBuf("	cmp rax, rdi\n")
			writeBuf("	sete al\n")
			writeBuf("	movzx rax, al\n")
			// 0000 => 0001, 0001 => 0000
			writeBuf("	xor rax, 1\n")
		case "<":
			writeBuf("	cmp rax, rdi\n")
			writeBuf("	setl al\n")
			writeBuf("	movzx rax, al\n")
		case "<=":
			writeBuf("	cmp rax, rdi\n")
			writeBuf("	setle al\n")
			writeBuf("	movzx rax, al\n")
		case ">":
			writeBuf("	cmp rax, rdi\n")
			writeBuf("	setg al\n")
			writeBuf("	movzx rax, al\n")
		case ">=":
			writeBuf("	cmp rax, rdi\n")
			writeBuf("	setge al\n")
			writeBuf("	movzx rax, al\n")
		case "&&":
			writeBuf("	and rax, rdi\n")
		case "||":
			writeBuf("	or rax, rdi\n")
		}
		writeBuf("	push rax\n")

	case *ast.BinaryExpr:
		genExpr(expr.Lhs)
		genExpr(expr.Rhs)

		writeBuf("	pop rdi\n")
		writeBuf("	pop rax\n")

		switch expr.Op {
		case "+":
			writeBuf("	add rax, rdi\n")
		case "-":
			writeBuf("	sub rax, rdi\n")
		case "*":
			writeBuf("	mul rdi\n")
		case "/":
			writeBuf("    xor rdx, rdx\n")
			writeBuf("    div rdi\n")
		case "%":
			writeBuf("    xor rdx, rdx\n")
			writeBuf("    div rdi\n")
			writeBuf("	mov rax, rdx\n")
		case "<<":
			// To change the cl value, changed the rcx value.
			// cl is lower 8 bit register of rcx register.
			writeBuf("	mov rcx, rdi\n")
			writeBuf("	shl rax, cl\n")
		case ">>":
			writeBuf("	mov rcx, rdi\n")
			writeBuf("	sar rax, cl\n")
		case "&":
			writeBuf("	and rax, rdi\n")
		case "|":
			writeBuf("	or rax, rdi\n")
		case "^":
			writeBuf("	xor rax, rdi\n")
		case "&^":
			writeBuf("	xor rdi, rax\n")
			writeBuf("	and rax, rdi\n")
		}
		writeBuf("	push rax\n")

	case *ast.UnaryExpr:
		genExpr(expr.Expr)
		writeBuf("    pop rax\n")

		switch expr.Op {
		case "-":
			writeBuf("    neg rax\n")
		case "!":
			writeBuf("    xor rax, 1\n")
		case "&":
			id, ok := expr.Expr.(*ast.Ident)
			utils.Assert(ok, "&a: a must be identifier")
			os, ok := offsets[id.Name]

			//fmt.Println(id.Name, ":", 8*os)

			utils.Assert(ok, "undefined identifier")
			writeBuf("    mov rax, rbp\n")
			writeBuf("    sub rax, %d\n", 8*os)
		}
		writeBuf("	push rax\n")

	// Dereference
	case *ast.PtrExpr:
		derefCount := 0
		for expr.Of != nil {
			expr = expr.Of
			derefCount++
		}
		genExpr(expr.Expr)
		writeBuf("    pop rax\n")
		for i := 0; i < derefCount; i++ {
			writeBuf("    mov rax, [rax]\n")
		}
		writeBuf("    push rax\n")
	case *ast.Ident:
		os, ok := offsets[expr.Name]
		utils.Assert(ok, "undefined identifier")
		writeBuf("	mov rax, QWORD PTR [rbp - %d]\n", 8*os)
		writeBuf("	push rax\n")
	}
}

func genDecl(vd *ast.VarDecl) {
	offsets[vd.Name] = varCount
	varCount++
}

func genSVDecl(svd *ast.SVDecl) {
	genExpr(svd.Val)
	writeBuf("	pop rax\n")
	writeBuf("	mov QWORD PTR [rbp - %d], rax\n", 8*varCount)
	offsets[svd.Name] = varCount
	varCount++
}

func genStmt(stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.ExprStmt:
		genExpr(stmt.Expr)
		writeBuf("	pop rax\n")
	case *ast.DeclStmt:
		switch decl := stmt.Decl.(type) {
		case *ast.VarDecl:
			genDecl(decl)
		case *ast.SVDecl:
			genSVDecl(decl)
		}
	case *ast.AssignStmt:
		genExpr(stmt.Val)
		writeBuf("	pop rax\n")
		os, ok := offsets[stmt.Name]
		utils.Assert(ok, "undefined identifier")
		writeBuf("	mov QWORD PTR [rbp - %d], rax\n", 8*os)
	case *ast.ReturnStmt:
		genExpr(stmt.Expr)
		writeBuf("	pop rax\n")

		// printNumStdout1()

		// writeBuf("	mov rsp, rbp\n")
		// writeBuf("	pop rbp\n")
		// writeBuf("	ret\n")

		// printNumStdout2()
		return
	case *ast.IfStmt:
		genExpr(stmt.Cond)
		writeBuf("	pop rax\n")
		writeBuf("	cmp rax, 0\n")

		lAlt := makeLabel()

		writeBuf("	je .L%s\n", lAlt)
		genStmts(stmt.Cons)
		if stmt.Alt != nil {
			lEnd := makeLabel()
			writeBuf("	jmp .L%s\n", lEnd)
			writeBuf(".L%s:\n", lAlt)
			genStmts(stmt.Alt)
			writeBuf(".L%s:\n", lEnd)
		} else {
			writeBuf(".L%s:\n", lAlt)
		}
	case *ast.ForSingleStmt:
		loop := makeLabel()
		slipOut := makeLabel()
		writeBuf(".LOOP%s:\n", loop)
		genExpr(stmt.Cond)
		writeBuf("	pop rax\n")
		writeBuf("	cmp rax, 0\n")
		writeBuf("	je .L%s\n", slipOut)
		genStmts(stmt.Stmts)
		writeBuf("	jmp .LOOP%s\n", loop)
		writeBuf(".L%s:\n", slipOut)
	case *ast.ForClauseStmt:
		loop := makeLabel()
		slipOut := makeLabel()
		genStmt(stmt.Init)
		writeBuf(".LOOP%s:\n", loop)
		genExpr(stmt.Cond)
		writeBuf("	pop rax\n")
		writeBuf("	cmp rax, 0\n")
		writeBuf("	je .L%s\n", slipOut)
		genStmts(stmt.Stmts)
		genStmt(stmt.Post)
		writeBuf("	jmp .LOOP%s\n", loop)
		writeBuf(".L%s:\n", slipOut)
	}
}

func genStmts(stmts []ast.Stmt) {
	for _, stmt := range stmts {
		rs, ok := stmt.(*ast.ReturnStmt)
		if !ok {
			genStmt(stmt)
		} else {
			genStmt(rs)
			writeBuf("	jmp _end\n")
			return
		}
	}
}

func gen(stmts []ast.Stmt) {
	if varNum > 0 {
		// なぜ一つの変数につき、rspを16下げる？（8ではなく）
		writeBuf("	sub rsp, %d\n", varNum*16)
		//writeBuf("	sub rsp, %d\n", varNum*8)
	}
	genStmts(stmts)
}

func Gen(stmts []ast.Stmt, identMap map[string]*ast.Ident) {
	initi(identMap)

	//writeBuf(".section	__TEXT,__text,regular,pure_instructions\n")

	writeBuf("	.intel_syntax noprefix\n")
	writeBuf("	.globl _main\n")

	writeBuf("_main:\n")
	writeBuf("	push rbp\n")
	writeBuf("	mov rbp, rsp\n")

	gen(stmts)

	writeBuf("_end:\n")

	printNumStdout1()

	writeBuf("	mov rsp, rbp\n")
	writeBuf("	pop rbp\n")
	writeBuf("	ret\n")

	printNumStdout2()

	// fmt.Printf("%s", out.String())
	makeAssemFile()
}

func makeLabel() string {
	l := fmt.Sprintf("%04d", labelCount)
	labelCount++
	return l
}

func printNumStdout1() {
	writeBuf("	lea	rdi, [rip + L_.str]\n")
	writeBuf("	mov	rsi, rax\n")
	writeBuf("	mov	al, 0\n")
	writeBuf("	call	_printf\n")

	// writeBuf("	xor	ecx, ecx\n")
	// writeBuf("	mov	dword ptr [rbp - 4], eax\n")
	// writeBuf("	mov	eax, ecx\n")
}

func printNumStdout2() {
	//writeBuf(".section	__TEXT,__cstring,cstring_literals\n")
	writeBuf("L_.str:\n")
	writeBuf("	.asciz	\"%%ld\\n\"\n")
}

func writeBuf(argv ...interface{}) {
	arg0, ok := argv[0].(string)
	utils.Assert(ok, "1st args of bufWrite() must be string type")
	str := fmt.Sprintf(arg0, argv[1:]...)
	out.WriteString(str)
}

func makeAssemFile() {
	file, err := os.Create("main.s")
	if err != nil {
		panic(err)
	}
	file.WriteString(out.String())
	file.Close()
}
