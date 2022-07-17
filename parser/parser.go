package parser

import (
	"fmt"

	"interperter/ast"
	"interperter/lexer"
	"interperter/token"
)

type Parser struct {
	// 当前词法分析器实例的指针, 在该实例上不停地调用 NextToken 方法获得下一个词法单元
	l        *lexer.Lexer
	curToken token.Token // 当前分析器分析到的词法单元

	// 如果 current token 没有提供到下一步的信息怎么做， 那么则需要 peekToken 来进行辅助分析
	// 如果有一行代码是 `5;` 如果当前的词法单元是 token.INT, 那么， 需要根据 peekToken 查看
	// 下一个词法单元来确定现在是位于末尾还是在算术表达式的开头， 总而言之， peekToken 是辅助用的
	peekToken token.Token // 下一个词法单元

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s insted",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// next token 方法主要用来向前移动 token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Statements: make([]ast.Statement, 0)}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		} else {
			// fmt.Println("遇到空白", p.curToken)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {

	case token.LET:
		return p.parseLetStatement()

	case token.RETURN:
		return p.parseReturnStatement()

	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// todo : 跳过表达式的处理
	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {

		p.peekError(t)
		return false
	}
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// 解析 return 语法块
func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
