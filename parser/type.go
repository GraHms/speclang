package parser

import (
	"speclang/ast"
	"speclang/token"
)

func (p *Parser) parseFieldDeclaration() *ast.FieldDeclaration {
	decl := &ast.FieldDeclaration{}

	// Parse the field name
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	decl.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}

	// Parse the field type
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	decl.Type = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}

	// Parse the field description
	if !p.expectPeek(token.STRING) {
		return nil
	}
	decl.Description = &ast.StringLiteral{Token: p.curTok, Value: p.curTok.Literal}

	return decl
}
func (p *Parser) parseTypeStatement() *ast.TypeStatement {
	stmt := &ast.TypeStatement{Token: p.curTok}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = p.parseIdentifier()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Fields = p.parseFieldList()

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return stmt
}

func (p *Parser) parseFieldList() []*ast.FieldDeclaration {
	fields := []*ast.FieldDeclaration{}

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		field := p.parseField()
		if field != nil {
			fields = append(fields, field)
		}
		p.nextToken()
	}

	return fields
}

func (p *Parser) parseField() *ast.FieldDeclaration {
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	identifier := p.parseIdentifier()

	if !p.expectPeek(token.COLON) {
		return nil
	}

	p.nextToken() // Skip the colon

	fieldType := p.parseType()

	// Perform a type assertion to convert the fieldType to ast.Identifier
	typeIdentifier, ok := fieldType.(*ast.Identifier)
	if !ok {
		p.peekError(token.IDENT)
	}

	return &ast.FieldDeclaration{Name: identifier, Type: typeIdentifier}

}

func (p *Parser) parseType() ast.Node {
	if p.curTokenIs(token.IDENT) {
		return p.parseIdentifier()
	}

	// Handle nested type statements
	if p.curTokenIs(token.LBRACE) {
		return p.parseNestedTypeStatement()
	}

	p.peekError(token.IDENT)
	return nil
}

func (p *Parser) parseNestedTypeStatement() *ast.NestedTypeStatement {
	stmt := &ast.NestedTypeStatement{Token: p.curTok}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Fields = p.parseFieldList()

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return stmt
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	if p.curTok.Type == token.IDENT {
		identifier := &ast.Identifier{
			Token: p.curTok,
			Value: p.curTok.Literal,
		}
		p.nextToken()
		return identifier
	}
	return nil
}
