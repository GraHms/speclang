package ast

import "speclang/token"

// AnnotationStatement represents an annotation statement.
type AnnotationStatement struct {
	Token token.Token // The '@' token
	Name  *Identifier
	Value Expression
}

func (as *AnnotationStatement) String() string {
	return as.TokenLiteral() + " " + as.Name.String() + " = " + as.Value.String() + ";"
}

func (as *AnnotationStatement) statementNode() {}

// TokenLiteral returns the literal representation of the annotation statement's token.
func (as *AnnotationStatement) TokenLiteral() string {
	return as.Token.Literal
}

// Identifier represents an identifier in the Speclang code.
type Identifier struct {
	Token token.Token // The IDENT token
	Value string      // The identifier's value
}

func (i *Identifier) String() string  { return i.Value }
func (i *Identifier) expressionNode() {}

// TokenLiteral returns the literal representation of the identifier's token.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
