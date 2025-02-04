package token

type TokenType string

const (
	IDENT = "IDENT"
	INT   = "INT"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	LPAREN   = "LAPAREN"
	RPAREN   = "RPAREN"
	LBRACKET = "LBRACKET"
	RBRACKET = "RBRACKET"

	ASSIGN = "ASSIGN"

	LET      = "LET"
	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"

	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"

	PLUS     = "PLUS"
	MINUS    = "MINUS"
	SLASH    = "SLASH"
	ASTERISK = "ASTERISK"
	BANG     = "BANG"

	LT     = "LT"
	GT     = "GT"
	EQ     = "EQ"
	NOT_EQ = "NOT_EQ"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
