package transpile

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/abibby/abc/runtime"
)

func Dir(srcDir, buildDir string) error {
	abcExt := ".abc"

	cFiles := []string{}

	err := filepath.Walk(srcDir, func(srcPath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(srcPath, abcExt) {
			return nil
		}

		buildPath := path.Join(buildDir, strings.TrimPrefix(srcPath, srcDir))
		buildPath = buildPath[:len(buildPath)-len(abcExt)] + ".c"

		b, err := os.ReadFile(srcPath)
		if err != nil {
			return err
		}

		cSrc, err := Transpile(srcPath, b)
		if err != nil {
			return err
		}

		err = os.WriteFile(buildPath, cSrc, 0o644)
		if err != nil {
			return err
		}

		if strings.HasSuffix(buildPath, ".c") {
			cFiles = append(cFiles, strings.TrimPrefix(buildPath, buildDir))
		}
		return nil
	})
	if err != nil {
		return err
	}

	runtimeDir := "src"
	err = fs.WalkDir(runtime.Files, runtimeDir, func(runtimePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		buildPath := path.Join(buildDir, strings.TrimPrefix(runtimePath, runtimeDir))
		b, err := runtime.Files.ReadFile(runtimePath)
		if err != nil {
			return err
		}
		err = os.WriteFile(buildPath, b, 0o644)
		if err != nil {
			return err
		}

		if strings.HasSuffix(buildPath, ".c") {
			cFiles = append(cFiles, strings.TrimPrefix(buildPath, buildDir))
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
