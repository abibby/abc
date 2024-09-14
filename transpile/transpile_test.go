package transpile_test

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/abibby/abc/transpile"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name   string
	ABCSrc []byte
	CSrc   []byte
}

//go:embed test-cases/*
var files embed.FS

func TestExamples(t *testing.T) {
	filter := os.Getenv("TEST_FILTER")

	testCases := []*TestCase{}
	err := fs.WalkDir(files, "test-cases", func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if d.IsDir() {
			return nil
		}

		if filter != "" {
			if !strings.Contains(p, filter) {
				return nil
			}
		}

		ext := path.Ext(p)
		name := strings.TrimSuffix(p, ext)
		b, err := files.ReadFile(p)
		if err != nil {
			return err
		}

		tc, err := parseMD(b, name)
		if err != nil {
			return err
		}

		testCases = append(testCases, tc)

		return nil
	})
	if !assert.NoError(t, err) {
		return
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			cSrc, err := transpile.Transpile(tc.Name+".abc", tc.ABCSrc)
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(t, string(tc.CSrc), string(cSrc))
		})
	}
}

func parseMD(b []byte, name string) (*TestCase, error) {
	tc := &TestCase{Name: name}
	parts := bytes.Split(b, []byte("```"))
	for _, part := range parts {
		part = bytes.TrimLeft(part, " ")
		if bytes.HasPrefix(part, []byte("abc\n")) {
			tc.ABCSrc = bytes.TrimPrefix(part, []byte("abc\n"))
		}
		if bytes.HasPrefix(part, []byte("c\n")) {
			tc.CSrc = bytes.TrimPrefix(part, []byte("c\n"))
		}
	}
	if tc.ABCSrc == nil {
		return nil, fmt.Errorf("no abc src in test")
	}
	if tc.CSrc == nil {
		return nil, fmt.Errorf("no c src in test")
	}
	return tc, nil
}
