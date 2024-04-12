package parser

import (
	"reflect"
	"testing"
)

func TestParseBlock(t *testing.T) {
	type args struct {
		start int
		src   string
	}
	testCases := []struct {
		name         string
		args         args
		expectedI    int
		expectedNode *BlockNode
		expectedErr  bool
	}{
		{
			name:      "",
			args:      args{0, "{}"},
			expectedI: 2,
			expectedNode: &BlockNode{
				LocationNode: NewLocationNode(0, 2),
				Statements:   []Node{},
			},
			expectedErr: false,
		},
		{
			name:      "",
			args:      args{0, "{{}}"},
			expectedI: 4,
			expectedNode: &BlockNode{
				LocationNode: NewLocationNode(0, 4),
				Statements: []Node{
					&BlockNode{
						LocationNode: NewLocationNode(1, 3),
						Statements:   []Node{},
					},
				},
			},
			expectedErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, got1, err := ParseBlock(tc.args.start, []byte(tc.args.src))
			if (err != nil) != tc.expectedErr {
				t.Errorf("ParseBlock() error = %v, wantErr %v", err, tc.expectedErr)
				return
			}
			if got != tc.expectedI {
				t.Errorf("ParseBlock() i = %v, want %v", got, tc.expectedI)
			}
			if !reflect.DeepEqual(got1, tc.expectedNode) {
				t.Errorf("ParseBlock() node = %v, want %v", got1, tc.expectedNode)
			}
		})
	}
}
