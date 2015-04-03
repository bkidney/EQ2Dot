package syntaxTree

import "github.com/bkidney/gofelex"

// A tempory struct to hold the node type and string
//  This will be factored away with a more generic implementation of Gofelex.

type SyntaxNode struct {
	NodeType gofelex.Token
	Literal  string
	TypeStr  string
}

func (node *SyntaxNode) String() string {
	return node.TypeStr
}
