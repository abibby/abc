package parser

type DeferNode struct {
	LocationNode
	Statement *StatementNode
}

func ParseDefer(start int, src []byte) (int, *DeferNode, error) {
	i := start

	i, _, err := ParseExact("defer")(i, src)
	if err != nil {
		return 0, nil, ErrWrongParser
	}

	i, _ = ParseWhitespace(i, src)

	i, statement, err := ParseStatementWithoutSemi(i, src)
	if err != nil {
		return i, nil, err
	}

	return i, &DeferNode{
		LocationNode: NewLocationNode(start, i),
		Statement:    statement,
	}, nil
}
