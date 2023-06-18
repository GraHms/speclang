package parser

import (
	"speclang/ast"
	"speclang/token"
)

// ... (Previous code)

func (p *Parser) parseFunctionStatement() ast.Statement {
	stmt := &ast.FunctionStatement{Token: p.curTok}

	// Parse the function name
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}

	// Parse the function parameters
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	stmt.Parameters = p.parseFunctionParameters()

	// Parse the function body
	stmt.Body = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseFunctionParameters() []*ast.Parameter {
	parameters := []*ast.Parameter{}

	// Check for empty parameter list
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return parameters
	}

	// Parse the first parameter
	p.nextToken()
	param := p.parseParameter()
	parameters = append(parameters, param)

	// Parse additional parameters if any
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		param = p.parseParameter()
		parameters = append(parameters, param)
	}

	// Ensure the parameter list is closed with a closing parenthesis
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return parameters
}

func (p *Parser) parseParameter() *ast.Parameter {
	param := &ast.Parameter{Token: p.curTok}

	// Parse the parameter name
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	param.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}

	// Parse the parameter type
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	param.Type = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}

	return param
}

// ... (Remaining code)
