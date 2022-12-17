package main

import "strconv"

// Here we have out beloved pretty printer

/////////////////////////
// Statements
/////////////////////////

func (stmt Seq) pretty() string {
	return stmt[0].pretty() + "; " + stmt[1].pretty()
}

func (decl Decl) pretty() string {
	return decl.lhs + " := " + decl.rhs.pretty()
}

func (assign Assign) pretty() string {
	return assign.lhs + " = " + assign.rhs.pretty()
}

func (ifthenelse IfThenElse) pretty() string {
	return "if" + ifthenelse.cond.pretty() + ifthenelse.thenStmt.pretty() + "else" + ifthenelse.elseStmt.pretty()
}

func (print Print) pretty() string {
	return "print" + print.expre.pretty()
}

func (block Block) pretty() string {
	return "{" + block.stmt.pretty() + "}"
}

func (while While) pretty() string {
	return "while" + while.cond.pretty() + while.stmt.pretty()
}

/////////////////////////
// Expressions
/////////////////////////

func (x Var) pretty() string {
	return (string)(x)
}

func (x Bool) pretty() string {
	if x {
		return "true"
	} else {
		return "false"
	}

}

func (x Num) pretty() string {
	return strconv.Itoa(int(x))
}

func (e Mult) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "*"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Plus) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "+"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e And) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "&&"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Or) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "||"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Neg) pretty() string {
	var x string

	x = "!"
	x += e[0].pretty()

	return x
}

func (e Equ) pretty() string {
	var x string

	x = e[0].pretty()
	x += "=="
	x += e[1].pretty()

	return x
}

func (e Les) pretty() string {
	var x string

	x = e[0].pretty()
	x += "<"
	x += e[1].pretty()

	return x
}
func (e Gro) pretty() string {
	var x string

	x = "("
	x += e[0].pretty()
	x += ")"

	return x
}
