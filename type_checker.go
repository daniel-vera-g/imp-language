package main

// type check

/////////////////////////
// Statements
/////////////////////////

func (stmt Seq) check(t TyState) bool {
	if !stmt[0].check(t) {
		return false
	}
	return stmt[1].check(t)
}

func (decl Decl) check(t TyState) bool {
	ty := decl.rhs.infer(t)
	if ty == TyIllTyped {
		return false
	}

	x := (string)(decl.lhs)
	t[x] = ty
	return true
}

func (a Assign) check(t TyState) bool {
	x := (string)(a.lhs)
	return t[x] == a.rhs.infer(t)
}

func (print Print) check(t TyState) bool {
	return print.expre.infer(t) != TyIllTyped
}

func (while While) check(t TyState) bool {
	// As in the documentation observations,
	// we need to pay attention to the variable inside a while function.
	// Therefore, as t is a map, we need to make a copy of it.
	tCopy := make(TyState)
	for k, v := range t {
		tCopy[k] = v
	}

	return while.cond.infer(t) == TyBool && while.stmt.check(tCopy)
}

func (ifthenelse IfThenElse) check(t TyState) bool {
	// Same procedure as in the while function
	// Difference is that we need to check both the then and else statements
	tCopyThen := make(TyState)
	tCopyElse := make(TyState)
	for k, v := range t {
		tCopyThen[k] = v
		tCopyElse[k] = v
	}

	return ifthenelse.cond.infer(t) == TyBool && ifthenelse.thenStmt.check(tCopyThen) && ifthenelse.elseStmt.check(tCopyElse)
}

/////////////////////////
// Expressions
/////////////////////////

// Type inferencer/checker

func (x Var) infer(t TyState) Type {
	y := (string)(x)
	ty, ok := t[y]
	if ok {
		return ty
	} else {
		return TyIllTyped // variable does not exist yields illtyped
	}
}

func (x Bool) infer(t TyState) Type {
	return TyBool
}

func (x Num) infer(t TyState) Type {
	return TyInt
}

func (e Mult) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		return TyInt
	}
	return TyIllTyped
}

func (e Plus) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		return TyInt
	}
	return TyIllTyped
}

func (e And) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyBool && t2 == TyBool {
		return TyBool
	}
	return TyIllTyped
}

func (e Or) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyBool && t2 == TyBool {
		return TyBool
	}
	return TyIllTyped
}

func (grp Gro) infer(t TyState) Type {
	return grp[0].infer(t)
}

func (less Les) infer(t TyState) Type {
	t1 := less[0].infer(t)
	t2 := less[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		return TyBool
	}
	return TyIllTyped
}

func (neg Neg) infer(t TyState) Type {
	t1 := neg[0].infer(t)
	if t1 == TyBool {
		return TyBool
	}
	return TyIllTyped
}

func (equ Equ) infer(t TyState) Type {
	t1 := equ[0].infer(t)
	t2 := equ[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		return TyBool
	}
	return TyIllTyped
}
