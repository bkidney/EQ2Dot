package syntaxTree_test

import (
	"reflect"
	"testing"

	"github.com/bkidney/EQ2Dot/syntaxTree"
)

func TestSyntaxTree_SingleNode(t *testing.T) {
	var out string
	expected := "root"
	tree := syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: expected})

	out = tree.String()
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Error:\nexp = %q\n\ngot = %q\n\n", expected, out)
	}
}

func TestSyntaxTree_OneChild(t *testing.T) {
	var out string

	expected := "root ( a )"
	tree := syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "root"})

	tree.InsertChild(syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "a"}))

	out = tree.String()
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Error:\nexp = %q\n\ngot = %q\n\n", expected, out)
	}
}

func TestSyntaxTree_ChildAndSiblings(t *testing.T) {
	var out string

	expected := "root ( c b a )"
	tree := syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "root"})

	tree.InsertChild(syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "a"}))
	tree.InsertChild(syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "b"}))
	tree.InsertChild(syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "c"}))

	out = tree.String()
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Error:\nexp = %q\n\ngot = %q\n\n", expected, out)
	}
}

func TestSyntaxTree_MultiLevelChildAndSiblings(t *testing.T) {
	var out string

	expected := "root ( c ( c3 c2 c1 ) b a ( a2 a1 ) )"
	tree := syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "root"})

	aTree := syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "a"})
	aTree.InsertChild(syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "a1"}))
	aTree.InsertChild(syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "a2"}))

	cTree := syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "c"})
	cTree.InsertChild(syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "c1"}))
	cTree.InsertChild(syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "c2"}))
	cTree.InsertChild(syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "c3"}))

	tree.InsertChild(aTree)
	tree.InsertChild(syntaxTree.New(syntaxTree.SyntaxNode{TypeStr: "b"}))
	tree.InsertChild(cTree)

	out = tree.String()
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Error:\nexp = %q\n\ngot = %q\n\n", expected, out)
	}
}
