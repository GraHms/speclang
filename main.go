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
				@uri("/")
				get allProducts() {
					# Define the query for retrieving all products
				}
			
				@uri("/{id}")
				post addProduct(@body ProductInput input) {
					# Define the query for adding a new product
				}
			}
			type Product {
				name string "binding:json, required"
				description string "binding:json, required"
				price float64 "binding:json, required"
			}`
