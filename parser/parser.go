package parser

import (
  "io"
  "fmt"
  "github.com/bkidney/gofelex"
)

type Parser struct {
  s *gofelex.Scanner
  buf struct {
    tok gofelex.Token // Last read token
    lit string // Last read literal
    n int // Buffer Size (max = 1)
  }
}

func NewParser(r io.Reader) *Parser {
  return &Parser{s: gofelex.NewScanner(r)}
}

func (p *Parser) Parse() (string, error) {
  p.scanIgnoreWhitespace()
  return p.query()
}

func (p *Parser) query() (string, error) {
  var out string
  var err error

  if p.buf.tok == gofelex.IDENT {
    p.scanIgnoreWhitespace()
    out, err = p.action()
  } else {
    err = fmt.Errorf("found %q, expected IDENT", p.buf.lit)
    return "", err
  }

  return out, nil 
}

func (p *Parser) action() (string, error) {
  return "", nil
}

func (p *Parser) join() (string, error) {
  return "", nil
}

func (p *Parser) logical() (string, error) {
  return "", nil
}

func (p *Parser) temporal() (string, error) {
  return "", nil
}

func (p *Parser) conditional() (string, error) {
  return "", nil
}

func (p *Parser) flow() (string, error) {
  return "", nil
}

// Utility Functions
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
