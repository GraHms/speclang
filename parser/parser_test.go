package parser

import (
	"speclang/ast"
	"speclang/lexer"

	"testing"
)

func TestParseAnnotationStatement(t *testing.T) {
	input := `@author("John Doe")`

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
