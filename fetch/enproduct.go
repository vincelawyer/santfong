package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getItemHref(td *goquery.Selection) string {
	href := ""
	td.Find("a").Each(func(_ int, a *goquery.Selection) {
		val, ok := a.Attr("href")
		if ok {
			href = val
		}
	})
	return href
}

func processImages(td *goquery.Selection) string {
	imagesRst := ""

	td.Find("img").Each(func(_ int, img *goquery.Selection) {
		val, ok := img.Attr("src")
		if ok {
			url := "http://shenfang.com.tw/" + val
			imagesRst += rstImage(url)
		}
	})

	return imagesRst
}

func processItem(tdname *goquery.Selection) string {
	text := TrimSpaceNewlineInString(strings.TrimSpace(tdname.Text()))
	if text != "" {
		href := getItemHref(tdname)
		url := "http://shenfang.com.tw/" + href
		return rstListItem(text, url)
	}
	return ""
}

func processTrItemImages(tr *goquery.Selection) string {
	var tdimg *goquery.Selection
	itemImagesRst := ""

	tr.Find("td").Each(func(i int, td *goquery.Selection) {
		if i == 1 || i == 4 {
			tdimg = td
		}
		if i == 3 || i == 5 {
			itemImagesRst += processItem(td)
			itemImagesRst += processImages(tdimg)
		}
	})

	return itemImagesRst
}

func CrawlEnglishProduct() {
	pageSrc := "http://shenfang.com.tw/product-1.htm"
	doc, err := NewDocumentFromNonUtf8Url(pageSrc, "big5")
	if err != nil {
		panic(err)
	}

	enProductRst := rstMeta("Product", "product", "product, product category", "all company products", "en", "2", pageSrc)

	// each tr consists of two (item, images)
	doc.Find("tr").Each(func(_ int, tr *goquery.Selection) {
		enProductRst += processTrItemImages(tr)
	})

	fmt.Print(enProductRst)
}
