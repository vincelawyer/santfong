package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getItemTitleHref(list productList, elm *goquery.Selection) (item productItem) {
	// get title of the final product
	a := elm.Find("a").First()
	item.Title = strings.TrimSpace(TrimSpaceNewlineInString(a.Text()))

	// get url of the final product
	href, ok := a.Attr("href")
	if ok {
		item.Href = fullUrl(list.Url, href)
	}

	return
}

func getImageSrc(list productList, elm *goquery.Selection) string {
	img := elm.Find("img").First()
	src, ok := img.Attr("src")
	if ok {
		downloadImage(fullUrl(list.Url, src))
		return fullUrl(list.Url, src)
	}

	panic("not here: getImageSrc")
}

func get2ItemData(list productList, tr *goquery.Selection) productList {
	var srcs []string
	tr.Find("td").Each(func(i int, td *goquery.Selection) {
		//fmt.Println(td.Html())
		switch i {
		case 0, 1, 4, 5:
			src := getImageSrc(list, td)
			srcs = append(srcs, src)
		case 2, 6:
			item := getItemTitleHref(list, td)
			item.ImageSrcs = append(item.ImageSrcs, srcs...)
			srcs = []string{}
			list.Items = append(list.Items, item)
		}
	})

	return list
}

func getProductListData2(list productList) productList {
	fmt.Println(list.Url)

	// convert URL from big5 to utf8
	doc, err := NewDocumentFromNonUtf8Url(list.Url, "big5")
	if err != nil {
		panic(err)
	}

	// get links of final product
	table := doc.Find(`table[height="203"]`).First()
	// one iteration get the link of one final product
	table.Find("tr").Each(func(i int, tr *goquery.Selection) {
		if i != 0 && i != 3 {
			return
		}
		list = get2ItemData(list, tr)
	})

	return list
}
