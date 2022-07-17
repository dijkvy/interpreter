package ast

import (
	"bytes"

	"interperter/token"
)

type Node interface {
	// TokenLiteral 返回关联的词法单元字面量
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type LetStatement struct {
	Token token.Token // 用来保存 let 词法单元的
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {
	// todo how to solve this
}

type Identifier struct {
	Token token.Token // 这里是用来保存 ident 词法单元的
	Value string
}

func (i *Identifier) String() string {
	return i.Value
}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {
	// todo how to solve this statement
}

// ReturnStatement 解析这种
// return 233;
// return sum(3,4);
type ReturnStatement struct {
	Token       token.Token // return 词法单元
	ReturnValue Expression
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) statementNode() {
	// todo solve me later
}

// ExpressionStatement 表达式解析器
// 用来解析比如:
// 2 + 3;
// (2 + 3) * 3
// 2 + 3 * 3
type ExpressionStatement struct {
	Token   token.Token
	Express Expression
}

func (es *ExpressionStatement) String() string {
	var out bytes.Buffer

	if es.Express != nil {
		out.WriteString(es.Express.String())
	}

	return out.String()
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) statementNode() {
	// todo solve me later
}
