package syntaxTree_test

import (
	"reflect"
	"testing"

	"github.com/bkidney/EQ2Dot/syntaxTree"
)

func TestSyntaxTree_SingleNode(t *testing.T) {
	var out string
	expected := "root"
	tree := syntaxTree.New(expected)

	out = tree.String()
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Error:\nexp = %q\n\ngot = %q\n\n", expected, out)
	}
}

func TestSyntaxTree_OneChild(t *testing.T) {
	var out string

	expected := "root ( a )"
	tree := syntaxTree.New("root")

	tree.InsertChild(syntaxTree.New("a"))

	out = tree.String()
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Error:\nexp = %q\n\ngot = %q\n\n", expected, out)
	}
}

func TestSyntaxTree_ChildAndSiblings(t *testing.T) {
	var out string

	expected := "root ( a b c )"
	tree := syntaxTree.New("root")

	tree.InsertChild(syntaxTree.New("a"))
	tree.InsertChild(syntaxTree.New("b"))
	tree.InsertChild(syntaxTree.New("c"))

	out = tree.String()
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Error:\nexp = %q\n\ngot = %q\n\n", expected, out)
	}
}

func TestSyntaxTree_MultiLevelChildAndSiblings(t *testing.T) {
	var out string

	expected := "root ( a ( a1 a2 ) b c ( c1 c2 c3 ) )"
	tree := syntaxTree.New("root")

	aTree := syntaxTree.New("a")
	aTree.InsertChild(syntaxTree.New("a1"))
	aTree.InsertChild(syntaxTree.New("a2"))

	cTree := syntaxTree.New("c")
	cTree.InsertChild(syntaxTree.New("c1"))
	cTree.InsertChild(syntaxTree.New("c2"))
	cTree.InsertChild(syntaxTree.New("c3"))

	tree.InsertChild(aTree)
	tree.InsertChild(syntaxTree.New("b"))
	tree.InsertChild(cTree)

	out = tree.String()
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Error:\nexp = %q\n\ngot = %q\n\n", expected, out)
	}
}
