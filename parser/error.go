package parser

import (
	"errors"
	"fmt"
)

var (
	ErrWrongParser = errors.New("wrong parser")
)

type Error struct {
	err  error
	file string
	src  []byte
	loc  int
}

func NewError(src []byte, loc int, err error) *Error {
	return &Error{
		err:  err,
		file: "src.abc",
		src:  src,
		loc:  loc,
	}
}

func (e *Error) Error() string {
	line, column := e.LineColumn()
	return fmt.Sprintf("%s:%d:%d %v", e.file, line, column, e.err)
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) LineColumn() (int, int) {
	line := 1
	column := 0

	for i, c := range e.src {
		if i > e.loc {
			return line, column
		}
		if c == '\n' {
			column = 0
			line++
			continue
		}

		column++
	}

	return line, column
}
func (e *Error) WithFile(f string) *Error {
	e.file = f
	return e
}
