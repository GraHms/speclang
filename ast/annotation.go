package ast

import "speclang/token"

type AtStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (as *AtStatement) statementNode() {}

// TokenLiteral returns the literal representation of the annotation statement's token.
func (as *AtStatement) TokenLiteral() string {
	return as.Token.Literal
}

func (as *AtStatement) String() string {
	return as.TokenLiteral() + " " + as.Name.String() + " = " + as.Value.String() + ";"
}
