package parser

type StructDefNode struct {
	LocationNode
	Props []*ArgumentNode
	Name  string
}

func ParseStructDef(start int, src []byte) (int, *StructDefNode, error) {
	i := start

	i, _, err := ParseExact("struct")(i, src)
	if err != nil {
		return 0, nil, ErrWrongParser
	}

	i, _ = ParseWhitespace(i, src)

	i, _, err = ParseExact("{")(i, src)
	if err != nil {
		return 0, nil, err
	}

	i, _ = ParseWhitespace(i, src)

	props := []*ArgumentNode{}
	var prop *ArgumentNode
	for src[i] != '}' {
		i, prop, err = ParseArgument(i, src)
		if err != nil {
			return 0, nil, err
		}

		i, _ = ParseWhitespace(i, src)

		i, _, err = ParseExact(";")(i, src)
		if err != nil {
			return 0, nil, err
		}

		i, _ = ParseWhitespace(i, src)

		props = append(props, prop)
	}
	i++

	return i, &StructDefNode{
		LocationNode: NewLocationNode(start, i),
		Props:        props,
	}, nil
}

func (n *StructDefNode) SetName(name string) {
	n.Name = name
}
