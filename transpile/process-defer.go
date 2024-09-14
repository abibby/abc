package transpile

import (
	"github.com/abibby/abc/parser"
)

type deferredStatements struct {
	statements []*parser.StatementNode
}

func newDeferredStatements() *deferredStatements {
	return &deferredStatements{}
}

func (s *deferredStatements) Add(statement *parser.StatementNode) {
	s.statements = append(s.statements, statement)
}
func (s *deferredStatements) Nodes() []parser.Node {
	nodes := make([]parser.Node, len(s.statements))
	for i, s := range s.statements {
		nodes[i] = s
	}
	return nodes
}

func processDefer(nodes []parser.Node) ([]parser.Node, error) {
	newNodes := make([]parser.Node, 0, len(nodes))

	for _, n := range nodes {
		newNode, err := processDeferNode(n, newDeferredStatements())
		if err != nil {
			return nil, err
		}
		newNodes = append(newNodes, newNode)
	}
	return newNodes, nil
}

func processDeferNode(node parser.Node, deferredStatements *deferredStatements) (parser.Node, error) {
	switch n := node.(type) {
	case *parser.BlockNode:
		return processDeferBlock(n, deferredStatements)
	case *parser.FunctionDefNode:
		return processDeferFunctionDef(n, deferredStatements)
	default:
		return n, nil
	}
}
func processDeferBlock(node *parser.BlockNode, deferredStatements *deferredStatements) (*parser.BlockNode, error) {
	statements := []parser.Node{}

	for _, s := range node.Statements {
		if sn, ok := s.(*parser.StatementNode); ok {
			s = sn.Value
		}
		switch s := s.(type) {
		case *parser.DeferNode:
			deferredStatements.Add(s.Statement)
		case *parser.BlockNode:
			block, err := processDeferBlock(s, deferredStatements)
			if err != nil {
				return nil, err
			}
			statements = append(statements, block)
		case *parser.ReturnNode:
			statements = append(statements, deferredStatements.Nodes()...)
			statements = append(statements, s)
		default:
			statements = append(statements, s)
		}
	}

	return &parser.BlockNode{
		LocationNode: node.LocationNode,
		Statements:   statements,
	}, nil
}
func processDeferFunctionDef(node *parser.FunctionDefNode, deferredStatements *deferredStatements) (*parser.FunctionDefNode, error) {
	block, err := processDeferBlock(node.Block, deferredStatements)
	if err != nil {
		return nil, err
	}
	return &parser.FunctionDefNode{
		LocationNode: node.LocationNode,
		ReturnType:   node.ReturnType,
		FunctionName: node.FunctionName,
		Arguments:    node.Arguments,
		Block:        block,
	}, nil
}
