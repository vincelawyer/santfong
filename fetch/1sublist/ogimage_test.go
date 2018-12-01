package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func replaceOgImage(fpath string) {
	oldprefix := "https://vincelawyer.github.io/santfong"
	newprefix := "https://sunteron.com"

	_, err := replaceLinesAndWrite(fpath, func(s string) string {
		if strings.HasPrefix(s, ":og_image: "+oldprefix) {
			tmp := strings.TrimPrefix(s, ":og_image: "+oldprefix)
			after := ":og_image: " + newprefix + tmp
			fmt.Println(after)
			return after
		}
		return s
	})
	if err != nil {
		panic(err)
	}
}

func WalkAllFilesInDir(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		// check if it is a regular file (not dir)
		if info.Mode().IsRegular() {
			replaceOgImage(path)
		}
		return nil
	})
}

func TestSetOgImage(t *testing.T) {
	dir := "../../content/pages/"

	err := WalkAllFilesInDir(dir)
	if err != nil {
		panic(err)
	}
}
