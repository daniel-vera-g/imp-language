package main

import "fmt"

// Evaluator

/////////////////////////
// Statements
/////////////////////////

func (stmt Seq) eval(s ValState) {
	stmt[0].eval(s)
	stmt[1].eval(s)
}

func (whi While) eval(s ValState) {
	for whi.cond.eval(s).valB { // A Boolean value is expected
		whi.stmt.eval(s)
	}
}

func (ite IfThenElse) eval(s ValState) {
	v := ite.cond.eval(s)
	if v.flag == ValueBool {
		switch {
		case v.valB:
			ite.thenStmt.eval(s)
		case !v.valB:
			ite.elseStmt.eval(s)
		}
	} else {
		fmt.Printf("if-then-else eval fail")
	}

}

func (pri Print) eval(s ValState) {
	value := pri.expre.eval(s)

	if value.flag == ValueInt { // If it is an integer
		fmt.Printf("\n %d", value.valI)
	} else if value.flag == ValueBool { // If it is a boolean
		fmt.Printf("\n %t", value.valB)
	}
}

// Maps are represented via points.
// Hence, maps are passed by "reference" and the update is visible for the caller as well.
func (decl Decl) eval(s ValState) {
	v := decl.rhs.eval(s)
	x := (string)(decl.lhs)
	s[x] = v
}

func (assign Assign) eval(s ValState) {
	v := assign.rhs.eval(s)
	x := (string)(assign.lhs)
	s[x] = v
}

/////////////////////////
// Expressions
/////////////////////////

func (x Bool) eval(s ValState) Val {
	return mkBool((bool)(x))
}

func (x Num) eval(s ValState) Val {
	return mkInt((int)(x))
}

func (e Mult) eval(s ValState) Val {
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI * n2.valI)
	}
	return mkUndefined()
}

func (e Plus) eval(s ValState) Val {
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI + n2.valI)
	}
	return mkUndefined()
}

func (e And) eval(s ValState) Val {
	// TODO use if else and do the eval there
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueBool && b1.valB == false:
		return mkBool(false)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB && b2.valB)
	}
	return mkUndefined()
}

func (e Or) eval(s ValState) Val {
	// TODO use if else and do the eval there
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueBool && b1.valB == true:
		return mkBool(true)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB || b2.valB)
	}
	return mkUndefined()
}

func (e Neg) eval(s ValState) Val {
	b1 := e[0].eval(s)
	if b1.flag == ValueBool {
		return mkBool(!b1.valB)
	}
	return mkUndefined()
}

func (e Equ) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)

	switch {

	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB == b2.valB)
	case b1.flag == ValueInt && b2.flag == ValueInt:

		if b1.valI == b2.valI {
			return mkBool(true)
		} else {
			return mkBool(false)
		}
	}
	return mkUndefined()
}
func (e Les) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)

	if b1.flag == ValueInt && b2.flag == ValueInt {
		if b1.valI < b2.valI {
			return mkBool(true)
		} else {
			return mkBool(false)
		}
	}
	return mkUndefined()
}
func (e Gro) eval(s ValState) Val {
	return e[0].eval(s)
}

func (v Var) eval(s ValState) Val {
	return s[(string)(v)]
}
