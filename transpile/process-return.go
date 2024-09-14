package transpile

import (
	"github.com/abibby/abc/parser"
)

func processReturn(nodes []parser.Node) ([]parser.Node, error) {
	newNodes := make([]parser.Node, 0, len(nodes))

	for _, n := range nodes {
		newNode, err := processReturnNode(n)
		if err != nil {
			return nil, err
		}
		newNodes = append(newNodes, newNode)
	}
	return newNodes, nil
}

func processReturnNode(node parser.Node) (parser.Node, error) {
	switch n := node.(type) {
	case *parser.FunctionDefNode:
		return processReturnFunctionDef(n)
	default:
		return n, nil
	}
}
func processReturnFunctionDef(node *parser.FunctionDefNode) (*parser.FunctionDefNode, error) {
	ret, ok := node.ReturnType.(*parser.BasicTypeNode)
	if !ok || ret.Value != "void" {
		return node, nil
	}

	statements := make([]parser.Node, len(node.Block.Statements)+1)
	copy(node.Block.Statements, statements)
	statements[len(statements)-1] = &parser.StatementNode{Value: &parser.ReturnNode{}}

	newBlock := &parser.BlockNode{
		LocationNode: node.LocationNode,
		Statements:   statements,
	}

	return &parser.FunctionDefNode{
		LocationNode: node.LocationNode,
		ReturnType:   node.ReturnType,
		FunctionName: node.FunctionName,
		Arguments:    node.Arguments,
		Block:        newBlock,
	}, nil
}
