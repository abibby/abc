package parser

type BlockNode struct {
	LocationNode
	Statements []Node
}

func ParseBlock(start int, src []byte) (int, *BlockNode, error) {
	i := start

	if src[i] != '{' {
		return 0, nil, ErrWrongParser
	}
	i++

	i, _ = ParseWhitespace(i, src)

	statements := []Node{}
	var err error
	for src[i] != '}' {
		var statement Node
		i, statement, err = ParseStatement(i, src)
		if err != nil {
			return 0, nil, err
		}
		statements = append(statements, statement)

		i, _ = ParseWhitespace(i, src)
	}

	// Skip over }
	i++

	return i, &BlockNode{
		LocationNode: NewLocationNode(start, i),
		Statements:   statements,
	}, nil
}
