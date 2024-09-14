package transpile

import (
	"bytes"
	"io"
)

type cWriter interface {
	io.Writer
	io.StringWriter
	Tab() cWriter
	Append(statement []byte)
	Bytes() []byte
}
type rootStatements struct {
	current    *bytes.Buffer
	statements [][]byte
}

func newStatements() cWriter {
	return &rootStatements{
		current:    &bytes.Buffer{},
		statements: [][]byte{},
	}
}

func (s *rootStatements) Tab() cWriter {
	return &subStatements{
		root: s,
		tab:  1,
	}
}
func (s *rootStatements) Append(statement []byte) {
	s.statements = append(s.statements, statement)
}
func (s *rootStatements) Write(statement []byte) (int, error) {
	return s.current.Write(statement)
}
func (s *rootStatements) WriteString(str string) (int, error) {
	return s.current.WriteString(str)
}
func (s *rootStatements) Bytes() []byte {
	b := &bytes.Buffer{}

	for _, statement := range s.statements {
		b.WriteString("\n")
		_, err := b.Write(statement)
		if err != nil {
			panic(err)
		}
		b.WriteString("\n")
	}

	_, err := b.Write(s.current.Bytes())
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}

type subStatements struct {
	root *rootStatements
	tab  int
}

func (s *subStatements) Tab() cWriter {
	return &subStatements{
		root: s.root,
		tab:  s.tab + 1,
	}
}
func (s *subStatements) Append(statement []byte) {
	s.root.Append(statement)
}
func (s *subStatements) Write(b []byte) (int, error) {
	b = bytes.ReplaceAll(b, []byte("\n"), []byte("\n    "))
	return s.root.Write(b)
}
func (s *subStatements) WriteString(str string) (int, error) {
	return s.Write([]byte(str))
}
func (s *subStatements) Bytes() []byte {
	return s.root.Bytes()
}
