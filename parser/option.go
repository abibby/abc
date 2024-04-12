package parser

import "errors"

type Parser func(start int, src []byte) (int, Node, error)

func Normalize[T Node](cb func(start int, src []byte) (int, T, error)) Parser {
	return func(start int, src []byte) (int, Node, error) {
		return cb(start, src)
	}
}

func NewOptionParser(start int, src []byte, fallbackErr error, parsers ...Parser) (int, Node, error) {
	var i int
	var node Node
	var err error

	for _, parser := range parsers {
		i, node, err = parser(start, src)
		if errors.Is(err, ErrWrongParser) {
			continue
		} else if err != nil {
			return 0, nil, err
		}
		return i, node, nil
	}

	return 0, nil, NewError(src, start, fallbackErr)
}
