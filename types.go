package main

// Types

// Types

type Type int

const (
	TyIllTyped Type = 0
	TyInt      Type = 1
	TyBool     Type = 2
)

func showType(t Type) string {
	var s string
	switch {
	case t == TyInt:
		s = "Int"
	case t == TyBool:
		s = "Bool"
	case t == TyIllTyped:
		s = "Illtyped"
	}
	return s
}

// Value State is a mapping from variable names to values
type ValState map[string]Val

// Value State is a mapping from variable names to types
type TyState map[string]Type

// Interface

type Exp interface {
	pretty() string
	eval(s ValState) Val
	infer(t TyState) Type
}

type Stmt interface {
	pretty() string
	eval(s ValState)
	check(t TyState) bool
}

// Statement cases (incomplete)

type Seq [2]Stmt

type Decl struct {
	lhs string
	rhs Exp
}
type Assign struct {
	lhs string
	rhs Exp
}

type While struct {
	cond Exp
	stmt Stmt
}

type IfThenElse struct {
	cond     Exp
	thenStmt Stmt
	elseStmt Stmt
}

type Print struct {
	expre Exp
}

// Expression cases (incomplete)

type Bool bool
type Num int
type Mult [2]Exp
type Plus [2]Exp
type And [2]Exp
type Or [2]Exp
type Neg [1]Exp
type Equ [2]Exp
type Les [2]Exp
type Gro [1]Exp

type Var string

// The Programm and Block types will be helpful for the parser

// prog      ::= block
type Prog struct {
	block Block
}

// block     ::= "{" statement "}"
type Block struct {
	stmt Stmt
}
