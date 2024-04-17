package parser

type NumberNode struct {
	LocationNode
	Value string
}

func ParseNumber(start int, src []byte) (int, *NumberNode, error) {
	end := start
	for i := start; i < len(src); i++ {
		c := src[i]
		end = i
		if ('0' > c || c > '9') && c != '_' && c != '.' {
			break
		}
	}
	if start == end {
		return 0, nil, ErrWrongParser
	}

	l := NewLocationNode(start, end-1)
	return end, &NumberNode{
		LocationNode: l,
		Value:        l.String(src),
	}, nil
}

func (n *NumberNode) GetType() string {
	return "int"
}
