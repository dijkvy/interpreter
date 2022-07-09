package token

import "fmt"

//TokenType 词法类型
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "OEF"

	// 标识符号
	IDENT = "IDENT"
	INT   = "INT"

	// ASSIGN 运算符
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// COMMA 分隔符号
	COMMA = ","
	// SEMICOLON 分号
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent 判断当前的输入是否是关键字， 如果是关键字， 返回关键字的常量标识符号
// 如果不是， 返回 token.Ident 表示用户自定义的
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
func init() {
	fmt.Println("loading... token.go")
}
