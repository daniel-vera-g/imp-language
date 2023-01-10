# IMP - Project

[![Go](https://github.com/daniel-vera-g/imp-language/actions/workflows/go.yml/badge.svg)](https://github.com/daniel-vera-g/imp-language/actions/workflows/go.yml)

> Simple imperative language

## About

Implementation of simple compiler with the following functionalities:

1. **Parser** :

- Parses a given Program, Statement, Expression,... to a proper AST that can be further processed
- For example: `1+0 -> plus(number(1), number(0))`

2. **Type checker** :

- Checks the types of the given Programm, Statement,...
- For example: `1 -> int`

3. **Evaluator** :

- Evaluates the given Expression
- For example: `1+0 -> 1`

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

1. Build: `go build main.go`
2. Run: `./imp-project 1+2*3` (First parameter is Program to compile)
3. Run tests: `go test -v`

## TODOs

- [x] GitHub Repository
  - [x] CI Pipeline
- [ ] Type checker
  - [x] Implementation
  - [ ] Tests
- [ ] Evaluator
  - [x] Implementation
  - [ ] Tests
- [ ] Parser
  - [ ] Lexer: Implementation
  - [ ] Parser: Implementation
  - [ ] Tests
- [ ] Code TODOs

## References

- Based on: https://sulzmann.github.io/ModelBasedSW/imp.html (See lecture notes)
- Used Parser(Pratt Parser):
  - https://dl.acm.org/doi/10.1145/512927.512931
  - https://crockford.com/javascript/tdop/tdop.html
