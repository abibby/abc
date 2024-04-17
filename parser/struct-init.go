package parser

type StructInitProp struct {
	Key   string
	Value Node
}

type StructInitNode struct {
	LocationNode
	Type  *IdentifierNode
	Props []*StructInitProp
}

func ParseStructInit(start int, src []byte) (int, *StructInitNode, error) {
	i := start

	i, typ, err := ParseIdentifier(i, src)
	if err != nil {
		return 0, nil, ErrWrongParser
	}

	i, _ = ParseWhitespace(i, src)

	i, _, err = ParseExact("{")(i, src)
	if err != nil {
		return 0, nil, ErrWrongParser
	}

	i, _ = ParseWhitespace(i, src)

	props := []*StructInitProp{}
	var name *IdentifierNode
	var value Node
	for src[i] != '}' {
		i, name, err = ParseIdentifier(i, src)
		if err != nil {
			return 0, nil, err
		}

		i, _ = ParseWhitespace(i, src)

		i, _, err = ParseExact(":")(i, src)
		if err != nil {
			return 0, nil, err
		}

		i, _ = ParseWhitespace(i, src)

		i, value, err = ParseValue(i, src)
		if err != nil {
			return 0, nil, err
		}

		i, _ = ParseWhitespace(i, src)

		i, _, err = ParseExact(",")(i, src)
		if err != nil {
			return 0, nil, err
		}

		i, _ = ParseWhitespace(i, src)

		props = append(props, &StructInitProp{
			Key:   name.Value,
			Value: value,
		})
	}
	i++

	return i, &StructInitNode{
		LocationNode: NewLocationNode(start, i),
		Type:         typ,
		Props:        props,
	}, nil
}

func (n *StructInitNode) GetType() string {
	return n.Type.Value
}
