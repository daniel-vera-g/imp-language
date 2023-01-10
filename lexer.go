package main

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

/////////////////////////
// tokenRegistry:
// Stores the tokens from the lexer in a map.
/////////////////////////

type tokenRegistry struct {
	symbolMap map[string]*token
}

// Create and populate a new token registry
// Initializes all the tokens, binding powers and functions
func newTokenRegistry() *tokenRegistry {
	tokenReg := &tokenRegistry{make(map[string]*token)}

	tokenReg.symbol("IDENTIFIER")
	tokenReg.symbol("NUMBER")
	tokenReg.symbol("true")
	tokenReg.symbol("false")

	tokenReg.consumable(")")
	tokenReg.consumable("else")
	tokenReg.consumable("EOF")
	tokenReg.consumable("{")
	tokenReg.consumable("}")
	tokenReg.consumable(";")

	tokenReg.infix("+", 50)
	tokenReg.infix("*", 60)
	tokenReg.infix("==", 20)
	tokenReg.infix("<", 30)

	tokenReg.infixRight("||", 25) // TODO same binding power as && ?
	tokenReg.infixRight("&&", 25)
	tokenReg.infixRight("=", 10)
	// TODO Difference declaration and assignment?
	tokenReg.infixRight(":=", 10)

	tokenReg.prefix("!")
	tokenReg.prefix("-")

	// Grouping expressions
	tokenReg.prefixNud("(", func(t *token, p *parser) *token {
		if p.lexer.peek().symbol != ")" {
			// Check inner expressions
			for {
				if p.lexer.peek().symbol == ")" {
					break
				}
				t.children = append(t.children, p.expression(0))
				if p.lexer.peek().symbol != "," {
					break
				}
				p.advance(",")
			}
		}
		p.advance(")")
		t.symbol = "()"
		t.value = "GROUP"
		return t
	})

	tokenReg.statement("if", func(t *token, p *parser) *token {
		t.children = append(t.children, p.expression(0))
		t.children = append(t.children, p.block())

		// Check for the else branch
		nextToken := p.lexer.peek()
		if nextToken.symbol == "else" {
			// Go into else branch
			p.lexer.next()
			next := p.lexer.peek()
			// Another if?
			if next.value == "if" {
				t.children = append(t.children, p.statement())
			} else {
				t.children = append(t.children, p.block())
			}
		}
		return t
	})

	tokenReg.statement("while", func(t *token, p *parser) *token {
		t.children = append(t.children, p.expression(0))
		t.children = append(t.children, p.block())
		return t
	})

	tokenReg.statement("{", func(t *token, p *parser) *token {
		t.children = append(t.children, p.statements()...)
		p.advance("}")
		return t
	})

	return tokenReg
}

// Is the given String a defined Symbol?
func (self *tokenRegistry) defined(symbol string) bool {
	_, ok := self.symbolMap[symbol]
	return ok
}

// Register new tokens
func (self *tokenRegistry) register(symbol string, bindingPower int, nud nudFunc, led ledFunc, std stdFunc) {
	if value, ok := self.symbolMap[symbol]; ok { // If the symbol is already defined
		if nud != nil && value.nud == nil {
			value.nud = nud
		}
		if led != nil && value.led == nil {
			value.led = led
		}
		if std != nil && value.std == nil {
			value.std = std
		}
		if bindingPower > value.bindingPower {
			value.bindingPower = bindingPower
		}
	} else {
		self.symbolMap[symbol] = &token{bindingPower: bindingPower, nud: nud, led: led, std: std}
	}
}

// An Infix token has a left hand side and a right hand side.
// Default Infix function.
func (self *tokenRegistry) infix(symbol string, bindingPower int) {
	self.register(symbol, bindingPower, nil, func(t *token, p *parser, left *token) *token {
		t.children = append(t.children, left)
		t.children = append(t.children, p.expression(t.bindingPower))
		return t
	}, nil)
}

// An InfixRight gives the left hand side a higher binding power than the right hand side.
// Has custom led function.
func (self *tokenRegistry) infixRight(symbol string, bindingPower int) {
	self.register(symbol, bindingPower, nil, func(t *token, p *parser, left *token) *token {
		t.children = append(t.children, left)
		t.children = append(t.children, p.expression(t.bindingPower-1))
		return t
	}, nil)
}

// A Prefix token has only a right hand side.
// Default Prefix function.
func (self *tokenRegistry) prefix(symbol string) {
	self.register(symbol, 0, func(t *token, p *parser) *token {
		t.children = append(t.children, p.expression(100))
		return t
	}, nil, nil)
}

// prefixNud has, like prefix, a higher binding power than the right hand side.
// It also takes a custom nud function.
func (self *tokenRegistry) prefixNud(symbol string, nud nudFunc) {
	self.register(symbol, 0, nud, nil, nil)
}

// Statement is like nud, but it only gets used at the beginning of a statement.
func (self *tokenRegistry) statement(symbol string, std stdFunc) {
	self.register(symbol, 0, nil, nil, std)
}

// A Symbol can be an identifier, number or boolean
func (self *tokenRegistry) symbol(symbol string) {
	self.register(symbol, 0, func(t *token, p *parser) *token { return t }, nil, nil)
}

// A consumable just gets consumed and not returned.
func (self *tokenRegistry) consumable(symbol string) {
	self.register(symbol, 0, nil, nil, nil)
}

/////////////////////////
// Tokens
/////////////////////////

type token struct {
	symbol       string
	value        string
	line         int
	column       int
	bindingPower int
	nud          nudFunc  // null denotation
	led          ledFunc  // left denotation
	std          stdFunc  // statement denotation
	children     []*token // for AST
}

// Generates a new token.
// Is used by the lexer.
// Uses the parameters given by the lexer and the information from the symbolMap.
func (self *tokenRegistry) token(symbol string, value string, line int, column int) *token {
	return &token{
		symbol:       symbol,
		value:        value,
		line:         line,
		column:       column,
		bindingPower: self.symbolMap[symbol].bindingPower,
		nud:          self.symbolMap[symbol].nud,
		led:          self.symbolMap[symbol].led,
		std:          self.symbolMap[symbol].std,
	}
}

// null denotation.
// Can be used with values and prefix operators
type nudFunc func(*token, *parser) *token

// left denotation.
// Can be used for infix and suffix operators
type ledFunc func(*token, *parser, *token) *token

// statement denotation.
// Can be used for statements
type stdFunc func(*token, *parser) *token

/////////////////////////
// Lexer
/////////////////////////

type impLexer struct {
	tokenRegistry *tokenRegistry
	sourceString  string
	indexInString int
	line          int
	col           int
	token         *token
	isCached      Bool
	last          *token
}

// This is the constructor for our lexer.
// It takes a string and returns a lexer.
func newLexer(sourceSyntax string) *impLexer {
	return &impLexer{newTokenRegistry(), sourceSyntax, 0, 1, 1, nil, false, nil}
}

// Moves over the String and parses the next token
func (self *impLexer) next() *token {
	// We're moving on...
	// Invalidate the cache from the peek operation
	self.isCached = false

	// Set index to the next character
	tmpIndex := -1 // Case we're beginning...
	if self.indexInString != tmpIndex {
		tmpIndex = self.indexInString
		self.removeWhitespace()
	}

	// Check for the end of the indexInString
	if 0 == len(self.sourceString[self.indexInString:]) {
		return self.tokenRegistry.token("EOF", "EOF", self.line, self.col)
	}

	// Now do actual parsing...
	var text bytes.Buffer
	r, size := utf8.DecodeRuneInString(self.sourceString[self.indexInString:])
	for size > 0 && !unicode.IsSpace(r) {
		if isFirstIdentifierChar(r) { // Parse Identifier or Keywords
			col := self.col
			self.consumeRune(&text, r, size)
			// Consume the rest of the identifier...(F.ex a word)
			for {
				r, size = utf8.DecodeRuneInString(self.sourceString[self.indexInString:])
				if size > 0 && isIdentifierChar(r) {
					self.consumeRune(&text, r, size)
				} else {
					break
				}
			}
			symbol := text.String()                 // We're done, assemble...
			if self.tokenRegistry.defined(symbol) { // Do we have a keyword?
				return self.tokenRegistry.token(symbol, symbol, self.line, col)
			} else { // We have an Identifier
				return self.tokenRegistry.token("IDENTIFIER", symbol, self.line, col)
			}
		} else if unicode.IsDigit(r) { // Parse Number
			col := self.col
			self.consumeRune(&text, r, size)
			// Consume the rest of the number...
			for {
				r, size = utf8.DecodeRuneInString(self.sourceString[self.indexInString:])
				if size > 0 && unicode.IsDigit(r) {
					self.consumeRune(&text, r, size)
				} else {
					break
				}
			}
			return self.tokenRegistry.token("NUMBER", text.String(), self.line, col)
		} else if isOperatorChar(r) { // Parse Operators
			col := self.col
			self.consumeRune(&text, r, size)

			// We also have two character operators: || && ==
			// Check for these...
			var twoChar bytes.Buffer
			twoChar.WriteRune(r)
			r, size = utf8.DecodeRuneInString(self.sourceString[self.indexInString:])
			if size > 0 && isOperatorChar(r) {
				twoChar.WriteRune(r)
				if self.tokenRegistry.defined(twoChar.String()) {
					self.consumeRune(&text, r, size)
					return self.tokenRegistry.token(twoChar.String(), twoChar.String(), self.line, col)
				}
			}
			// No two character operator found, return the one character operator
			text := text.String()
			if self.tokenRegistry.defined(text) {
				return self.tokenRegistry.token(text, text, self.line, col)
			}
		} else {
			// We have an unknown character...
			break
		}
	}
	panic("Unknown character: " + string(r))
}

// Is the character a valid operator?
func isOperatorChar(r rune) bool {
	possibleOperators := "|+*&!=()<;:={}"
	for _, op := range possibleOperators {
		if r == op {
			return true
		}
	}
	return false
}

// Consumes a rune and appends it to the text buffer.
func (self *impLexer) consumeRune(text *bytes.Buffer, r rune, size int) {
	text.WriteRune(r)
	self.indexInString += size
	self.col++
}

// It is a proper beginning of an identifier if it is a letter or an underscore.
func isFirstIdentifierChar(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

// It is a proper character for the identifier
func isIdentifierChar(r rune) bool {
	return isFirstIdentifierChar(r) || unicode.IsDigit(r)
}

// Removes the white space from the source string.
func (self *impLexer) removeWhitespace() {
	rune, size := utf8.DecodeRuneInString(self.sourceString[self.indexInString:])
	for size > 0 && unicode.IsSpace(rune) {
		if rune == '\n' {
			self.line++
			self.col = 1
		} else {
			self.col++
		}
		self.indexInString += size
		rune, size = utf8.DecodeRuneInString(self.sourceString[self.indexInString:])
	}
}

// Peeks to the next token without moving the indexInString.
// Creates a cache for the next query
func (self *impLexer) peek() *token {
	if self.isCached {
		return self.token
	}

	// Save state
	tmpIndex := self.indexInString
	tmpLine := self.line
	tmpCol := self.col

	// Get & Cache next token for future peeks)
	nextToken := self.next()
	self.token = nextToken
	self.isCached = true

	// Restore state
	self.indexInString = tmpIndex
	self.line = tmpLine
	self.col = tmpCol

	return nextToken
}
