package parser

import (
	"fmt"
	"io"

	"github.com/bkidney/EQ2Dot/syntaxTree"
	"github.com/bkidney/gofelex"
)

type Parser struct {
	s   *gofelex.Scanner
	buf struct {
		tok gofelex.Token // Last read token
		lit string        // Last read literal
		n   int           // Buffer Size (max = 1)
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: gofelex.NewScanner(r)}
}

func (p *Parser) Parse() (string, error) {
	ast := syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "root"})

	p.scanIgnoreWhitespace()
	tree, err := p.query()
	ast.InsertChild(tree)

	return ast.String(), err
}

func (p *Parser) query() (*syntaxTree.SyntaxTree, error) {
	var out, ret *syntaxTree.SyntaxTree
	var err error

	err = nil

	if p.buf.tok == gofelex.IDENT || p.buf.tok == gofelex.OPEN {
		out, err = p.action()
	} else {
		err = fmt.Errorf("found %q, expected IDENT or OPEN", p.buf.lit)
		return nil, err
	}

	ret = out

	if isJoin(p.buf.tok) {
		out, err = p.join()
		out.InsertChild(ret)
		ret = out
	}

	return ret, err
}

func (p *Parser) action() (*syntaxTree.SyntaxTree, error) {
	var ret *syntaxTree.SyntaxTree
	var err error

	err = nil

	if p.buf.tok == gofelex.IDENT {
		node := syntaxTree.SyntaxNode{
			NodeType: p.buf.tok,
			Literal:  p.buf.lit,
			TypeStr:  "IDENT",
		}

		ret = syntaxTree.New(node)
	} else {

		p.scanIgnoreWhitespace()
		ret, err = p.query()

		if p.buf.tok != gofelex.CLOSE {
			err = fmt.Errorf("found %q, expected CLOSE", p.buf.lit)
		}
	}

	p.scanIgnoreWhitespace()
	return ret, err
}

func (p *Parser) join() (*syntaxTree.SyntaxTree, error) {
	var ret, out *syntaxTree.SyntaxTree
	var err error

	if p.buf.tok == gofelex.LOGICAL {
		ret, err = p.logical()
	} else if p.buf.tok == gofelex.TEMPORAL {
		ret, err = p.temporal()
	} else if p.buf.tok == gofelex.CONDITION {
		ret, err = p.conditional()
	} else if p.buf.tok == gofelex.FLOW {
		ret, err = p.flow()
	} else {
		err = fmt.Errorf("found %q, expected join type", p.buf.lit)
		return nil, err
	}

	out, err = p.query()

	ret.InsertChild(out)

	return ret, err
}

func (p *Parser) logical() (*syntaxTree.SyntaxTree, error) {
	var out *syntaxTree.SyntaxTree
	var err error

	node := syntaxTree.SyntaxNode{
		NodeType: p.buf.tok,
		Literal:  p.buf.lit,
		TypeStr:  "LOGICAL",
	}
	out = syntaxTree.New(node)
	err = nil

	p.scanIgnoreWhitespace()
	return out, err
}

func (p *Parser) temporal() (*syntaxTree.SyntaxTree, error) {
	var out *syntaxTree.SyntaxTree
	var err error

	node := syntaxTree.SyntaxNode{
		NodeType: p.buf.tok,
		Literal:  p.buf.lit,
		TypeStr:  "TEMPORAL",
	}
	out = syntaxTree.New(node)
	err = nil

	p.scanIgnoreWhitespace()
	return out, err
}

func (p *Parser) conditional() (*syntaxTree.SyntaxTree, error) {
	var out *syntaxTree.SyntaxTree
	var err error

	node := syntaxTree.SyntaxNode{
		NodeType: p.buf.tok,
		Literal:  p.buf.lit,
		TypeStr:  "CONDITION",
	}
	out = syntaxTree.New(node)
	err = nil

	p.scanIgnoreWhitespace()
	return out, err
}

func (p *Parser) flow() (*syntaxTree.SyntaxTree, error) {
	var out *syntaxTree.SyntaxTree
	var err error

	node := syntaxTree.SyntaxNode{
		NodeType: p.buf.tok,
		Literal:  p.buf.lit,
		TypeStr:  "FLOW",
	}
	out = syntaxTree.New(node)
	err = nil

	p.scanIgnoreWhitespace()
	return out, err
}

// Utility Functions
func isJoin(tok gofelex.Token) bool {
	return tok == gofelex.LOGICAL ||
		tok == gofelex.TEMPORAL ||
		tok == gofelex.CONDITION ||
		tok == gofelex.FLOW
}

func (p *Parser) scan() (tok gofelex.Token, lit string) {

	// Return any token in the buffer
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read a new one in
	tok, lit = p.s.Scan()

	// And save to buffer
	p.buf.tok, p.buf.lit = tok, lit

	return
}

func (p *Parser) scanIgnoreWhitespace() {
	var tok gofelex.Token

	tok, _ = p.scan()
	if tok == gofelex.WS {
		tok, _ = p.scan()
	}
}

func (p *Parser) unscan() { p.buf.n = 1 }
