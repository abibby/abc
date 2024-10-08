package parser

import "fmt"

func ParseTopLevel(start int, src []byte) (int, Node, error) {
	return NewOptionParser[Node](
		start, src, fmt.Errorf("unknown top level construct"),
		Normalize(ParseFunctionDef),
		Normalize(ParseTypeDef),
	)
}
