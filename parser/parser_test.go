package parser

import (
	"testing"

	"interperter/ast"
	"interperter/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let  x = 5;
	let y = 10;
	let foobar = 838383;
`
	//	input = `
	//	let  x  5;
	//	let y  10;
	//	let foobar  838383;
	//`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	// 检查是否有语法错误
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() return nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		t.Logf("[%d] stmt:%v\n", i, stmt.TokenLiteral())
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

	}

}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got = %T", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatment. got = %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStatement.Name.Value not %s, got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got = %s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	if len(p.Errors()) == 0 {
		return
	}

	t.Errorf("parser has %d errors\n", len(p.Errors()))

	for _, msg := range p.Errors() {
		t.Errorf("parser error:%q", msg)
	}
	t.FailNow()

}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 233;
	return 233 332;
	
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got =%t",
				stmt,
			)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral not 'return', got=%q",
				returnStmt.TokenLiteral(),
			)
		}
	}

}
