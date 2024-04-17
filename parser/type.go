package parser

import "fmt"

type TypeNode interface {
	Node
}

func ParseType(start int, src []byte) (int, Node, error) {
	return NewOptionParser[TypeNode](
		start, src, fmt.Errorf("unknown type"),
		Normalize(ParseStructDef),
		Normalize(ParseBasicType),
		Normalize(ParsePointerType),

		// TODO: replace with something type specific
		Normalize(ParseIdentifier),
	)
}
