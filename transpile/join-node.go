package transpile

import "github.com/abibby/abc/parser"

type JoinNode struct {
	parser.LocationNode
	Nodes []parser.Node
	Glue  string
}

func joinNodes[T parser.Node](nodes []T, glue string) *JoinNode {
	genericNodes := make([]parser.Node, len(nodes))
	for i, n := range nodes {
		genericNodes[i] = n
	}
	return &JoinNode{
		LocationNode: parser.NewLocationNode(0, 0),
		Nodes:        genericNodes,
		Glue:         glue,
	}
}

func transpileJoinNode(s statements, n *JoinNode) error {
	for i, node := range n.Nodes {
		if i > 0 {
			_, err := s.WriteString(n.Glue)
			if err != nil {
				return err
			}
		}
		err := transpileNode(s, node)
		if err != nil {
			return err
		}
	}
	return nil
}
