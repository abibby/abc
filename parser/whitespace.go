package parser

type WhitespaceNode struct {
	LocationNode
}

func ParseWhitespace(start int, src []byte) (int, *WhitespaceNode) {
	for i := start; i < len(src); i++ {
		switch src[i] {
		case ' ', '\n', '\t':
			// do nothing
		default:
			return i, &WhitespaceNode{LocationNode: NewLocationNode(start, i-1)}
		}
	}

	return len(src), &WhitespaceNode{LocationNode: NewLocationNode(start, len(src)-1)}
}
