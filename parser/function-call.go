package parser

type FunctionCallNode struct {
	LocationNode
	Name      *IdentifierNode
	Arguments []*IdentifierNode
}

func ParseFunctionCall(start int, src []byte) (int, *FunctionCallNode, error) {
	i := start
	i, name, err := ParseIdentifier(i, src)
	if err != nil {
		return 0, nil, err
	}

	i, _, err = ParseExact("(")(i, src)
	if err != nil {
		return 0, nil, ErrWrongParser
	}

	args := []*IdentifierNode{}

	var arg *IdentifierNode
	for src[i] != ')' {
		i, arg, err = ParseIdentifier(i, src)
		if err != nil {
			return 0, nil, err
		}

		i, _ = ParseWhitespace(i, src)

		args = append(args, arg)

		if src[i] != ',' {
			break
		}
		i++

		i, _ = ParseWhitespace(i, src)

	}

	i, _ = ParseWhitespace(i, src)

	i, _, err = ParseExact(")")(i, src)
	if err != nil {
		return 0, nil, err
	}

	return i, &FunctionCallNode{
		LocationNode: NewLocationNode(start, i),
		Name:         name,
		Arguments:    args,
	}, nil
}

func (n *FunctionCallNode) GetType() string {
	return "unknown"
}
