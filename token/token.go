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
	AT         = "@"
	FLOAT      = "FLOAT"
	ANNOTATION = "ANNOTATION"

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
	ENDPOINT = "endpoint"
	GET      = "get"
	POST     = "post"
	TYPE     = "type"
)

var keywords = map[string]TokenType{
	"uri":      URI,
	"endpoint": ENDPOINT,
	"get":      GET,
	"post":     POST,
	"type":     TYPE,
}

// LookupIdentifier checks if the given identifier is a keyword.
// If it is, it returns the corresponding TokenType.
// Otherwise, it returns IDENT as the TokenType.
func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
