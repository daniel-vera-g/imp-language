# IMP - Project

[![Go](https://github.com/daniel-vera-g/imp-language/actions/workflows/go.yml/badge.svg)](https://github.com/daniel-vera-g/imp-language/actions/workflows/go.yml)

> Simple imperative language

- Based on: https://sulzmann.github.io/ModelBasedSW/imp.html
- See lecture notes

## About

Implementation of simple compiler with the following functionalities:

1. Type checker: Checks the types of the given Programm, Statement,... (F.ex 1 -> int)
2. Evaluator: Evaluates the given Expression (F. ex 0 + 1 -> 1)
3. Parser: Parses a given Program, Statement, Expression,... to a proper AST that can be evaluated

_Syntax definition:_

```go
vars       Variable names, start with lower-case letter

prog      ::= block
block     ::= "{" statement "}"
statement ::=  statement ";" statement           -- Command sequence
            |  vars ":=" exp                     -- Variable declaration
            |  vars "=" exp                      -- Variable assignment
            |  "while" exp block                 -- While
            |  "if" exp block "else" block       -- If-then-else
            |  "print" exp                       -- Print

exp ::= 0 | 1 | -1 | ...     -- Integers
     | "true" | "false"      -- Booleans
     | exp "+" exp           -- Addition
     | exp "*" exp           -- Multiplication
     | exp "||" exp          -- Disjunction
     | exp "&&" exp          -- Conjunction
     | "!" exp               -- Negation
     | exp "==" exp          -- Equality test
     | exp "<" exp           -- Lesser test
     | "(" exp ")"           -- Grouping of expressions
     | vars                  -- Variables
```

## Usage

1. Build and run: `go build main.go && ./imp-project`
2. Run tests: `go test -v`

## TODO

## General

- [ ] GitHub Repository
  - [ ] CI Pipeline
- [ ] Type checker
  - [ ] Tests
- [ ] Evaluator
  - [ ] Tests
- [ ] Parser
  - [ ] Tests
