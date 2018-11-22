package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var sublistOgImage = ""

func getParentUrlTitle(rstpath string) (url, title string) {
	_, err := replaceLines(rstpath, func(s string) string {
		if strings.HasPrefix(s, ":title: ") {
			title = strings.TrimPrefix(s, ":title: ")
		}
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

func createFinalProductRst(parentTitle, parentRstPath, title, href, src string) {
	ogImg := getRstImagePath(src)
	s := rstMeta(title, titleToSlug(title), "product, "+parentTitle, title+" - "+parentTitle, "en", "", href, ogImg)
	targetPath := getFinalProductRstPath(parentRstPath, title)
	fmt.Print(s)
	writeToFile(targetPath, s)
}

func processTr(parentUrl, parentTitle, parentRstPath string, tr *goquery.Selection) (rst string) {
	a := tr.Find("a").First()
	title := strings.TrimSpace(TrimSpaceNewlineInString(a.Text()))

	href, ok := a.Attr("href")
	if ok {
		href = fullUrl(parentUrl, href)
	}

	img := tr.Find("img").First()
	src, ok := img.Attr("src")
	if ok {
		src = fullUrl(parentUrl, src)
		downloadImage(src)
	}

	if sublistOgImage == "" {
		sublistOgImage = ":og_image: " + getRstImagePath(src)
	}

	rst = titleImageToRstList(title, href, src)
	createFinalProductRst(parentTitle, parentRstPath, title, href, src)
	return
}

func handleProductList(rstpath string) {
	parentUrl, parentTitle := getParentUrlTitle(rstpath)
	fmt.Println(parentUrl)

	// convert URL from big5 to utf8
	doc, err := NewDocumentFromNonUtf8Url(parentUrl, "big5")
	if err != nil {
		panic(err)
	}

	// get links of final product
	rstAll := ""
	table := doc.Find("#AutoNumber3").First()
	// one iteration get one link of final product
	table.Find("tr").Each(func(_ int, tr *goquery.Selection) {
		rstAll += processTr(parentUrl, parentTitle, rstpath, tr)
	})

	// add og:image metadata
	rstAll = sublistOgImage + "\n\n" + rstAll

	// append rst back to en product list
	AppendStringToFile(rstpath, rstAll)
}

func main() {
	rstpath := "../../content/pages/en/product/conduit-pipe/list.rst"
	handleProductList(rstpath)
}
