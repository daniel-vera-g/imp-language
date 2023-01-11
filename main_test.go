package main

import (
	"fmt"
	"testing"
)

// Parser Tests

type parser_tests struct {
	code        string
	description string
}

func Test_parser_statements(t *testing.T) {

	statements_parser_test := []parser_tests{
		{"x := 1; x = 3;", "Sequence, Declaration & Assignment"},
		{"x := true; while x { x = false; y := 4;}", "While"},
		{"if false { x := 2+1; } else { x := 2; }", "ifthenelse"},
		{"print 1+2", "print"},
	}

	for _, test := range statements_parser_test {

		fmt.Printf("\n------------------------\n")
		fmt.Printf("%s", test.description)
		fmt.Printf("\n------------------------\n")

		l := newLexer(test.code)
		p := parser{lexer: l}

		statements := p.statements()
		ast := prog(block(iterateStatements(statements)))
		ast.handleProgram()
	}
}

func Test_parser_expressions(t *testing.T) {

	expression_parser_tests := []parser_tests{
		{"x := 1 + 2;", "Plus"},
		{"x := 1 * 2;", "Mult"},
		{"x := 1 + 2 * 3;", "Mult & Plus"},
		{"x := true && false;", "And"},
		{"x := true || false;", "Or"},
		{"x := !true;", "Neg"},
		{"x := 1 == 2;", "Equal"},
		{"x := 1 < 2;", "LessThan"},
		{"x := (1+2)*3;", "Grouping"},
	}

	for _, test := range expression_parser_tests {

		fmt.Printf("\n------------------------\n")
		fmt.Printf("%s", test.description)
		fmt.Printf("\n------------------------\n")

		l := newLexer(test.code)
		p := parser{lexer: l}

		statements := p.statements()
		ast := prog(block(iterateStatements(statements)))
		ast.handleProgram()
	}
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

// Tests Mail - Zwischenstand

func Test_shortCircuit(t *testing.T) {
	astOr := decl("x", or(boolean(true), boolean(false)))
	astAnd := decl("x", and(boolean(false), boolean(true)))
	handleStmt(astOr)
	handleStmt(astAnd)
}
