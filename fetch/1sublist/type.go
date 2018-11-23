package main

import (
	"fmt"
	"strings"
)

type productItem struct {
	Title     string
	Href      string
	ImageSrcs []string

	EnTitle string // English title, used for Chinese page
}

type productList struct {
	RstPath string
	Title   string
	Url     string
	Items   []productItem

	EnTitle string // English title, used for Chinese page
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

		tags := "product, " + titleToSlug(l.Title)
		summary := item.Title + " - " + l.Title
		s := rstMeta(item.Title, titleToSlug(item.Title),
			tags, summary,
			"en", "", item.Href, ogImg)
		targetPath := getFinalProductRstPath(l.RstPath, item.Title)
		fmt.Println(s)
		writeToFile(targetPath, s)
	}
}

func (l productList) OgImageRst() string {
	return ":og_image: " + getRstImagePath(l.Items[0].ImageSrcs[0])
}

func (l *productList) SetEnglishTitle(enlist productList) {
	l.EnTitle = enlist.Title

	for i, item := range enlist.Items {
		l.Items[i].EnTitle = item.Title
	}
}

func newProductList(rstpath string) productList {
	list := productList{
		RstPath: rstpath,
	}
	list.SetUrlTitle()

	return list
}
