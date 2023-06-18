package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"speclang/ast"
	"speclang/lexer"
	"speclang/parser"
	"strings"
)

func main() {
	// Read the speclang file
	filePath := "product.speclang"
	input, err := ReadSpeclangFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the speclang input
	program := ParseSpeclang(input)

	// Generate Swagger specification
	swagger := GenerateSwaggerSpec(program)

	// Write Swagger specification to a file
	outputPath := "spec.yaml"
	err = WriteSwaggerSpec(outputPath, swagger)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Swagger specification written to %s\n", outputPath)

}

// ReadSpeclangFile reads the content of a speclang file.
func ReadSpeclangFile(filePath string) (string, error) {
	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(input), nil
}

// ParseSpeclang parses the speclang input and returns the program AST.
func ParseSpeclang(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.Parse()
}

// GenerateSwaggerSpec generates the Swagger specification from the program AST.
func GenerateSwaggerSpec(program *ast.Program) string {
	var sb strings.Builder
	sb.WriteString("swagger: '2.0'\n")
	sb.WriteString("info:\n")
	sb.WriteString("  version: 1.0.0\n")
	sb.WriteString("  title: My API\n")
	sb.WriteString("paths:\n")

	var currentEndpointURI string
	for _, stmt := range program.Statements {
		switch stmt := stmt.(type) {
		case *ast.AnnotationStatement:
			if stmt.Name.Value == "uri" {
				currentEndpointURI = stmt.Value.(*ast.StringLiteral).Value
			}
		case *ast.EndpointStatement:
			endpointURI := currentEndpointURI
			sb.WriteString("  " + endpointURI + ":\n")
			for _, innerStmt := range stmt.Block.Statements {
				switch innerStmt := innerStmt.(type) {
				case *ast.AnnotationStatement:
					if innerStmt.Name.Value == "uri" {
						// Store the function URI if needed for further processing
						// currentFunctionURI = innerStmt.Value.(*ast.StringLiteral).Value
					}
				case *ast.FunctionStatement:
					sb.WriteString("    " + strings.ToLower(innerStmt.Token.Literal) + ":\n")
					sb.WriteString("      summary: " + innerStmt.Name.Value + "\n")
					sb.WriteString("      operationId: " + innerStmt.Name.Value + "\n")
					sb.WriteString("      parameters:\n")
					// Add parameters here if needed
					sb.WriteString("      responses:\n")
					sb.WriteString("        200:\n")
					sb.WriteString("          description: Success\n")
				}
			}
		}
	}

	return sb.String()
}

// WriteSwaggerSpec writes the Swagger specification to a file.
func WriteSwaggerSpec(outputPath string, swaggerSpec string) error {
	return os.WriteFile(outputPath, []byte(swaggerSpec), 0644)
}
