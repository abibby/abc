package parser

import "fmt"

func ParseType(start int, src []byte) (int, Node, error) {
	return NewOptionParser(
		start, src, fmt.Errorf("unknown type"),
		Normalize(ParseStructDef),
		Normalize(ParseBasicType),
	)
}
