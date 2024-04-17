package parser

type SetNamer interface {
	SetName(name string)
}

type TypeDefNode struct {
	LocationNode
	Name *IdentifierNode
	Type Node
}

func ParseTypeDef(start int, src []byte) (int, *TypeDefNode, error) {
	i := start

	i, _, err := ParseExact("type")(i, src)
	if err != nil {
		return 0, nil, ErrWrongParser
	}

	i, _ = ParseWhitespace(i, src)

	i, name, err := ParseIdentifier(i, src)
	if err != nil {
		return 0, nil, err
	}

	i, _ = ParseWhitespace(i, src)

	i, typ, err := ParseType(i, src)
	if err != nil {
		return 0, nil, err
	}

	if n, ok := typ.(SetNamer); ok {
		n.SetName(name.Value)
	}

	return i, &TypeDefNode{
		LocationNode: NewLocationNode(start, i),
		Name:         name,
		Type:         typ,
	}, nil
}
