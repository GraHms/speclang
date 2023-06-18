package ast

import (
	"bytes"
	"speclang/token"
)

// Node represents a node in the abstract syntax tree (AST).
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents a statement in the "speclang" language.
type Statement interface {
	Node
	statementNode()
}

// Expression is the interface that all expression nodes implement.
type Expression interface {
	Node
	expressionNode()
}

// Program represents a "speclang" program.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// EndpointStatement represents an endpoint declaration in "speclang".
type EndpointStatement struct {
	Token   token.Token
	Path    string
	Methods []MethodStatement
}

func (es *EndpointStatement) statementNode()       {}
func (es *EndpointStatement) TokenLiteral() string { return es.Token.Literal }
func (es *EndpointStatement) String() string {
	var out bytes.Buffer

	out.WriteString("@uri(\"")
	out.WriteString(es.Path)
	out.WriteString("\")\n")
	out.WriteString("endpoint {\n")

	for _, method := range es.Methods {
		out.WriteString(method.String())
		out.WriteString("\n")
	}

	out.WriteString("}")

	return out.String()
}

// MethodStatement represents a method declaration within an endpoint in "speclang".
type MethodStatement struct {
	Token      token.Token
	Method     string
	Path       string
	Parameters []ParameterStatement
}

func (ms *MethodStatement) TokenLiteral() string { return ms.Token.Literal }
func (ms *MethodStatement) String() string {
	var out bytes.Buffer

	out.WriteString("\t@uri(\"")
	out.WriteString(ms.Path)
	out.WriteString("\")\n")
	out.WriteString("\t")
	out.WriteString(ms.Method)
	out.WriteString(" {\n")

	for _, param := range ms.Parameters {
		out.WriteString(param.String())
		out.WriteString("\n")
	}

	out.WriteString("\t}")

	return out.String()
}

// ParameterStatement represents a parameter declaration within a method in "speclang".
type ParameterStatement struct {
	Token token.Token
	Name  string
	Type  string
}

func (ps *ParameterStatement) TokenLiteral() string { return ps.Token.Literal }
func (ps *ParameterStatement) String() string {
	var out bytes.Buffer

	out.WriteString("\t\t")
	out.WriteString(ps.Name)
	out.WriteString(": ")
	out.WriteString(ps.Type)

	return out.String()
}

// TypeDefinition represents a type definition in "speclang".
type TypeDefinition struct {
	Token  token.Token
	Name   string
	Fields []FieldDefinition
}

func (td *TypeDefinition) statementNode()       {}
func (td *TypeDefinition) TokenLiteral() string { return td.Token.Literal }
func (td *TypeDefinition) String() string {
	var out bytes.Buffer

	out.WriteString("type ")
	out.WriteString(td.Name)
	out.WriteString(" {\n")

	for _, field := range td.Fields {
		out.WriteString(field.String())
		out.WriteString("\n")
	}

	out.WriteString("}")

	return out.String()
}

// FieldDefinition represents a field definition within a type in "speclang".
type FieldDefinition struct {
	Token token.Token
	Name  string
	Type  string
	Tags  []string
}

func (fd *FieldDefinition) TokenLiteral() string { return fd.Token.Literal }
func (fd *FieldDefinition) String() string {
	var out bytes.Buffer

	out.WriteString("\t")
	out.WriteString(fd.Name)
	out.WriteString(" ")
	out.WriteString(fd.Type)

	for _, tag := range fd.Tags {
		out.WriteString(" \"")
		out.WriteString(tag)
		out.WriteString("\"")
	}

	return out.String()
}
