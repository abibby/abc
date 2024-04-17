package parser

import (
	"bytes"
	"fmt"
)

type StringNode struct {
	LocationNode
	Value string
}

func ParseString(start int, src []byte) (int, *StringNode, error) {
	i := start
	i, _, err := ParseExact(`"`)(i, src)
	if err != nil {
		return 0, nil, ErrWrongParser
	}
	str := &bytes.Buffer{}

	for i < len(src) {
		c := src[i]
		i++
		switch c {
		case '"':
			return i, &StringNode{
				LocationNode: NewLocationNode(start, i-1),
				Value:        str.String(),
			}, nil
		case '\\':
			i++
			switch src[i] {
			case 'n':
				str.WriteByte('\n')
				i++
			case 't':
				str.WriteByte('\t')
				i++
			case 'r':
				str.WriteByte('\r')
				i++
			}
		default:
			str.WriteByte(c)
		}
	}

	return 0, nil, NewError(src, i, fmt.Errorf("expected \" found EOF"))
}

func (n *StringNode) GetType() string {
	return "string"
}
