package main

import (
	"fmt"
	"strings"
)

type productItem struct {
	Title     string
	Href      string
	ImageSrcs []string
}

type productList struct {
	RstPath string
	Title   string
	Url     string
	Items   []productItem
}

func (l *productList) SetUrlTitle() {
	_, err := replaceLines(l.RstPath, func(s string) string {
		if strings.HasPrefix(s, ":title: ") {
			l.Title = strings.TrimPrefix(s, ":title: ")
		}
		if strings.HasPrefix(s, ":source: ") {
			l.Url = strings.TrimPrefix(s, ":source: ")
		}
		return s
	})
	if err != nil {
		panic(err)
	}
}

func (l productList) CreateFinalProductRstFiles() {
	fmt.Println()
	for _, item := range l.Items {
		ogImg := getRstImagePath(item.ImageSrcs[0])

		s := rstMeta(item.Title, titleToSlug(item.Title),
			"product, "+l.Title,
			item.Title+" - "+l.Title,
			"en", "", item.Href, ogImg)
		targetPath := getFinalProductRstPath(l.RstPath, item.Title)
		fmt.Println(s)
		writeToFile(targetPath, s)
	}
}

func newProductList(rstpath string) productList {
	list := productList{
		RstPath: rstpath,
	}
	list.SetUrlTitle()

	return list
}
