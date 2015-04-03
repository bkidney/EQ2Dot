package syntaxTree

type SyntaxTree struct {
	node    SyntaxNode
	sibling *SyntaxTree
	child   *SyntaxTree
}

func New(node SyntaxNode) *SyntaxTree {
	tree := &SyntaxTree{}
	tree.node = node
	tree.sibling = nil
	tree.child = nil

	return tree
}

func (tree *SyntaxTree) InsertChild(child *SyntaxTree) {
	if tree.child == nil {
		tree.child = child
	} else {
		temp := tree.child
		tree.child = child
		child.sibling = temp
	}
}

func (tree *SyntaxTree) GetChild() *SyntaxTree {
	return tree.child
}

func (tree *SyntaxTree) GetSibling() *SyntaxTree {
	return tree.sibling
}

// Returns depth first walk of tree (node must implement String())
func (tree *SyntaxTree) String() (out string) {
	out = tree.node.String()
	if tree.child != nil {
		out = out + " ( " + tree.child.String() + " )"
	}
	if tree.sibling != nil {
		out = out + " " + tree.sibling.String()
	}
	return
}
