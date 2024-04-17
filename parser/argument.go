package parser

type ArgumentNode struct {
	LocationNode
	Type Node
	Name *IdentifierNode
}

func ParseArgument(start int, src []byte) (int, *ArgumentNode, error) {
	i := start
	i, typ, err := ParseType(i, src)
	if err != nil {
		return 0, nil, err
	}

	i, _ = ParseWhitespace(i, src)

	i, name, err := ParseIdentifier(i, src)
	if err != nil {
		return 0, nil, err
	}

	return i, &ArgumentNode{
		LocationNode: NewLocationNode(start, i),
		Type:         typ,
		Name:         name,
	}, nil
}
