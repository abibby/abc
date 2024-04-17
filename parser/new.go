package parser

type NewNode struct {
	LocationNode
	Value ValueNode
	Type  string
}

func ParsePointer(start int, src []byte) (int, *NewNode, error) {
	i := start

	i, _, err := ParseExact("&")(i, src)
	if err != nil {
		return 0, nil, ErrWrongParser
	}

	i, _ = ParseWhitespace(i, src)

	i, value, err := ParseValue(i, src)
	if err != nil {
		return 0, nil, err
	}

	return i, &NewNode{
		LocationNode: NewLocationNode(start, i),
		Value:        value,
	}, nil
}

func (n *NewNode) GetType() string {
	return n.Value.GetType() + "*"
}
