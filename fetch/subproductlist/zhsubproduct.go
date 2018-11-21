package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func createAllChineseSubproductPages() {
	rstpath := "../../content/pages/zh/product/product.rst"
	productDir := path.Dir(rstpath)

	lines, err := FileToLines(rstpath)
	if err != nil {
		panic(err)
	}

	count := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "- `") {
			summary, url := getTitleUrlFromLine(line)
			title, slug := zhSummaryToTitleSlug(summary)
			subdir := getSubdirPath(productDir, slug)
			createDirIfNotExist(subdir)
			content := rstMeta(title, slug, "product, 產品, 產品分類", summary, "zh", "", url)
			err := ioutil.WriteFile(path.Join(subdir, "list.rst"), []byte(content), 0644)
			if err != nil {
				panic(err)
			}
			count++
		}
	}
	fmt.Println(count)
}

func modifyChineseProductPage() {
	rstpath := "../../content/pages/zh/product/product.rst"

	lines, err := FileToLines(rstpath)
	if err != nil {
		panic(err)
	}

	var lines2 []string
	for _, line := range lines {
		if strings.HasPrefix(line, "- `") {
			summary, url := getTitleUrlFromLine(line)
			_, slug := zhSummaryToTitleSlug(summary)
			l2 := strings.Replace(line, url, "{filename}"+slug+"/list.rst", 1)
			lines2 = append(lines2, l2)
		} else {
			lines2 = append(lines2, line)
		}
	}

	WriteLinesToFile(lines2, rstpath)
}

func main() {
	//createAllChineseSubproductPages()
	//modifyChineseProductPage()
}
