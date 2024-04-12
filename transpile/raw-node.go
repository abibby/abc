package transpile

import "github.com/abibby/abc/parser"

type RawNode struct {
	parser.LocationNode
	Raw string
}

func NewRawNode(raw string) *RawNode {
	return &RawNode{
		LocationNode: parser.NewLocationNode(0, 0),
		Raw:          raw,
	}
}

func transpileRawNode(s statements, n *RawNode) error {
	_, err := s.Write([]byte(n.Raw))
	return err
}
