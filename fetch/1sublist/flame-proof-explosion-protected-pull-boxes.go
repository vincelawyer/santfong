package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Get info of one final product
func getProductItemDataFLAME(list productList, tr *goquery.Selection, itr int) productList {
	item := productItem{}

	// get title of the final product
	a := tr.Find("a").First()
	item.Title = strings.TrimSpace(TrimSpaceNewlineInString(a.Text()))
	if itr == 0 {
		item.Title += " 1"
	}
	if itr == 2 {
		item.Title += " 2"
	}

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

func getProductListData7(list productList) productList {
	fmt.Println(list.Url)

	var doc *goquery.Document
	var err error
	if list.Url == "http://shenfang.com.tw/022/d3-1.htm" {
		doc, err = NewDocumentFromNonUtf8File("xxx.html", "utf8")

	} else {
		doc, err = NewDocumentFromNonUtf8File("xxx2.html", "utf8")
	}
	if err != nil {
		panic(err)
	}

	// get links of final product
	table := doc.Find(`table[height="213"]`).First()
	// one iteration get the link of one final product
	table.Find("tr").Each(func(i int, tr *goquery.Selection) {
		fmt.Println(i, "@@@")
		if i%2 == 0 {
			list = getProductItemDataFLAME(list, tr, i)
		}
	})

	return list
}
