package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// TokenTypes
const (
	// 부가 표현 요소
	ILLEGAL = "ILLEGAL" // 어떤 토큰이나 문자를 렉서가 알 수 없을떄 사용하는 타입
	EOF     = "EOF"     // 파일의 끝을 표현하는 타입 (파서에게 이제 그만 멈춰도 돼 라는 용도로 사용된다.)

	// 식별자
	IDENT = "IDENT"
	INT   = "INT"

	// 연산자
	ASSIGN = "="
	PLUS   = "+"

	// 구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// 예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
