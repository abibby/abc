package parser

type ReturnNode struct {
	LocationNode
	Value ValueNode
}

func ParseReturn(start int, src []byte) (int, *ReturnNode, error) {
	i, _, err := ParseExact("return")(start, src)
	if err != nil {
		return 0, nil, ErrWrongParser
	}

	i, _ = ParseWhitespace(i, src)

	i, typ, err := ParseValue(i, src)
	if err != nil {
		return 0, nil, err
	}

	return i, &ReturnNode{
		LocationNode: NewLocationNode(start, i),
		Value:        typ,
	}, nil
}
