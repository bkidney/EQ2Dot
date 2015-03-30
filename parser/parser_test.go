package parser_test

import (
  "strings"
  "testing"
  "reflect"

  "github.com/bkidney/EQ2Dot/parser"
)

func TestParser_ParseQuery(t *testing.T) {
  var tests = []struct {
    s string
    err string
  }{
    // Simple pattern match 
    {
      s: `?node:ip:send(?srcIP,?destIP) and destIP in [203.0.113.12,192.168.1.100]`,
    },
    // Errors
    {
      s: `and ?srcIP`,
      err: `found "and", expected IDENT`,
    },
    {
      s: `?node:ip:send(?srcIP,?destIP) and in [203.0.113.12,192.168.1.100]`,
    },
  }

  for i, tt := range tests {
    _, err := parser.NewParser(strings.NewReader(tt.s)).Parse()
    if !reflect.DeepEqual(tt.err, errstring(err)) {
      t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
    }
  }
}

// errstring returns the string representation of an error
func errstring (err error) string {
  if err != nil {
    return err.Error()
  }
  return ""
}
