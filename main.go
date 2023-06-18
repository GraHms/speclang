package main

import (
	"io"
	"speclang/lexer"
	"speclang/parser"
)

func main() {
	lex := lexer.New(input)

	p := parser.New(lex)
	program := p.Parse()

	for _, stmt := range program.Statements {
		println(stmt.String())
	}

}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {

		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}

const input = `@uri("http://example.com/thing")
			  endpoint Product {
				@uri("/")
				get allProducts() {
					# Define the query for retrieving all products
				}
			
				@uri("/{id}")
				post product() {
                   # Define the query for retrieving a single product
				}
				
			}
			type Product {
				name string "binding:json, required"
				description string "binding:json, required"
				price float64 "binding:json, required"
			}`
