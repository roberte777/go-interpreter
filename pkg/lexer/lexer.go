package lexer

import (
	"github.com/roberte777/go-interpreter/pkg/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte //current character being examined
}

func New(input string) *Lexer {
	//TODO: pass in file name and line number somehow to give better errors
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	//check if we are at the end of the input
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		//read the next char
		l.ch = l.input[l.readPosition]
	}
	//move the current position forward 1
	l.position = l.readPosition
	//keep readPostion one character ahead of current position
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	//make a token out of the current character
	l.eatWhitespace()
	var tok token.Token
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
			//need the early return because readIdentifier will move us to the
			//next character, so l.readChar() after this switch would move too
			//far
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}
	//set up for the next character
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	//ending is not inclusive
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	//TODO: implement this
	//if I want to allow ?, !, etc. to be part of an identifier,
	//I need to add them to this list
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	//TODO: allow floating point numbers
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
