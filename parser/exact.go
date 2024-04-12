package parser

import "fmt"

type ExactNode struct {
	LocationNode
	Value string
}

func ParseExact(s string) func(start int, src []byte) (int, *ExactNode, error) {
	return func(start int, src []byte) (int, *ExactNode, error) {
		cur := 0
		for i := start; i < len(src); i++ {
			if cur >= len(s) {
				return i, &ExactNode{
					LocationNode: NewLocationNode(start, i),
					Value:        s,
				}, nil
			}
			if s[cur] != src[i] {
				return 0, nil, NewError(src, start, fmt.Errorf("expected %c found %c", s[cur], src[i]))
			}
			cur++
		}
		return 0, nil, NewError(src, start, fmt.Errorf("expected %c found EOF", s[cur]))
	}
}
