package parser

import (
	"speclang/ast"
	"speclang/lexer"

	"testing"
)

func TestParseAnnotationStatement(t *testing.T) {
	input := `@author("John Doe")
`

	l := lexer.New(input)
	p := New(l)

	program := p.Parse()
	if program == nil {
		t.Fatalf("Parse returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.AnnotationStatement)
	if !ok {
		t.Fatalf("Expected AnnotationStatement, got %T", program.Statements[0])
	}

	expectedName := "author"
	if stmt.Name.Value != expectedName {
		t.Errorf("Expected name '%s', got '%s'", expectedName, stmt.Name.Value)
	}

	expectedValue := "John Doe"
	if stmt.Value.TokenLiteral() != expectedValue {
		t.Errorf("Expected value '%s', got '%s'", expectedValue, stmt.Value.TokenLiteral())
	}
}

func TestParseEndpointStatement(t *testing.T) {
	input := `@uri("/api/products")
endpoint Product {
	@uri("/")
	get allProducts() {
		# Define the query for retrieving all products
	}

	@uri("/{id}")
	post addProduct(@body ProductInput input) {
		# Define the query for adding a new product
	}
}`

	l := lexer.New(input)
	p := New(l)
	program := p.Parse()

	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement. Got %d", len(program.Statements))
	}

	endpointStmt, ok := program.Statements[0].(*ast.EndpointStatement)
	if !ok {
		t.Fatalf("Statement is not an EndpointStatement. Got %T", program.Statements[0])
	}

	// Test URI
	expectedURI := "/api/products"
	if endpointStmt.URI.Value != expectedURI {
		t.Errorf("Expected URI to be %q. Got %q", expectedURI, endpointStmt.URI.Value)
	}

	// Test Name
	expectedName := "Product"
	if endpointStmt.Name.Value != expectedName {
		t.Errorf("Expected Name to be %q. Got %q", expectedName, endpointStmt.Name.Value)
	}

	// Test Block
	expectedNumFunctions := 2
	if len(endpointStmt.Block.Statements) != expectedNumFunctions {
		t.Fatalf("Expected %d functions in block. Got %d", expectedNumFunctions, len(endpointStmt.Block.Statements))
	}

	// Test first function in block
	firstFunc, ok := endpointStmt.Block.Statements[0].(*ast.FunctionStatement)
	if !ok {
		t.Fatalf("First statement in block is not a FunctionStatement. Got %T", endpointStmt.Block.Statements[0])
	}
	expectedFirstFuncName := "allProducts"
	if firstFunc.Name.Value != expectedFirstFuncName {
		t.Errorf("Expected first function name to be %q. Got %q", expectedFirstFuncName, firstFunc.Name.Value)
	}

	// Test second function in block
	secondFunc, ok := endpointStmt.Block.Statements[1].(*ast.FunctionStatement)
	if !ok {
		t.Fatalf("Second statement in block is not a FunctionStatement. Got %T", endpointStmt.Block.Statements[1])
	}
	expectedSecondFuncName := "addProduct"
	if secondFunc.Name.Value != expectedSecondFuncName {
		t.Errorf("Expected second function name to be %q. Got %q", expectedSecondFuncName, secondFunc.Name.Value)
	}
}
