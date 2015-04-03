package syntaxTree

type SyntaxTree struct {
	node    string
	sibling *SyntaxTree
	child   *SyntaxTree
}

func New(node string) *SyntaxTree {
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

// Returns depth first walk of tree (node must implement String())
func (tree *SyntaxTree) String() (out string) {
	out = tree.node
	if tree.child != nil {
		out = out + " ( " + tree.child.String() + " )"
	}
	if tree.sibling != nil {
		out = out + " " + tree.sibling.String()
	}
	return
}
