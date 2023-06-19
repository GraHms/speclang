package parser

import (
	"speclang/ast"
	"speclang/lexer"
	"speclang/token"
)

type Parser struct {
	lexer                *lexer.Lexer
	errors               []string
	curTok               token.Token
	peekTok              token.Token
	mostRecentAnnotation *ast.AnnotationStatement // Track most recent annotation statement
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{lexer: lexer, errors: []string{}, mostRecentAnnotation: nil}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curTok = p.peekTok
	p.peekTok = p.lexer.NextToken()
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curTok.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.expectPeek(token.NEWLINE)
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curTok.Type {
	case token.ANNOTATION:
		return p.parseAnnotationStatement()
	case token.ENDPOINT:
		return p.parseEndpointStatement()
	case token.GET:
		return p.parseFunctionStatement()
	case token.POST:
		return p.parseFunctionStatement()
	case token.PUT:
		return p.parseFunctionStatement()
	case token.DELETE:
		return p.parseFunctionStatement()
	case token.PATCH:
		return p.parseFunctionStatement()
	case token.TYPE:
		return p.parseTypeStatement()

	default:
		p.addError("Expected statement, got " + p.curTok.Literal + " instead")
		return nil
	}

}

func (p *Parser) parseAnnotationStatement() ast.Statement {
	stmt := &ast.AnnotationStatement{Token: p.curTok}

	// Parse the annotation name
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}

	// Parse the annotation value
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	// Parse the annotation argument
	p.nextToken()
	if p.curTok.Type != token.STRING {
		p.peekError(token.STRING)
		return nil
	}
	stmt.Value = &ast.StringLiteral{Token: p.curTok, Value: p.curTok.Literal}

	// Ensure the annotation argument is followed by a closing parenthesis
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	p.mostRecentAnnotation = stmt
	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curTok.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekTok.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekTokenIs(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
	msg := "expected next token to be %s, got %s instead"
	p.errors = append(p.errors, msg)
}

func (p *Parser) addError(msg string) {
	p.errors = append(p.errors, msg)
}

func (p *Parser) Errors() []string {
	return p.errors
}
