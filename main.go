package main

import (
	"fmt"
	"os"
)

// Simple imperative language

// Handle expressions
// Type checks and evaluates an expression
func handleExpr(e Exp) {
	valueMapping := make(map[string]Val)
	typeMapping := make(map[string]Type)
	fmt.Printf("\n ******* EXPRESSIONS ******* \n")
	fmt.Printf("\n %s", e.pretty())
	fmt.Printf("\n %s", showVal(e.eval(valueMapping)))
	fmt.Printf("\n %s", showType(e.infer(typeMapping)))
}

// Handle statements
// Type checks and evaluates a statement
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

// Handle program
// Type checks and evaluates a program
func (prog Prog) handleProgram() {

	// First do some type checking
	isTypeSave := prog.checkProgType()

	if isTypeSave {
		fmt.Printf("\n Successfully checked the program")

		valueMapping := make(map[string]Val)

		// Evaluate the program
		// Save the Name with the Value in valueMapping
		prog.block.stmt.eval(valueMapping)
		printValueStatemnt(valueMapping)
	} else {
		fmt.Printf("\n Error checking the program types!")
		return
	}
}

func main() {
	stmt := os.Args[1]
	// TODO accept input parameters, parse them, type check them, and evaluate them

	l := newLexer(stmt)
	p := parser{lexer: l}

	// TODO For Prog and Block?
	// See: https://ilias.h-ka.de/ilias.php?ref_id=719974&cmdClass=ilobjforumgui&thr_pk=18066&page=0&cmd=viewThread&cmdNode=xm:o9&baseClass=ilrepositorygui
	// Ich schlage "entweder oder" vor. D.h. entweder prog ::= block oder prog ::= statement

	statements := p.statements()
	for _, s := range statements {
		handleStmt(buildAstStmt(s))
	}
}
