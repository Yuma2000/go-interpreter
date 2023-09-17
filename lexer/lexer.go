package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置（現在の文字を指し示す）
	readPosition int  // これから読み込む位置（現在の文字の次）
	ch           byte // 現在検査中の文字
}

// New initia
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken analyzes the Lexer's current character and returns the next token from the input.
// If the character doesn't match any known token type, it returns an ILLEGAL token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

// skipWhitespace advances the Lexer's position past any consective whitespace characters.
// It recognizes spaces, tabs, newlines, and carriage returns as whitespace.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readIdentifier reads characters from the Lexer's input starting from the current position
// until it encounters a character that is not part of a valid identifier.
// It returns the identifier as a string.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads a sequence of digit characters from the Lexer's input
// starting from the current position, and returns the number as a string.
// It advances the Lexer's position to just after the last digit of the number.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isDigit returns true if the given byte represents a digit character
// (i.e., '0' to '9'), and false otherwise.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// isLetter checks if the given byte represents an alphabet character (either uppercase or lowercase)
// or an underscore. It returns true if the byte is a valid letter character, and false otherwise.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// newToken creates a new Token object using the provided token type and byte value.
// The byte value is converted to a string to represent the Token's literal.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
