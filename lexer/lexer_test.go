package lexer

import (
	"testing"

	"interperter/token"
)

func TestNextToken(t *testing.T) {

	var input = `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
}

let result = add(five, ten);

!-/*5;
5< 10 >5;

if ( 5 < 10) {
	return true;
} else {
	return false;
}

10  == 10;

10 != 9;
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{

		// 定义变量
		// let five = 5;
		{expectedLiteral: "let", expectedType: token.LET},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		// let ten = 10;
		{expectedLiteral: "let", expectedType: token.LET},
		{expectedType: token.IDENT, expectedLiteral: "ten"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.INT, expectedLiteral: "10"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},

		// let add = func(x, y) {x +y;}
		// 一些函数定义的语句块
		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENT, expectedLiteral: "add"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.FUNCTION, expectedLiteral: "fn"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.IDENT, expectedLiteral: "x"},
		{expectedType: token.COMMA, expectedLiteral: ","},
		{expectedType: token.IDENT, expectedLiteral: "y"},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.LBRACE, expectedLiteral: "{"},
		{expectedType: token.IDENT, expectedLiteral: "x"},
		{expectedType: token.PLUS, expectedLiteral: "+"},
		{expectedType: token.IDENT, expectedLiteral: "y"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.RBRACE, expectedLiteral: "}"},

		// let result = add(five, ten);
		// 函数调用并结果赋值
		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENT, expectedLiteral: "result"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.IDENT, expectedLiteral: "add"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.IDENT, expectedLiteral: "five"},
		{expectedType: token.COMMA, expectedLiteral: ","},
		{expectedType: token.IDENT, expectedLiteral: "ten"},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},

		// !-/*5;
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		// 5< 10 > 5;
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		//if ( 5 < 10) {return true;} else {return false;}
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		// 10 == 10;
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		// 10 != 9;
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		t.Logf("%#v", tok)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i,
				tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i,
				tt.expectedLiteral, tok.Literal)
		}
	}
}
