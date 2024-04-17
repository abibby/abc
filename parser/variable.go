package parser

type VariableNode struct {
	LocationNode
	Name *IdentifierNode
}

func ParseVariable(start int, src []byte) (int, *VariableNode, error) {
	i := start
	i, name, err := ParseIdentifier(i, src)
	if err != nil {
		return 0, nil, err
	}

	return i, &VariableNode{
		LocationNode: NewLocationNode(start, i),
		Name:         name,
	}, nil
}

func (n *VariableNode) GetType() string {
	return "unknown"
}
