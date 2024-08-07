package ast

import (
	"bytes"
	"monkey/token"
)

// Ast를 구성하는 모든 노드는 Node interface를 구현해야 한다.
type Node interface {
	// TokenLiteral은 토큰에 대응되는 리터럴을 반환합니다.
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	// 꼭 필요한 메소드는 아니지만, Go 컴파일러가 에러를 처리하는데 유용하게 쓰인다.
	statementNode()
}

type Expression interface {
	Node
	// 꼭 필요한 메소드는 아니지만, Go 컴파일러가 에러를 처리하는데 유용하게 쓰인다.
	expressionNode()
}

// Program 구조체는 Parser(구문분석기)가 생성하는 노드들의 루트 노드가 된다.
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
	Token token.Token // 이건 어디에 사용되나?
	// Name 변수는 바인딩 식별자 값을 저장합니다.
	Name *Identifier
	// Value 변수는 값을 생성하는 표현식을 저장합니다.
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
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

type Identifier struct {
	Token token.Token
	Value string
}

// 식별자가 값을 생성하기 때문에 expression node이다.
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // 표현식의 첫 번째 토큰
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
