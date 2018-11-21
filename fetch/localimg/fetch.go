package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"unicode"

	"github.com/mozillazg/go-pinyin"
)

var a = pinyin.NewArgs()

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func getImageDir(rstpath string) string {
	return path.Join(path.Dir(rstpath), "image")
}

func getImageFilename(url string) (s string) {
	filename := path.Base(url)
	for _, r := range filename {
		if unicode.Is(unicode.Han, r) {
			s += string(pinyin.Pinyin(string(r), a)[0][0])
		} else {
			s += string(r)
		}
	}
	return
}

func main() {
	rstpath := "../../content/pages/zh/product/product.rst"
	imgDir := getImageDir(rstpath)
	createDirIfNotExist(imgDir)

	lines, err := FileToLines(rstpath)
	if err != nil {
		panic(err)
	}

	var lines2 []string
	for _, line := range lines {
		if strings.HasPrefix(line, "  .. image::") {
			url := strings.TrimPrefix(line, "  .. image:: ")
			filename := getImageFilename(url)
			filepath := path.Join(imgDir, filename)
			download(url, filepath)
			lines2 = append(lines2, "  .. image:: {filename}image/"+filename)
			lines2 = append(lines2, "     :name: "+url)
		} else {
			lines2 = append(lines2, line)
		}
	}

	fc := strings.Join(lines2, "\n") + "\n"
	fmt.Print(fc)
}
