package parser

import "fmt"

func ParseValue(start int, src []byte) (int, Node, error) {
	return NewOptionParser(
		start, src, fmt.Errorf("unknown value"),
		Normalize(ParseNumber),
		Normalize(ParseString),
	)
}
