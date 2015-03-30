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
  var out, ret string
  var err error

  err = nil

  if p.buf.tok == gofelex.IDENT {
    out, err = p.action()
  } else {
    err = fmt.Errorf("found %q, expected IDENT", p.buf.lit)
    return "", err
  }

  ret = out

  if p.buf.tok != gofelex.EOF {
    out, err = p.join()
    ret = ret + " " + out
  }

  return ret, err 
}

func (p *Parser) action() (string, error) {
  var out string
  var err error

  out = "IDENT"
  err = nil

  p.scanIgnoreWhitespace()
  return out, err
}

func (p *Parser) join() (string, error) {
  var ret, out string
  var err error

  if p.buf.tok == gofelex.LOGICAL {
    out, err = p.logical()
  } else if p.buf.tok == gofelex.TEMPORAL {
    out, err = p.temporal()
  } else if p.buf.tok == gofelex.CONDITION {
    out, err = p.conditional()
  } else if p.buf.tok == gofelex.FLOW {
    out, err = p.flow()
  } else {
    err = fmt.Errorf("found %q, expected join type")
    return "", err
  }

  ret = out
  out, err = p.query()

  ret = ret + " " + out

  return ret, err
}

func (p *Parser) logical() (string, error) {
  var out string
  var err error

  out = "LOGICAL"
  err = nil

  p.scanIgnoreWhitespace()
  return out, err
}

func (p *Parser) temporal() (string, error) {
  var out string
  var err error

  out = "TEMPORAL"
  err = nil

  p.scanIgnoreWhitespace()
  return out, err
}

func (p *Parser) conditional() (string, error) {
  var out string
  var err error

  out = "CONDITION"
  err = nil

  p.scanIgnoreWhitespace()
  return out, err
}

func (p *Parser) flow() (string, error) {
  var out string
  var err error

  out = "FLOW"
  err = nil

  p.scanIgnoreWhitespace()
  return out, err
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
