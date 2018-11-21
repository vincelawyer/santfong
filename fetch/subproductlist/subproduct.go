package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func getTitleUrlFromLine(line string) (title, url string) {
	l2 := strings.TrimPrefix(line, "- `")
	l3 := strings.TrimSuffix(l2, ">`_")
	l4 := strings.Split(l3, " <")
	title = strings.TrimSpace(l4[0])
	url = strings.TrimSpace(l4[1])
	return
}

func getSubdirPath(productDir, slug string) string {
	return path.Join(productDir, slug)
}

func WriteLinesToFile(lines []string, filename string) {
	err := ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")+"\n"), 0644)
	if err != nil {
		panic(err)
	}
}

func createAllEnglishSubproductPages() {
	rstpath := "../../content/pages/en/product/product.rst"
	productDir := path.Dir(rstpath)

	lines, err := FileToLines(rstpath)
	if err != nil {
		panic(err)
	}

	count := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "- `") {
			title, url := getTitleUrlFromLine(line)
			slug := titleToSlug(title)
			subdir := getSubdirPath(productDir, slug)
			createDirIfNotExist(subdir)
			content := rstMeta(title, slug, "product, product category", title, "en", "", url)
			err := ioutil.WriteFile(path.Join(subdir, "list.rst"), []byte(content), 0644)
			if err != nil {
				panic(err)
			}
			count++
		}
	}
	fmt.Println(count)
}

func modifyEnglishProductPage() {
	rstpath := "../../content/pages/en/product/product.rst"

	lines, err := FileToLines(rstpath)
	if err != nil {
		panic(err)
	}

	var lines2 []string
	for _, line := range lines {
		if strings.HasPrefix(line, "- `") {
			title, url := getTitleUrlFromLine(line)
			slug := titleToSlug(title)
			l2 := strings.Replace(line, url, "{filename}"+slug+"/list.rst", 1)
			lines2 = append(lines2, l2)
		} else {
			lines2 = append(lines2, line)
		}
	}

	WriteLinesToFile(lines2, rstpath)
}
