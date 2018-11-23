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

		var tags, summary, slug, lang string
		if item.EnTitle == "" {
			tags = "product, " + titleToSlug(l.Title)
			summary = item.Title + " - " + l.Title
			slug = titleToSlug(item.Title)
			lang = "en"
		} else {
			tags = "product, 產品, " + titleToSlug(l.EnTitle)
			summary = fmt.Sprintf("%s (%s) - %s (%s)", item.Title, item.EnTitle, l.Title, l.EnTitle)
			slug = titleToSlug(item.EnTitle)
			lang = "zh"
		}

		s := rstMeta(item.Title, slug, tags, summary,
			lang, "", item.Href, ogImg)
		targetPath := ""
		if item.EnTitle == "" {
			targetPath = getFinalProductRstPath(l.RstPath, item.Title)
		} else {
			targetPath = getFinalProductRstPath(l.RstPath, item.EnTitle)
		}
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
