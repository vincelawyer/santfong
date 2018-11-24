package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func get2ItemDataFLEXIBLE(list productList, tr *goquery.Selection) productList {
	var srcs []string
	tr.Find("td").Each(func(i int, td *goquery.Selection) {
		//fmt.Println(td.Html())
		switch i {
		case 1, 4:
			//fmt.Println(i, "@@@")
			srcss := getImageSrcs(list, td)
			srcs = append(srcs, srcss...)
		case 2, 5:
			//fmt.Println(i, "@@@")
			item := getItemTitleHref(list, td)
			item.ImageSrcs = append(item.ImageSrcs, srcs...)
			srcs = []string{}
			list.Items = append(list.Items, item)
		}
	})

	return list
}

func getProductListData4(list productList) productList {
	fmt.Println(list.Url)

	// convert URL from big5 to utf8
	doc, err := NewDocumentFromNonUtf8Url(list.Url, "big5")
	if err != nil {
		panic(err)
	}

	// get links of final product
	table := doc.Find(`table[height="246"]`).First()
	// one iteration get the link of one final product
	table.Find("tr").Each(func(i int, tr *goquery.Selection) {
		list = get2ItemDataFLEXIBLE(list, tr)
	})

	return list
}
