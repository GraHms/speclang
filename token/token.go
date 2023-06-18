package token

// TokenType represents the type of a token.
type TokenType string

// Token represents a token in the "speclang" language.
type Token struct {
	Type    TokenType
	Literal string
}

// TokenType constants
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT      = "IDENT" // add, foobar, x, y, ...
	INT        = "INT"   // 1234567890
	STRING     = "STRING"
	FLOAT      = "FLOAT"
	ANNOTATION = "ANNOTATION"
	NEWLINE    = "NEWLINE"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	EQ     = "=="
	NOT_EQ = "!="

	LT = "<"
	GT = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	URI      = "@uri"
	ENDPOINT = "ENDPOINT"
	GET      = "GET"
	POST     = "POST"
	PUT      = "PUT"
	DELETE   = "DELETE"
	PATCH    = "PATCH"
	TYPE     = "TYPE"
)

var keywords = map[string]TokenType{
	"endpoint": ENDPOINT,
	"get":      GET,
	"post":     POST,
	"put":      PUT,
	"type":     TYPE,
}

func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
