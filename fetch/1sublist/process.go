package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// Get info of one final product
func getProductItemData(list productList, tr *goquery.Selection) productList {
	item := productItem{}

	// get title of the final product
	a := tr.Find("a").First()
	item.Title = strings.TrimSpace(TrimSpaceNewlineInString(a.Text()))

	// get url of the final product
	href, ok := a.Attr("href")
	if ok {
		item.Href = fullUrl(list.Url, href)
	}

	// get image links of the final product
	tr.Find("img").Each(func(_ int, img *goquery.Selection) {
		src, ok := img.Attr("src")
		if ok {
			item.ImageSrcs = append(item.ImageSrcs, fullUrl(list.Url, src))
			downloadImage(fullUrl(list.Url, src))
		}
	})

	list.Items = append(list.Items, item)
	return list
}

func getProductListData(list productList) productList {
	fmt.Println(list.Url)

	// convert URL from big5 to utf8
	doc, err := NewDocumentFromNonUtf8Url(list.Url, "big5")
	if err != nil {
		panic(err)
	}

	// get links of final product
	table := doc.Find(`table[height="223"]`).First()
	// one iteration get the link of one final product
	table.Find("tr").Each(func(i int, tr *goquery.Selection) {
		//fmt.Println(i, "@@@")
		list = getProductItemData(list, tr)
	})

	return list
}

func writeAll(list productList) {
	rstAll := list.OgImageRst() + "\n\n"
	for _, item := range list.Items {
		rstAll += item.ToRstList()
	}
	AppendStringToFile(list.RstPath, rstAll)

	list.CreateFinalProductRstFiles()
}

func main() {
	enrstpath := "../../content/pages/en/product/conduit-fittings/list.rst"
	zhrstpath := getChineseRstPath(enrstpath)

	enlist := newProductList(enrstpath)
	enlist = getProductListData6(enlist)
	PrettyPrint(enlist)
	writeAll(enlist)

	zhlist := newProductList(zhrstpath)
	zhlist = getProductListData6(zhlist)
	zhlist.SetEnglishTitle(enlist)
	PrettyPrint(zhlist)
	writeAll(zhlist)
}
