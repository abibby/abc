package parser

import (
	"fmt"
)

type FunctionDefNode struct {
	LocationNode
	ReturnType   Node
	FunctionName *IdentifierNode
	Arguments    []*ArgumentNode
	Block        *BlockNode
}

func ParseFunctionDef(start int, src []byte) (int, *FunctionDefNode, error) {
	i := start

	i, _, err := ParseExact("func")(i, src)
	if err != nil {
		return 0, nil, ErrWrongParser
	}

	i, _ = ParseWhitespace(i, src)

	i, functionName, err := ParseIdentifier(i, src)
	if err != nil {
		return 0, nil, err
	}

	if src[i] != '(' {
		return 0, nil, NewError(src, i, fmt.Errorf("expected ( received %c", src[i]))
	}
	i++

	i, _ = ParseWhitespace(i, src)

	var arg *ArgumentNode
	args := []*ArgumentNode{}
	for src[i] != ')' {
		i, arg, err = ParseArgument(i, src)
		if err != nil {
			return 0, nil, err
		}
		args = append(args, arg)
	}

	i++

	i, _ = ParseWhitespace(i, src)

	i, returnType, err := ParseType(i, src)
	if err != nil {
		return 0, nil, err
	}

	i, _ = ParseWhitespace(i, src)

	i, block, err := ParseBlock(i, src)
	if err != nil {
		return 0, nil, err
	}

	return i, &FunctionDefNode{
		LocationNode: NewLocationNode(start, i),
		ReturnType:   returnType,
		FunctionName: functionName,
		Arguments:    args,
		Block:        block,
	}, nil
}
