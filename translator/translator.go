package translator

import (
	"strconv"
	"strings"

	"github.com/bkidney/EQ2Dot/syntaxTree"
	"github.com/bkidney/gofelex"
)

type Translator struct {
	ast            *syntaxTree.SyntaxTree
	nodeNumber     int
	subgraphNumber int
}

func NewTranslator(ast *syntaxTree.SyntaxTree) *Translator {
	translator := &Translator{}
	translator.ast = ast
	translator.nodeNumber = 2 // 2 is first node after query
	translator.subgraphNumber = 0

	return translator
}

func (translator *Translator) Dot() string {
	var out string

	out = translator.generatePreamble()
	out = out + translator.generateDot(translator.ast.GetChild())
	out = out + translator.generatePostamble()
	return out
}

func (translator *Translator) generatePreamble() string {
	out := "digraph query {\n\tnode [shape = none];\n\t0 [label = \"start\"];\n\tnode [shape = circle];\n\t0 -> 1;\n"
	return out
}

func (translator *Translator) generatePostamble() string {
	lastNode := translator.getLastNode()
	return "\t" + strconv.Itoa(lastNode) + " [shape = doublecircle];\n}"
}

func (translator *Translator) generateDot(ast *syntaxTree.SyntaxTree) string {
	var out string
	var currNode = ast.GetNode()

	// Here we look at the node
	//  If binary operator, get child translation and child.sibling translation
	//  Combine as necessay
	//  Withins are boxed
	if currNode.NodeType == gofelex.TEMPORAL {
		temporalType := strings.ToUpper(currNode.Literal)
		if temporalType == "WITHIN" {
			out = translator.createSubGraph(ast.GetChild())
		} else {
			lhs := ast.GetChild()

			if lhs.GetNode().NodeType == gofelex.IDENT {
				edge := ast.GetChild().GetNode().Literal
				nodeNum := translator.getNextNode()
				next := translator.generateDot(ast.GetChild().GetSibling())

				out = translator.createNode(edge, nodeNum) + next
			} else {
				out = translator.generateDot(lhs) + translator.generateDot(lhs.GetSibling())
			}
		}

	}
	// Logical connections
	//  and is literal -> literal
	//  or is parallel linking
	if currNode.NodeType == gofelex.LOGICAL {
		if currNode.Literal == "and" {
			edge := ast.GetChild().GetNode().Literal
			nodeNum := translator.getNextNode()
			next := translator.generateDot(ast.GetChild().GetSibling())

			out = translator.createNode(edge, nodeNum) + next
		} else {
			edge := ast.GetChild().GetNode().Literal
			lastNode := translator.getLastNode()

			next := translator.generateDot(ast.GetChild().GetSibling())
			siblingLastNode := translator.getLastNode()

			out = translator.createNodeEx(edge, lastNode, siblingLastNode) + next
		}
	}
	//  flow combines literals
	if currNode.NodeType == gofelex.FLOW {
		lhs := ast.GetChild().GetNode().Literal
		op := currNode.Literal
		rhs := ast.GetChild().GetSibling().GetNode().Literal

		edge := lhs + " " + op + " " + rhs

		out = translator.createNode(edge, translator.getNextNode())
	}
	//  condition combinines literal
	if currNode.NodeType == gofelex.CONDITION {
		lhs := ast.GetChild().GetNode().Literal
		op := currNode.Literal
		rhs := ast.GetChild().GetSibling().GetNode().Literal

		edge := lhs + " " + op + " " + rhs

		out = translator.createNode(edge, translator.getNextNode())
	}

	if currNode.NodeType == gofelex.IDENT {
		nodeNum := translator.getNextNode()
		out = translator.createNode(currNode.Literal, nodeNum)
	}

	return out
}

func (translator *Translator) createNode(edge string, nodeNum int) string {
	return translator.createNodeEx(edge, nodeNum-1, nodeNum)
}

func (translator *Translator) createNodeEx(edge string, startNode int, endNode int) string {
	var out string

	out = "\t" + strconv.Itoa(startNode) + " -> " + strconv.Itoa(endNode) + " [label = \"" + edge + "\"];\n"

	return out
}

func (translator *Translator) createSubGraph(ast *syntaxTree.SyntaxTree) string {
	var out string
	enterEdge := "Call - " + ast.GetNode().Literal
	exitEdge := "Ret - " + ast.GetNode().Literal

	out = translator.createNode(enterEdge, translator.getNextNode())
	out = out + "\tsubgraph cluster_" + strconv.Itoa(translator.getSubgraphNum()) + " {\n\trank = same;\n\tstyle=\"dashed\";\n"
	out = out + translator.generateDot(ast.GetSibling())
	out = out + "\t}"

	out = out + translator.createNode(exitEdge, translator.getNextNode())

	return out
}

func (translator *Translator) getNextNode() int {
	ret := translator.nodeNumber
	translator.nodeNumber++
	return ret
}

func (translator *Translator) getLastNode() int {
	return translator.nodeNumber - 1
}

func (translator *Translator) getSubgraphNum() int {
	ret := translator.subgraphNumber
	translator.subgraphNumber++
	return ret
}
