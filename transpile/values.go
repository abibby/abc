package transpile

import "github.com/abibby/abc/parser"

func transpileVariableNode(s statements, n *parser.VariableNode) error {
	s.WriteString(n.Name.Value)
	return nil
}
