package parser

type IdentifierNode struct {
	LocationNode
	Value string
}

func ParseIdentifier(start int, src []byte) (int, *IdentifierNode, error) {
	end := start
	for i := start; i < len(src); i++ {
		c := src[i]
		end = i
		if i == start && '0' <= c && c <= '9' {
			return 0, nil, ErrWrongParser
		}
		if ('a' > c || c > 'z') && ('A' > c || c > 'Z') && ('0' > c || c > '9') && c != '_' {
			break
		}
	}
	if start == end {
		return 0, nil, ErrWrongParser
	}

	l := NewLocationNode(start, end-1)
	return end, &IdentifierNode{
		LocationNode: l,
		Value:        l.String(src),
	}, nil
}
