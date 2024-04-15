package parser

type ArgumentNode struct {
	LocationNode
	Type Node
	Name *IdentifierNode
}

func ParseArgument(start int, src []byte) (int, *ArgumentNode, error) {
	i := start
	i, argType, err := ParseType(i, src)
	if err != nil {
		return 0, nil, err
	}

	i, _ = ParseWhitespace(i, src)

	i, argName, err := ParseIdentifier(i, src)
	if err != nil {
		return 0, nil, err
	}

	return i, &ArgumentNode{
		LocationNode: NewLocationNode(start, i),
		Type:         argType,
		Name:         argName,
	}, nil
}
