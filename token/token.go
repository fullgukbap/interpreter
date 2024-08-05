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
	/*
		식별자는 값이나 어떠한 코드의 일부분을 담거나 참조할 수 있는 수단이다.
		변수, 상수, 함수명 혹은 코드루프에 대한 레이블을 지정하기 위하여 사용되는 테스트를 칭한다.
		출처: https://m.blog.naver.com/on21life/221565135551
	*/
	IDENT = "IDENT"

	// 리터럴
	INT = "INT"

	// 연산자
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"

	// 구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	LT = "<"
	RT = ">"

	// 예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent 함수는 ident가 예약어 인지 아닌지 확인하여 ident에 올바른 tokenType을 반환합니다.
// 식별자는 예약어를 포함하기 떄문에 이와 같은 함수가 필요합니다.
func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}

	return IDENT
}
