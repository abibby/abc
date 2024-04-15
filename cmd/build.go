/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/abibby/abc/transpile"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		buildDir := ".abc/"
		err := os.RemoveAll(buildDir)
		if err != nil {
			return err
		}

		err = os.MkdirAll(buildDir, 0o777)
		if err != nil {
			return err
		}

		srcDir := "."
		if len(args) > 0 {
			srcDir = args[0]
		}

		err = transpile.Dir(srcDir, buildDir)
		if err != nil {
			return err
		}

		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		cFiles := []string{}
		files, err := os.ReadDir(buildDir)
		if err != nil {
			return err
		}
		for _, f := range files {
			filePath := f.Name()
			if f.IsDir() || !strings.HasSuffix(filePath, ".c") {
				continue
			}
			cFiles = append(cFiles, filePath)
		}

		c := exec.Command("gcc", append(cFiles, "-Wall", "-Wextra", "-std=c89", "-pedantic", "-Wmissing-prototypes", "-Wstrict-prototypes", "-Wold-style-definition", "-o", path.Join(wd, "a.out"))...)
		c.Dir = buildDir
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		return c.Run()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
