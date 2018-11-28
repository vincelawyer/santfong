package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func replaceUlmark(path string) {
	lines, err := FileToLines(path)
	if err != nil {
		panic(err)
	}

	isInUlImage := false
	var lines2 []string
	for _, line := range lines {
		if strings.Contains(line, "ul-mark.png") {
			isInUlImage = true
		}
		if !isInUlImage {
			lines2 = append(lines2, line)
		}
		if strings.Contains(line, "ul-max-width") {
			isInUlImage = false
			if strings.Contains(path, "/en/") {
				lines2 = append(lines2, "UL listed Santfong series products need to be requested, this page is for non UL listed Shenfang series products.")
			} else {
				lines2 = append(lines2, "UL系列為Santfong產品需另外洽詢，本網頁規格圖片為非UL系列Shenfang產品。")
			}
		}
	}
	err = WriteLinesToFile(lines2, path)
	if err != nil {
		panic(err)
	}
}

func WalkAllFilesInDir2(dir string, fn func(string)) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		// check if it is a regular file (not dir)
		if info.Mode().IsRegular() {
			fn(path)
		}
		return nil
	})
}

func TestReplaceUlmark(t *testing.T) {
	dir := "../../content/pages/"

	err := WalkAllFilesInDir2(dir, replaceUlmark)
	if err != nil {
		panic(err)
	}
}
