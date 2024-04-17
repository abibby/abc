package parser

type DeclarationNode struct {
	LocationNode
	Name  *IdentifierNode
	Type  Node
	Value Node
}

func ParseDeclaration(start int, src []byte) (int, *DeclarationNode, error) {
	i := start

	i, _, err := ParseExact("var")(i, src)
	if err != nil {
		return 0, nil, ErrWrongParser
	}

	i, _ = ParseWhitespace(i, src)

	i, name, err := ParseIdentifier(i, src)
	if err != nil {
		return i, nil, err
	}

	i, _ = ParseWhitespace(i, src)

	i, typ, err := ParseType(i, src)
	if err != nil {
		return 0, nil, err
	}

	i, _ = ParseWhitespace(i, src)

	i, _, err = ParseExact("=")(i, src)
	if err != nil {
		return 0, nil, err
	}

	i, _ = ParseWhitespace(i, src)

	i, value, err := ParseValue(i, src)
	if err != nil {
		return 0, nil, err
	}

	return i, &DeclarationNode{
		LocationNode: NewLocationNode(start, i),
		Name:         name,
		Type:         typ,
		Value:        value,
	}, nil
}
