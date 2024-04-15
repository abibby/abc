package parser

type BasicTypeNode struct {
	LocationNode
	Value string
}

var BasicTypes = []string{
	"string",
	"int",
}

func ParseBasicType(start int, src []byte) (int, *BasicTypeNode, error) {
	for _, typ := range BasicTypes {
		i, _, err := ParseExact(typ)(start, src)
		if err == nil {
			return i, &BasicTypeNode{
				LocationNode: NewLocationNode(start, i),
				Value:        typ,
			}, nil
		}
	}

	return 0, nil, ErrWrongParser
}
