package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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
			var functionParam string
			for _, innerStmt := range stmt.Block.Statements {

				switch innerStmt := innerStmt.(type) {
				case *ast.AnnotationStatement:
					if innerStmt.Name.Value == "uri" {
						functionParam = innerStmt.Value.(*ast.StringLiteral).Value
					}
				case *ast.FunctionStatement:
					sb.WriteString("    " + strings.ToLower(innerStmt.Token.Literal) + ":\n")
					sb.WriteString("      summary: " + innerStmt.Name.Value + "\n")
					sb.WriteString("      operationId: " + innerStmt.Name.Value + "\n")
					// parse function parameter if {VALUE} is present
					if strings.Contains(functionParam, "{") {
						parameters := parsePathParameters(functionParam)
						sb.WriteString("      parameters:\n")
						for _, param := range parameters {
							sb.WriteString("      - name: " + param + "\n")
							sb.WriteString("        in: path\n")
							sb.WriteString("        required: true\n")
							sb.WriteString("        type: string\n")
						}
					}
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

func parsePathParameters(uri string) []string {
	re := regexp.MustCompile(`{([^}]*)}`)
	matches := re.FindAllStringSubmatch(uri, -1)
	parameters := make([]string, len(matches))
	for i, match := range matches {
		parameters[i] = match[1]
	}
	return parameters
}
