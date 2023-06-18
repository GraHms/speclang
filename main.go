package main

import (
	"fmt"
	"speclang/ast"
	"speclang/lexer"
	"speclang/parser"
)

func main() {
	input := `
	@uri("/products")
	endpoint Product {
		@uri("/")
		get AllProducts() {
			// Retrieve all products
		}

		@uri("/{id}")
		get ProductById(id int) {
			// Retrieve product by ID
		}

		@uri("/")
		post addProduct(product ProductInput) {
			// Add a new product
		}

		@uri("/{id}")
		patch updateProductById(id int, product ProductInput) {
			// Update product by ID
		}

		@uri("/{id}")
		delete ProductById(id int) {
			// Delete product by ID
		}
	}
	`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.Parse()

	var currentEndpointURI string
	var currentFunctionURI string
	for _, stmt := range program.Statements {
		switch stmt := stmt.(type) {
		case *ast.AnnotationStatement:
			if stmt.Name.Value == "uri" {
				currentEndpointURI = stmt.Value.(*ast.StringLiteral).Value
			}
		case *ast.EndpointStatement:
			endpointURI := currentEndpointURI
			fmt.Printf("Endpoint: %s\n", endpointURI)
			fmt.Printf("Endpoint: %s\n", stmt.URI.Value)
			for _, innerStmt := range stmt.Block.Statements {

				switch innerStmt := innerStmt.(type) {
				case *ast.AnnotationStatement:
					if innerStmt.Name.Value == "uri" {
						currentFunctionURI = innerStmt.Value.(*ast.StringLiteral).Value
					}
				case *ast.FunctionStatement:
					fmt.Printf("METHOD: %s, FUNCTION: %s, URI: %s%s\n", innerStmt.Token.Literal, innerStmt.Name.Value, endpointURI, currentFunctionURI)
				}
			}
		}
	}
}
