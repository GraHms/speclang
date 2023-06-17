package main

import (
	"fmt"
	"speclang/lexer"
	"speclang/token"
)

func main() {
	lex := lexer.New(input)
	tokens := []token.Token{}
	for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
		tokens = append(tokens, tok)
		fmt.Printf("%+v\n", tok)
	}
}

const input = `@uri("http://example.com/thing")
endpoint Product {
`
