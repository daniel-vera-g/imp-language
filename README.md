# IMP - Project

> Simple imperative language

Implementation of simple compiler with the following functionalities:

1. Type checker
2. Evaluator
3. Interpreter

- Based on: https://sulzmann.github.io/ModelBasedSW/imp.html
- See lecture notes

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

1. Run project: `go run main.go`
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
