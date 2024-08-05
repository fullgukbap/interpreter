package lexer

import "monkey/token"

type Lexer struct {
	input        string // source code
	position     int    // input에서의 현재 위치(ch의 값의 문자를 가리킴)
	readPosition int    // input 에서의 다음 위치
	ch           byte   // 현재 문자, ch의 타입이 byte이므로 ASCII만 지원한다. 만약 한글과 같은 언어를 지원하고 한다면 RUNE 타입을 사용하자
}

// New 함수는 input(source code)을 기반으로 생성한 lexer의 주소값을 반환합니다.
func New(input string) *Lexer {
	lexer := Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
		ch:           0, // ch는 현재 읽고 있는 문자를 뜻하며, 초기값으로 ASCII 코드 문자 0인 'NUL'에 해당하며 '아직 아무것도 읽지 않은 상태여서 문자가 없음' 및 '파일의 끝에 도달해 문자가 없음'을 두 가지 의미를 내포하고 있으며, 본 초기화 과정에서의 ch변수값의 의미는 전자이다.
	}
	// ch의 의미 중 하나인 '아직 아무것도 읽지 않은 상태여서 문자가 없음'을 없애기 위해 초기화 과정에서 lexer.readChar() 함수 호출 (즉 이 말은 lexer를 준비상태로 만든다)
	lexer.readChar()
	/*
		lexer {
			input: "hello",
			position: 0,
			readPosition: 1,
			ch: 'h'
		}
	*/

	return &lexer
}

// readChar 함수는 포인터 리시버로써 lexer의 값을 참조하여 종래의 input의 다음 문자를 읽어 ch, position, readPosition의 값을 갱신합니다.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 본 0의 의미는 '파일의 끝에 도달해 문자가 없음'을 뜻한다.
		// return 여기에 return 걸어주면 좋을듯. 그러면 l.positoin과 l.readPosition이 계속 올라가는 경우를 없앨 수 있음
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

// NextToken 함수는 Lexer 구조체의 포인터 리시버로써 l 구조체의 input을 기반으로 토큰을 추출하여 반환합니다.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
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
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.RT, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		// 위에 조건을 충족하지 않으면 식별자로 보겠다.
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // 여기서 early exist를 해줘야 할까? (아직 이유를 모르겠다.)
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok // 여기서 early exist를 해줘야 할까? (아직 이유를 모르겠다.)
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// skipWhitespace 함수는 렉서가 처리할 필요가 없는 문자를 생략하여 넘깁니다.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// newToken 함수는 tokenType과 ch를 기반으로 monkey/token.Token을 생성하여 반환합니다.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	tok := token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}

	return tok
}

// readIdentifier 함수는 문자열을 추출합니다.
// 단 이 함수가 동작하기 위해서는 position이 isLetter()에 충족되는 문자를 가리키고 있어야 합니다.
func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

// isLetter 함수는 문자인지 아닌지 검사하여 참, 거짓을 반환합니다.
// monkey 언어에서의 문자의 조건은 다음과 같습니다.
// 'a' ~ 'z' || 'A' ~ 'Z' || ch == '_'
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// readNumber 함수는 숫자를 추출합니다.
// 단 이 함수가 동작하기 위해서는 position이 isDigit()에 충족되는 문자를 가리키고 있어야 합니다.
// 본 함수는 오직 10진수만 받아들일 수 있도록 코딩되었다.
func (l *Lexer) readNumber() string {
	startPosition := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

// isDigit 함수는 숫자인지 아닌지 검하여 참, 거짓을 반환합니다.
// monkey 언어에서의 숫자의 조건은 다음과 같습니다.
// '0' ~ '9'
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
