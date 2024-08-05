package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	type expectedToken struct {
		expectedType    token.TokenType
		expectedLiteral string
	}

	testCases := []struct {
		name     string
		input    string
		expected []expectedToken
	}{
		// case 1
		{
			"NextToken(1)",

			`=+(){},;`,

			[]expectedToken{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},

		// case 2
		{
			"NextToken(2)",

			`let five = 5;
			let ten = 10;
			
			let add = fn(x, y) {
				x + y;
			};

			let result = add(five, ten);
			`,

			[]expectedToken{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},

				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},

				{token.LET, "let"},
				{token.IDENT, "add"},
				{token.ASSIGN, "="},

				{token.FUNCTION, "fn"},
				{token.LPAREN, "("},
				{token.IDENT, "x"},
				{token.COMMA, ","},
				{token.IDENT, "y"},
				{token.RPAREN, ")"},

				{token.LBRACE, "{"},
				{token.IDENT, "x"},
				{token.PLUS, "+"},
				{token.IDENT, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},

				{token.LET, "let"},
				{token.IDENT, "result"},
				{token.ASSIGN, "="},

				{token.IDENT, "add"},
				{token.LPAREN, "("},
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},

				{token.EOF, ""},
			},
		},

		// case 3
		{
			"NextToken(3)",

			`!-/*<5>;`,

			[]expectedToken{
				{token.BANG, "!"},
				{token.MINUS, "-"},
				{token.SLASH, "/"},
				{token.ASTERISK, "*"},
				{token.LT, "<"},
				{token.INT, "5"},
				{token.RT, ">"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			lexer := New(tc.input)

			for i, expected := range tc.expected {
				tok := lexer.NextToken()
				if tok.Type != expected.expectedType {
					tt.Fatalf("%s'test[%d] - token type worng. expected=%q got=%q", tc.name, i, expected.expectedType, tok.Type)
				}

				if tok.Literal != expected.expectedLiteral {
					tt.Fatalf("%s'test[%d] - token literal worng. expected=%q got=%q", tc.name, i, expected.expectedLiteral, tok.Literal)
				}
			}

		})
	}
}

func TestReadIdentifier(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"ReadIdentifier(1)", "hello", "hello"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			lexer := New(tc.input)
			got := lexer.readIdentifier()
			if got != tc.expected {
				tt.Fatalf("readIdentifier worng. expected=%q got=%q", tc.expected, got)
			}
		})
	}
}
