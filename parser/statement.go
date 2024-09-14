package parser

import "fmt"

type StatementNode struct {
	LocationNode
	Value Node
}

func ParseStatementWithoutSemi(start int, src []byte) (int, *StatementNode, error) {
	i, statement, err := NewOptionParser[Node](
		start, src, fmt.Errorf("unknown statement"),
		Normalize(ParseBlock),
		Normalize(ParseDeclaration),
		Normalize(ParseFunctionCall),
		Normalize(ParseReturn),
		Normalize(ParseDefer),
	)
	if err != nil {
		return 0, nil, err
	}
	return i, &StatementNode{
		LocationNode: NewLocationNode(start, i),
		Value:        statement,
	}, nil
}
func ParseStatement(start int, src []byte) (int, *StatementNode, error) {
	i, statement, err := ParseStatementWithoutSemi(start, src)
	if err != nil {
		return 0, nil, err
	}

	i, _ = ParseWhitespace(i, src)

	i, _, err = ParseExact(";")(i, src)
	if err != nil {
		return 0, nil, err
	}
	return i, statement, nil
}
