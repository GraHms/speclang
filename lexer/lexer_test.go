package lexer

import (
	"speclang/token"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `@uri("http://example.com/thing")
endpoint Product {

	@uri("/")
	get allProducts() {

		// Define the query for retrieving all products
	}

	@uri("/{id}")
	post addProduct(@body ProductInput input) {
		// Define the query for adding a new product
	}
}

type Product {
	name string "binding:json, required"
	description string "binding:json, required"
	price float64 "binding:json, required"
}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{

		{token.ANNOTATION, "@"},
		{token.IDENT, "uri"},
		{token.LPAREN, "("},
		{token.STRING, "http://example.com/thing"},
		{token.RPAREN, ")"},
		//{token.ENDPOINT, "endpoint"},
		//{token.IDENT, "Product"},
		//{token.LBRACE, "{"},
		//{token.ANNOTATION, "@uri"},
		//{token.STRING, `"/"`},
		//{token.GET, "get"},
		//{token.IDENT, "allProducts"},
		//{token.LPAREN, "("},
		//{token.RPAREN, ")"},
		//{token.LBRACE, "{"},
		//{token.RBRACE, "}"},
		//{token.ANNOTATION, "@uri"},
		//{token.STRING, `"/{id}"`},
		//{token.POST, "post"},
		//{token.IDENT, "addProduct"},
		//{token.LPAREN, "("},
		//{token.IDENT, "ProductInput"},
		//{token.IDENT, "input"},
		//{token.RPAREN, ")"},
		//{token.LBRACE, "{"},
		//{token.RBRACE, "}"},
		//{token.TYPE, "type"},
		//{token.IDENT, "Product"},
		//{token.LBRACE, "{"},
		//{token.IDENT, "name"},
		//{token.STRING, `"binding:json, required"`},
		//{token.IDENT, "string"},
		//{token.IDENT, "description"},
		//{token.STRING, `"binding:json, required"`},
		//{token.IDENT, "string"},
		//{token.IDENT, "price"},
		//{token.FLOAT, "float64"},
		//{token.STRING, `"binding:json, required"`},
		//{token.RBRACE, "}"},
		//{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("test[%d] - TokenType wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - Literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}
