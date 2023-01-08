package main

import "fmt"

type parser struct {
	lexer *impLexer
}

// Parses an Expression
// Gets called on Prefix operators
func (self *parser) expression(rbp int) *token {
	var left *token
	t := self.lexer.next()

	if t.nud != nil {
		left = t.nud(t, self) // Note: Returns itself if it's a prefix operator
	} else {
		panic(fmt.Sprint("NOT PREFIX", t))
	}
	for rbp < self.lexer.peek().bindingPower { // Check binding power...
		t := self.lexer.next() // ...and if it's higher, parse the next token
		if t.led != nil {
			left = t.led(t, self, left) // Select new left token. Should be an expression.
		} else {
			panic(fmt.Sprint("NOT INFIX", t))
		}
	}

	return left
}

// Advance to the next token if it's the expected one
// Used f.ex to check the end of the statement (;)
func (self *parser) advance(expected string) *token {
	tok := self.lexer.next()
	if tok.symbol != expected {
		panic(fmt.Sprint("WAS LOOKING FOR", expected, "GOT", tok))
	}
	return tok
}

// Parses a Block
func (self *parser) block() *token {
	tok := self.lexer.next()
	if tok.symbol != "{" {
		panic(fmt.Sprint("WAS LOOKING FOR BLOCK START", tok))
	}
	block := tok.std(tok, self)
	return block
}

// Parses a Statement
func (self *parser) statement() *token {
	tok := self.lexer.peek()
	if tok.std != nil {
		tok = self.lexer.next()
		return tok.std(tok, self)
	}
	res := self.expression(0)
	self.advance(";")
	return res
}

// Parses multiple statements.
// Used f.ex at the top level
func (self *parser) statements() []*token {
	stmts := []*token{}
	next := self.lexer.peek()
	for next.symbol != "EOF" && next.symbol != "}" {
		stmts = append(stmts, self.statement())
		next = self.lexer.peek()
	}
	return stmts
}

/////////////////////////
// Statements
// Helper functions to build ASTs by hand
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
