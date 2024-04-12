package parser

import (
	"fmt"
)

type FunctionNode struct {
	LocationNode
	ReturnType   *IdentifierNode
	FunctionName *IdentifierNode
	Arguments    []*ArgumentNode
	Block        *BlockNode
}

func ParseFunction(start int, src []byte) (int, *FunctionNode, error) {
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

	i, returnType, err := ParseIdentifier(i, src)
	if err != nil {
		return 0, nil, err
	}

	i, _ = ParseWhitespace(i, src)

	i, block, err := ParseBlock(i, src)
	if err != nil {
		return 0, nil, err
	}

	return i, &FunctionNode{
		LocationNode: NewLocationNode(start, i),
		ReturnType:   returnType,
		FunctionName: functionName,
		Arguments:    args,
		Block:        block,
	}, nil
}

type ArgumentNode struct {
	LocationNode
	Type *IdentifierNode
	Name *IdentifierNode
}

func ParseArgument(start int, src []byte) (int, *ArgumentNode, error) {
	i := start
	i, argType, err := ParseIdentifier(i, src)
	if err != nil {
		return 0, nil, err
	}

	i, _ = ParseWhitespace(i, src)

	i, argName, err := ParseIdentifier(i, src)
	if err != nil {
		return 0, nil, err
	}

	return i, &ArgumentNode{
		LocationNode: NewLocationNode(start, i),
		Type:         argType,
		Name:         argName,
	}, nil
}
