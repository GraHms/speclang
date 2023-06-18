package ast

import "speclang/token"

type FunctionStatement struct {
	Token      token.Token // The token.FUNCTION token
	Name       *Identifier
	Parameters []*Parameter
	Body       *BlockStatement
}

func (f FunctionStatement) TokenLiteral() string {
	return f.Token.Literal
}

func (f FunctionStatement) String() string {
	var out string
	out += f.TokenLiteral() + " "
	out += f.Name.String()
	out += "("
	for i, p := range f.Parameters {
		if i > 0 {
			out += ", "
		}
		out += p.Name.String()
		out += ": "
		out += p.Type.String()
	}
	out += ") "
	out += f.Body.String()
	return out
}

func (f FunctionStatement) statementNode() {
}

type Parameter struct {
	Token token.Token // The token.IDENT token
	Name  *Identifier
	Type  *Identifier
}
