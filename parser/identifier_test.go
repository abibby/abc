package parser

import (
	"reflect"
	"testing"
)

func TestParseIdentifier(t *testing.T) {
	type args struct {
		start int
		src   string
	}
	tests := []struct {
		name     string
		args     args
		wantI    int
		wantNode *IdentifierNode
		wantErr  bool
	}{
		{
			name:    "error",
			args:    args{0, "="},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParseIdentifier(tt.args.start, []byte(tt.args.src))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIdentifier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantI {
				t.Errorf("ParseIdentifier() got = %v, want %v", got, tt.wantI)
			}
			if !reflect.DeepEqual(got1, tt.wantNode) {
				t.Errorf("ParseIdentifier() got1 = %v, want %v", got1, tt.wantNode)
			}
		})
	}
}
