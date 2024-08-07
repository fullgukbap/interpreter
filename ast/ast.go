package ast

import "monkey/token"

// Ast를 구성하는 모든 노드는 Node interface를 구현해야 한다.
type Node interface {
	// TokenLiteral은 토큰에 대응되는 리터럴을 반환합니다.
	TokenLiteral() string
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

type LetStatement struct {
	Token token.Token // 이건 어디에 사용되나?
	// Name 변수는 바인딩 식별자 값을 저장합니다.
	Name *Identifier
	// Value 변수는 값을 생성하는 표현식을 저장합니다.
	Value Expression
}

func (l *LetStatement) statementNode()       {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

// 식별자가 값을 생성하기 때문에 expression node이다.
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
