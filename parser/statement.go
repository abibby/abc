package parser

import "fmt"

func ParseStatement(start int, src []byte) (int, Node, error) {
	i, statement, err := NewOptionParser(
		start, src, fmt.Errorf("unknown statement"),
		Normalize(ParseBlock),
		Normalize(ParseDeclaration),
		Normalize(ParseFunctionCall),
	)
	if err != nil {
		return 0, nil, err
	}

	i, _ = ParseWhitespace(i, src)

	i, _, err = ParseExact(";")(i, src)
	if err != nil {
		return 0, nil, err
	}
	return i, statement, nil
}
