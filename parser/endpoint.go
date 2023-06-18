package parser

import (
	"speclang/ast"
	"speclang/token"
)

func (p *Parser) parseEndpointStatement() ast.Statement {
	stmt := &ast.EndpointStatement{Token: p.curTok}

	recentAnnotation := p.mostRecentAnnotation
	stmt.URI = recentAnnotation.Value.(*ast.StringLiteral)
	// Parse the endpoint name
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}

	// Parse the endpoint block
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	stmt.Block = p.parseEndpointBlock()

	return stmt
}
