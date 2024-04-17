package parser

type PointerTypeNode struct {
	LocationNode
	Type Node
}

func ParsePointerType(start int, src []byte) (int, *PointerTypeNode, error) {
	i, _, err := ParseExact("*")(start, src)
	if err != nil {
		return 0, nil, ErrWrongParser
	}
	i, typ, err := ParseType(i, src)
	if err != nil {
		return 0, nil, err
	}

	return i, &PointerTypeNode{
		LocationNode: NewLocationNode(start, i),
		Type:         typ,
	}, nil
}
