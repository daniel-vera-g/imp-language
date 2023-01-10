package main

import (
	"testing"
)

// Parser Tests

func Test_parser1(t *testing.T) {
	l := newLexer("x := 1 ; x=x+2;")
	p := parser{lexer: l}

	statements := p.statements()
	// stmt := statements[0]
	ast := prog(block(iterateStatements(statements)))
	ast.handleProgram()
}

// Expressions tests

func Test_ex1(t *testing.T) {
	ast := plus(mult(number(1), number(2)), number(0))
	handleExpr(ast)
}

func Test_ex2(t *testing.T) {
	ast := and(boolean(false), number(0))
	handleExpr(ast)
}

func Test_ex3(t *testing.T) {
	ast := or(boolean(false), number(0))
	handleExpr(ast)
}

// Statements tests

func Test_st1(t *testing.T) {
	ast := seq(decl("x", number(1)), printStmt(plus(mult(number(1), number(2)), number(0))))
	handleStmt(ast)
}

func Test_st2(t *testing.T) {
	ast := printStmt(and(boolean(false), number(0)))
	handleStmt(ast)
}

func Test_st3(t *testing.T) {
	ast := printStmt(or(boolean(false), number(0)))
	handleStmt(ast)
}

func TestMain(t *testing.T) {
	main()
}

// Program tests

func Test_p1(t *testing.T) {
	ast := prog(block(seq(decl("x", number(1)), printStmt(plus(mult(number(1), number(2)), number(0))))))
	ast.handleProgram()
}

func Test_p2(t *testing.T) {
	ast := prog(block(printStmt(and(boolean(false), number(0)))))
	ast.handleProgram()
}

func Test_p3(t *testing.T) {
	ast := prog(block(printStmt(or(boolean(false), number(0)))))
	ast.handleProgram()
}

// TODO add more tests here
// See: https://gobyexample.com/testing
