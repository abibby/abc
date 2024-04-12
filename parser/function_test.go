package parser_test

import (
	"testing"

	"github.com/abibby/abc/parser"
	"github.com/stretchr/testify/assert"
)

func TestParseFunction(t *testing.T) {
	testCases := []struct {
		Name         string
		Src          string
		ExpectedI    int
		ExpectedNode parser.Node
	}{
		{
			Name:      "func",
			Src:       "int main() {}",
			ExpectedI: 13,
			ExpectedNode: &parser.FunctionNode{
				LocationNode: parser.NewLocationNode(0, 13),
				ReturnType:   &parser.IdentifierNode{LocationNode: parser.NewLocationNode(0, 3), Value: "int"},
				FunctionName: &parser.IdentifierNode{LocationNode: parser.NewLocationNode(4, 8), Value: "main"},
				Arguments:    []*parser.ArgumentNode{},
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {
			i, node, err := parser.ParseFunction(0, []byte(tc.Src))
			if assert.NoError(t, err) {
				assert.Equal(t, tc.ExpectedI, i)
				assert.Equal(t, tc.ExpectedNode, node)
			}
		})
	}
}
