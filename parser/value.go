package parser

import "fmt"

type ValueNode interface {
	Node
	GetType() string
}

func ParseValue(start int, src []byte) (int, ValueNode, error) {
	return NewOptionParser[ValueNode](
		start, src, fmt.Errorf("unknown value"),
		Normalize(ParseNumber),
		Normalize(ParseString),
		Normalize(ParseStructInit),
		Normalize(ParseFunctionCall),
		Normalize(ParsePointer),
		Normalize(ParseVariable),
	)
}
