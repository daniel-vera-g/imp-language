package main

import (
	"fmt"
)

// Simple imperative language

// Handle expressions
func handleExpr(e Exp) {
	valueMapping := make(map[string]Val)
	typeMapping := make(map[string]Type)
	fmt.Printf("\n ******* EXPRESSIONS ******* \n")
	fmt.Printf("\n %s", e.pretty())
	fmt.Printf("\n %s", showVal(e.eval(valueMapping)))
	fmt.Printf("\n %s", showType(e.infer(typeMapping)))
}

// Handle statements
func handleStmt(st Stmt) {
	fmt.Printf("\n ******* STATEMENTS ******* \n")
	fmt.Printf("\n %s", st.pretty())

	valueMapping := make(map[string]Val)
	typeMapping := make(map[string]Type)

	// Check whether the statement is well typed
	// Save the Name with the Type in isCorrect
	isCorrect := st.check(typeMapping)

	if isCorrect {
		fmt.Printf("\n Successfully checked the statement")
	} else {
		fmt.Printf("\n Error checking the statement!")
		return
	}

	// Evaluate the statement
	// Save the Name with the Value in s
	st.eval(valueMapping)
	printValueStatemnt(valueMapping)
}

func main() {}
