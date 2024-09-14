package transpile

import "github.com/abibby/abc/parser"

type scope struct{}

func newScope() *scope {
	return &scope{}
}

func (s *scope) Child() *scope {
	return s
}

func ProcessLifetimes(nodes []parser.Node) ([]parser.Node, error) {
	return processLifetimes(nodes, newScope())
}

func processLifetimes(nodes []parser.Node, scope *scope) ([]parser.Node, error) {
	newNodes := make([]parser.Node, 0, len(nodes))

	for _, n := range nodes {
		newNode, err := processLifetimesNode(n, scope)
		if err != nil {
			return nil, err
		}
		newNodes = append(newNodes, newNode)
	}
	return newNodes, nil
}
func processLifetimesNode(node parser.Node, scope *scope) (parser.Node, error) {
	switch n := node.(type) {
	case *parser.BlockNode:
		return processLifetimesBlock(n, scope)
	case *parser.FunctionDefNode:
		return processLifetimesFunctionDef(n, scope)
	default:
		return n, nil
	}
}
func processLifetimesBlock(node *parser.BlockNode, scope *scope) (*parser.BlockNode, error) {
	scope = scope.Child()
	statements := []parser.Node{}

	// for _, s := range node.Statements {

	// }

	return &parser.BlockNode{
		LocationNode: node.LocationNode,
		Statements:   statements,
	}, nil
}
func processLifetimesFunctionDef(node *parser.FunctionDefNode, scope *scope) (*parser.FunctionDefNode, error) {
	block, err := processLifetimesBlock(node.Block, scope)
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
