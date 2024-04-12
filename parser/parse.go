package parser

type Node interface {
	Start() int
	End() int
	Len() int
	String(src []byte) string
}

type LocationNode struct {
	start int
	end   int
}

func (n *LocationNode) Start() int {
	return n.start
}
func (n *LocationNode) End() int {
	return n.end
}
func (n *LocationNode) Len() int {
	return n.end - n.start
}
func (n *LocationNode) String(src []byte) string {
	return string(src[n.start : n.end+1])
}

func NewLocationNode(start, end int) LocationNode {
	return LocationNode{start: start, end: end}
}

type NodeParser func(start int, src []byte) (int, Node, error)

func Parse(src []byte) ([]Node, error) {
	nodes := []Node{}
	var node Node
	var err error

	i := 0
	for i < len(src) {
		i, node, err = ParseFunction(i, src)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)

		i, node = ParseWhitespace(i, src)
		nodes = append(nodes, node)
	}
	return nodes, nil
}
