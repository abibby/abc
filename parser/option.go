package parser

import (
	"errors"
	"fmt"
	"reflect"
)

type Parser func(start int, src []byte) (int, Node, error)

func Normalize[T Node](cb func(start int, src []byte) (int, T, error)) Parser {
	return func(start int, src []byte) (int, Node, error) {
		return cb(start, src)
	}
}

func NewOptionParser[T Node](start int, src []byte, fallbackErr error, parsers ...Parser) (int, T, error) {
	var i int
	var node Node
	var err error

	for _, parser := range parsers {
		i, node, err = parser(start, src)
		if errors.Is(err, ErrWrongParser) {
			continue
		} else if err != nil {
			var zero T
			return 0, zero, err
		}
		n, ok := node.(T)
		if !ok {
			var zero T
			return 0, zero, NewError(src, start, fmt.Errorf("node %s must be of type %s", reflect.TypeOf(node), reflect.TypeOf(([]T)(nil)).Elem()))
		}
		return i, n, nil
	}

	var zero T
	return 0, zero, NewError(src, start, fallbackErr)
}
