package translator_test

// Note: Test tests depend on a functioning parser, test it first.

import (
	"reflect"
	"strings"
	"testing"

	"github.com/bkidney/EQ2Dot/parser"
	"github.com/bkidney/EQ2Dot/translator"
)

func TestTranlator_ProduceDot(t *testing.T) {
	var tests = []struct {
		s   string
		out string
	}{
		// Simple pattern match
		{
			s:   `?node:ip:send(?srcIP,?destIP) and destIP in [203.0.113.12,192.168.1.100]`,
			out: "digraph query {\n\tnode [shape = none];\n\t0 [label = \"start\"];\n\tnode [shape = circle];\n\t0 -> 1;\n\t1 -> 2 [label = \"?node:ip:send(?srcIP,?destIP)\"];\n\t2 -> 3 [label = \"destIP in [203.0.113.12,192.168.1.100]\"];\n\t3 [shape = doublecircle];\n}",
		},
		// A more complete example
		{
			s: `
		 {DBServerNode:myDB:openSession(_):?sessionID Within
		  DBServerNode:myDb:userAuthenticate:(?user)} Precedes
			DBServerNode:myDB:sqlQuery(sessionID):?resultData Precedes
			  ?egressNode:ip::send(?outData,203.0.113.12)
		 and resultData FlowsTo* outData
		 `,
			out: "digraph query {\n\tnode [shape = none];\n\t0 [label = \"start\"];\n\tnode [shape = circle];\n\t0 -> 1;\n\t1 -> 2 [label = \"Call - DBServerNode:myDB:openSession(_):?sessionID\"];\n\tsubgraph cluster_0 {\n\trank = same;\n\tstyle=\"dashed\";\n\t2 -> 3 [label = \"DBServerNode:myDb:userAuthenticate:(?user)\"];\n\t}\t3 -> 4 [label = \"Ret - DBServerNode:myDB:openSession(_):?sessionID\"];\n\t4 -> 5 [label = \"DBServerNode:myDB:sqlQuery(sessionID):?resultData\"];\n\t5 -> 6 [label = \"?egressNode:ip::send(?outData,203.0.113.12)\"];\n\t6 -> 7 [label = \"resultData FlowsTo* outData\"];\n\t7 [shape = doublecircle];\n}",
		},
		// A simple example using 'or' with simplified identifiers.
		{
			s:   `a or {b and c}`,
			out: "digraph query {\n\tnode [shape = none];\n\t0 [label = \"start\"];\n\tnode [shape = circle];\n\t0 -> 1;\n\t1 -> 3 [label = \"a\"];\n\t1 -> 2 [label = \"b\"];\n\t2 -> 3 [label = \"c\"];\n\t3 [shape = doublecircle];\n}",
		},
	}

	for i, tt := range tests {
		ast, _ := parser.NewParser(strings.NewReader(tt.s)).Parse()
		out := translator.NewTranslator(ast).Dot()
		if !reflect.DeepEqual(tt.out, out) {
			t.Errorf("%d. %q\noutput mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.out, out)
		}
	}
}

// errstring returns the string representation of an error
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
