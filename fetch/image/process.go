package main

import (
	"fmt"
	"strings"
)

func getUrl(line string) string {
	ll := strings.Split(line, ":name: ")
	return ll[1]
}

func main() {
	//rstpath := "../../content/pages/en/product/product.rst"
	rstpath := "../../content/pages/zh/product/product.rst"

	var ls []string
	_, err := replaceLines(rstpath, func(s string) string {
		if strings.HasPrefix(s, "  .. image:: ") {
			return ""
		}
		if strings.Contains(s, ":name:") {
			url := getUrl(s)
			fmt.Println(url)
			path := getLocalImagePathFromUrl(url)
			fmt.Println(path)
			download(url, path)

			path = strings.TrimPrefix(path, "../../content")
			ss := "  .. image:: {filename}" + path
			ls = append(ls, ss)
		}
		ls = append(ls, s)
		return s
	})
	if err != nil {
		panic(err)
	}
	err = WriteLinesToFile(ls, "tmp.rst")
	if err != nil {
		panic(err)
	}
}
