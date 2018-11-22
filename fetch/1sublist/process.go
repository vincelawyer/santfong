package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type productItem struct {
	Title     string
	Href      string
	ImageSrcs []string
}

type productList struct {
	RstPath string
	Items   []productItem
}

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
	s := rstMeta(title, titleToSlug(title),
		"product, "+parentTitle,
		title+" - "+parentTitle,
		"en", "", href, ogImg)
	targetPath := getFinalProductRstPath(parentRstPath, title)
	fmt.Print(s)
	writeToFile(targetPath, s)
}

// Get info of one final product
func processTr(parentUrl, parentTitle, parentRstPath string, tr *goquery.Selection) (rst string) {
	// get title of the final product
	a := tr.Find("a").First()
	title := strings.TrimSpace(TrimSpaceNewlineInString(a.Text()))

	// get url of the final product
	href, ok := a.Attr("href")
	if ok {
		href = fullUrl(parentUrl, href)
	}

	// get image link of the final product
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

func handleProductList(list productList) {
	parentUrl, parentTitle := getParentUrlTitle(list.RstPath)
	fmt.Println(parentUrl)

	// convert URL from big5 to utf8
	doc, err := NewDocumentFromNonUtf8Url(parentUrl, "big5")
	if err != nil {
		panic(err)
	}

	// get links of final product
	rstAll := ""
	table := doc.Find("#AutoNumber3").First()
	// one iteration get the link of one final product
	table.Find("tr").Each(func(_ int, tr *goquery.Selection) {
		rstAll += processTr(parentUrl, parentTitle, list.RstPath, tr)
	})

	// add og:image metadata
	rstAll = sublistOgImage + "\n\n" + rstAll

	// append rst back to en product list
	AppendStringToFile(list.RstPath, rstAll)
}

func main() {
	list := productList{
		RstPath: "../../content/pages/en/product/conduit-pipe/list.rst",
	}
	handleProductList(list)
}
