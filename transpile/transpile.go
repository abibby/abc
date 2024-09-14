package transpile

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/abibby/abc/parser"
)

func Transpile(file string, src []byte) ([]byte, error) {
	nodes, err := parser.Parse(src)
	if perr, ok := err.(*parser.Error); ok {
		return nil, perr.WithFile(file)
	} else if err != nil {
		return nil, err
	}

	nodes, err = processReturn(nodes)
	if err != nil {
		return nil, err
	}

	nodes, err = processDefer(nodes)
	if err != nil {
		return nil, err
	}

	s := newStatements()

	for _, node := range nodes {
		if _, ok := node.(*parser.WhitespaceNode); ok {
			continue
		}
		_, err = s.WriteString("\n")
		if err != nil {
			return nil, err
		}
		err = transpileNode(s, node)
		if err != nil {
			return nil, err
		}
		_, err = s.WriteString("\n")
		if err != nil {
			return nil, err
		}
	}

	return append([]byte("#include \"runtime.h\"\n"), s.Bytes()...), nil
}

func transpileNode(s cWriter, n parser.Node) error {
	switch n := n.(type) {
	case *parser.FunctionDefNode:
		return transpileFunctionNode(s, n)
	case *parser.WhitespaceNode:
		return transpileWhitespaceNode(s, n)
	case *parser.IdentifierNode:
		return transpileIdentifierNode(s, n)
	case *parser.DeclarationNode:
		return transpileDeclarationNode(s, n)
	case *parser.BlockNode:
		return transpileBlockNode(s, n)
	case *parser.StatementNode:
		return transpileStatementNode(s, n)
	case *parser.NumberNode:
		return transpileNumberNode(s, n)
	case *parser.StringNode:
		return transpileStringNode(s, n)
	case *parser.ArgumentNode:
		return transpileArgumentNode(s, n)
	case *parser.FunctionCallNode:
		return transpileFunctionCallNode(s, n)
	case *parser.TypeDefNode:
		return transpileTypeDefNode(s, n)
	case *parser.StructDefNode:
		return transpileStructDefNode(s, n)
	case *parser.StructInitNode:
		return transpileStructInitNode(s, n)
	case *parser.BasicTypeNode:
		return transpileBasicTypeNode(s, n)
	case *parser.NewNode:
		return transpileNewNode(s, n)
	case *parser.PointerTypeNode:
		return transpilePointerTypeNode(s, n)
	case *parser.VariableNode:
		return transpileVariableNode(s, n)
	case *parser.ReturnNode:
		return transpileReturnNode(s, n)
	case *RawNode:
		return transpileRawNode(s, n)
	case *JoinNode:
		return transpileJoinNode(s, n)

	default:
		return fmt.Errorf("no transpiler for %s", reflect.TypeOf(n))
	}
}

func transpileNodes(s cWriter, nodes ...parser.Node) error {
	for _, n := range nodes {
		err := transpileNode(s, n)
		if err != nil {
			return err
		}
	}
	return nil
}

func transpileWhitespaceNode(s cWriter, _ *parser.WhitespaceNode) error {
	_, err := s.Write([]byte(" "))
	return err
}

func transpileFunctionNode(s cWriter, n *parser.FunctionDefNode) error {
	defS := newStatements()
	transpileNodes(defS,
		n.ReturnType,
		NewRawNode(" "),
		n.FunctionName,
		NewRawNode("("),
		joinNodes(n.Arguments, ", "),
		NewRawNode(");"),
	)

	s.Append(defS.Bytes())
	return transpileNodes(s,
		n.ReturnType,
		NewRawNode(" "),
		n.FunctionName,
		NewRawNode("("),
		joinNodes(n.Arguments, ", "),
		NewRawNode(") "),
		n.Block,
	)
}

func transpileIdentifierNode(s cWriter, n *parser.IdentifierNode) error {
	_, err := s.Write([]byte(n.Value))
	return err
}

func transpileBlockNode(s cWriter, n *parser.BlockNode) error {
	_, err := s.WriteString("{")
	if err != nil {
		return err
	}
	sTabbed := s.Tab()
	for _, n := range n.Statements {
		_, err = sTabbed.WriteString("\n")
		if err != nil {
			return nil
		}
		err := transpileNode(sTabbed, n)
		if err != nil {
			return err
		}
	}

	_, err = s.WriteString("\n}")
	if err != nil {
		return err
	}
	return nil
}

func transpileStatementNode(s cWriter, n *parser.StatementNode) error {
	return transpileNodes(s,
		n.Value,
		NewRawNode(";"),
	)
}

func transpileDeclarationNode(s cWriter, n *parser.DeclarationNode) error {
	return transpileNodes(s,
		n.Type,
		NewRawNode(" "),
		n.Name,
		NewRawNode(" = "),
		n.Value,
	)
}
func transpileNumberNode(s cWriter, n *parser.NumberNode) error {
	_, err := fmt.Fprint(s, n.Value)
	return err
}

func transpileStringNode(s cWriter, n *parser.StringNode) error {
	b, err := json.Marshal(n.Value)
	if err != nil {
		return err
	}
	_, err = s.WriteString("new_string(" + string(b) + ")")
	return err
}

func transpileArgumentNode(s cWriter, n *parser.ArgumentNode) error {
	return transpileNodes(s,
		n.Type,
		NewRawNode(" "),
		n.Name,
	)
}
func transpileFunctionCallNode(s cWriter, n *parser.FunctionCallNode) error {
	return transpileNodes(s,
		n.Name,
		NewRawNode("("),
		joinNodes(n.Arguments, ", "),
		NewRawNode(")"),
	)
}

func transpileReturnNode(s cWriter, n *parser.ReturnNode) error {
	_, err := s.WriteString("return")
	if err != nil {
		return err
	}
	if n.Value == nil {
		return nil
	}
	return transpileNode(s, n.Value)
}
