package lexer

import "interperter/token"

type Lexer struct {
	input        string // 输入的内容
	position     int    // 所输入字符串中的当前位置(指向的是当前的字符)
	readPosition int    // 所输入字符串中的当前读取位置(指向当前字符之后的后一个字符)
	ch           byte   // 当前正在查看的字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input, position: 0, readPosition: 0, ch: 0}
	l.readChar()
	return l
}
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case ',':
		tok = newToken(token.COMMA, l.ch)

	case '+':
		tok = newToken(token.PLUS, l.ch)

	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)

	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)

		// default
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// ...
	l.readChar()
	return tok
}

// 跳过空白字符
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' ||
		l.ch == '\t' ||
		l.ch == '\n' ||
		l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch >= 'z' ||
		'A' <= ch && ch >= 'Z'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readChar 主要是读取 input 中的下一个字符
// 并将 改变 position 和 readPosition 的位置
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition = l.readPosition + 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// read a word
func (l *Lexer) readIdentifier() string {

	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	// sub string
	return l.input[position:l.position]
}
