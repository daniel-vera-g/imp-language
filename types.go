package main

// Our Types are defined here

// The Programm and Block types will be helpful for the parser

// prog      ::= block
type Prog struct {
	block Block
}

// block     ::= "{" statement "}"
type Block struct {
	stmt Stmt
}
