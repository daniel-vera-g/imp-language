package main

// Parser

// Helper functions to build ASTs by hand

/////////////////////////
// Statements
/////////////////////////

func print(x Exp) Stmt {
	return Print{expre: x}
}

func seq(x, y Stmt) Stmt {
	return Seq{x, y}
}

func decl(x string, y Exp) Stmt {
	return Decl{lhs: string(x), rhs: y}
}

func assign(x string, y Exp) Stmt {
	return Assign{lhs: string(x), rhs: y}
}

func ifthenelse(x Exp, y Stmt, z Stmt) Stmt {
	return IfThenElse{cond: x, thenStmt: y, elseStmt: z}
}

func while(x Exp, y Stmt) Stmt {
	return While{cond: x, stmt: y}
}

/////////////////////////
// Expressions
/////////////////////////

func number(x int) Exp {
	return Num(x)
}

func boolean(x bool) Exp {
	return Bool(x)
}

func plus(x, y Exp) Exp {
	return (Plus)([2]Exp{x, y})

	// The type Plus is defined as the two element array consisting of Exp elements.
	// Plus and [2]Exp are isomorphic but different types.
	// We first build the AST value [2]Exp{x,y}.
	// Then cast this value (of type [2]Exp) into a value of type Plus.

}

func mult(x, y Exp) Exp {
	return (Mult)([2]Exp{x, y})
}

func and(x, y Exp) Exp {
	return (And)([2]Exp{x, y})
}

func or(x, y Exp) Exp {
	return (Or)([2]Exp{x, y})
}

func prog(bl Block) Prog {
	return Prog{block: bl}
}

func block(st Stmt) Block {
	return Block{stmt: st}
}

func vars(x string) Var {
	return Var(x)
}

func neg(x Exp) Exp {
	return (Neg)([1]Exp{x})
}

func gro(x Exp) Exp {
	return (Gro)([1]Exp{x})
}

func equ(x, y Exp) Exp {
	return (Equ)([2]Exp{x, y})
}

func les(x, y Exp) Exp {
	return (Les)([2]Exp{x, y})
}
