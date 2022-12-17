package main

import (
	"testing"
)

// Test expressions

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

// Test statements

func Test_st1(t *testing.T) {
	ast := seq(decl("x", number(1)), print(plus(mult(number(1), number(2)), number(0))))
	handleStmt(ast)
}

func Test_st2(t *testing.T) {
	ast := print(and(boolean(false), number(0)))
	handleStmt(ast)
}

func Test_st3(t *testing.T) {
	ast := print(or(boolean(false), number(0)))
	handleStmt(ast)
}

func TestMain(t *testing.T) {
	main()
}

// TODO add more tests here
// See: https://gobyexample.com/testing
