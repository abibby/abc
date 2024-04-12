package main_test

import (
	"embed"
	"io/fs"
	"path"
	"strings"
	"testing"

	"github.com/abibby/abc/transpile"
	"github.com/stretchr/testify/assert"
)

//go:embed test-cases/*
var files embed.FS

func TestExamples(t *testing.T) {
	type TestCase struct {
		abcSrc []byte
		cSrc   []byte
	}
	testCases := map[string]*TestCase{}
	err := fs.WalkDir(files, "test-cases", func(p string, d fs.DirEntry, walkErr error) error {
		if d.IsDir() {
			return nil
		}

		ext := path.Ext(p)
		name := strings.TrimSuffix(p, ext)
		tc, ok := testCases[name]
		if !ok {
			tc = &TestCase{}
			testCases[name] = tc
		}
		b, err := files.ReadFile(p)
		if err != nil {
			panic(err)
		}
		if ext == ".abc" {
			tc.abcSrc = b
		} else if ext == ".c" {
			tc.cSrc = b
		} else {
			t.Errorf("invalid test file %s", p)
			return nil
		}

		return nil
	})
	if !assert.NoError(t, err) {
		return
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			cSrc, err := transpile.Transpile(name+".abc", tc.abcSrc)
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(t, string(tc.cSrc), string(cSrc))
		})
	}
}
