package transpile

import "github.com/abibby/abc/parser"

func transpileVariableNode(s cWriter, n *parser.VariableNode) error {
	s.WriteString(n.Name.Value)
	return nil
}
