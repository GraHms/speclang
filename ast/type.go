package ast

import "speclang/token"

type TypeStatement struct {
	Token       token.Token
	Name        *Identifier
	Fields      []*FieldDeclaration
	Description *StringLiteral
}

func (ts *TypeStatement) String() string {
	return ts.TokenLiteral() + " " + ts.Name.String() + " = " + ts.Description.String() + ";"
}

type FieldDeclaration struct {
	Name        *Identifier
	Type        *Identifier
	Description *StringLiteral
}

// TokenLiteral Implement the Node interface for TypeStatement
func (ts *TypeStatement) TokenLiteral() string {
	return ts.Token.Literal
}

func (ts *TypeStatement) statementNode() {}

// TokenLiteral Implement the Node interface for FieldDeclaration
func (fd *FieldDeclaration) TokenLiteral() string {
	return fd.Name.Token.Literal
}

func (fd *FieldDeclaration) statementNode() {}

type NestedTypeStatement struct {
	Token  token.Token         // The token associated with the nested type statement
	Fields []*FieldDeclaration // The fields of the nested type
}

func (nts *NestedTypeStatement) TokenLiteral() string {
	return nts.Token.Literal
}

func (nts *NestedTypeStatement) statementNode() {}
func (nts *NestedTypeStatement) String() string {
	return nts.TokenLiteral()
}
