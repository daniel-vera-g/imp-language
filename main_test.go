package main

import (
	"testing"
)

func Test_ex1(t *testing.T) {
	ast := plus(mult(number(1), number(2)), number(0))

	runExpr(ast)
}

func Test_ex2(t *testing.T) {
	ast := and(boolean(false), number(0))
	runExpr(ast)
}

func Test_ex3(t *testing.T) {
	ast := or(boolean(false), number(0))
	runExpr(ast)
}

func TestMain(t *testing.T) {
	main()
}

// TODO add more tests here
// See: https://gobyexample.com/testing
