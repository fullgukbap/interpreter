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

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
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
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	}

	l.readChar()
	return tok
}

// newToken 함수는 tokenType과 ch를 기반으로 monkey/token.Token을 생성하여 반환합니다.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	tok := token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}

	return tok
}
