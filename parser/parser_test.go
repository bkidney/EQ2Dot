package parser_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/bkidney/EQ2Dot/parser"
)

func TestParser_ParseQuery(t *testing.T) {
	var tests = []struct {
		s   string
		out string
		err string
	}{
		// Simple pattern match
		{
			s:   `?node:ip:send(?srcIP,?destIP) and destIP in [203.0.113.12,192.168.1.100]`,
			out: `IDENT LOGICAL IDENT CONDITION IDENT`,
		},
		// A more complete example
		{
			s: `
DBServerNode:myDB:openSession(_):?sessionID Within
    DBServerNode:myDb:userAuthenticate:(?user) Precedes
        DBServerNode:myDB:sqlQuery(sessionID):?resultData Precedes
          ?egressNode:ip::send(?outData,203.0.113.12)
and resultData FlowsTo* outData
      `,
			out: `IDENT TEMPORAL IDENT TEMPORAL IDENT TEMPORAL IDENT LOGICAL IDENT FLOW IDENT`,
		},
		// Errors
		{
			s:   `and ?srcIP`,
			out: ``,
			err: `found "and", expected IDENT`,
		},
		{
			s:   `?node:ip:send(?srcIP,?destIP) and in [203.0.113.12,192.168.1.100]`,
			out: ``,
			err: `found "in", expected IDENT`,
		},
	}

	for i, tt := range tests {
		out, err := parser.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.out, out) {
			t.Errorf("%d. %q\noutput mismatch:\n\nexp%#v\n\ngot=%#v\n\n", i, tt.s, tt.out, out)
		}
	}
}

// errstring returns the string representation of an error
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
,
