package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func processTrItemImages2(tr *goquery.Selection) string {
	var tdimg *goquery.Selection
	itemImagesRst := ""

	tr.Find("td").Each(func(i int, td *goquery.Selection) {
		if i == 1 || i == 5 {
			tdimg = td
		}
		if i == 2 || i == 6 {
			itemImagesRst += processItem(td)
			itemImagesRst += processImages(tdimg)
		}
	})

	return itemImagesRst
}

func CrawlChineseProduct() {
	pageSrc := "http://shenfang.com.tw/product.htm"
	doc, err := NewDocumentFromNonUtf8Url(pageSrc, "big5")
	if err != nil {
		panic(err)
	}

	zhProductRst := rstMeta("產品", "product", "product, 產品, 產品分類", "公司產品", "zh", "2", pageSrc)

	// each tr consists of two (item, images)
	doc.Find("tr").Each(func(_ int, tr *goquery.Selection) {
		zhProductRst += processTrItemImages2(tr)
	})

	fmt.Print(zhProductRst)
}
