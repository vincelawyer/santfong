package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getUrlInRst(rstpath string) (url string) {
	_, err := replaceLines(rstpath, func(s string) string {
		if strings.HasPrefix(s, ":source: ") {
			url = strings.TrimPrefix(s, ":source: ")
		}
		return s
	})
	if err != nil {
		panic(err)
	}
	return
}

func getAllImageUrlFromWebpage(url string) (imgurls []string) {
	// convert URL from big5 to utf8
	doc, err := NewDocumentFromNonUtf8Url(url, "big5")
	if err != nil {
		panic(err)
	}

	doc.Find(`img`).Each(func(i int, img *goquery.Selection) {
		src, ok := img.Attr("src")
		if ok {
			downloadImage(fullUrl(url, src))
			if strings.HasSuffix(src, "LOGO.JPG") {
				return
			}
			imgurls = append(imgurls, fullUrl(url, src))
		}
	})

	return
}

func createImageRstFromUrl(imgurls []string) (rst string) {
	rst += "\n"

	for _, imgUrl := range imgurls {
		rst += "\n"
		rst += fmt.Sprintf(".. image:: %s\n", getRstImagePath(imgUrl))
		rst += fmt.Sprintf("   :name: %s\n", imgUrl)
		rst += fmt.Sprintf("   :alt: product\n")
		rst += fmt.Sprintf("   :class: img-fluid\n")
	}
	return
}
