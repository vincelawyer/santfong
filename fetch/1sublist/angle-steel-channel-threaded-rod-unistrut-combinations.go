package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func get2ItemDataANGLE(list productList, tr *goquery.Selection, itr int) productList {
	var srcs []string
	tr.Find("td").Each(func(i int, td *goquery.Selection) {
		//fmt.Println(td.Html())
		switch i {
		case 0, 3:
			//fmt.Println(i, "@@@")
			srcss := getImageSrcs(list, td)
			srcs = append(srcs, srcss...)
		case 2, 5:
			//fmt.Println(i, "@@@")
			item := getItemTitleHref(list, td)
			if itr == 3 && i == 5 {
				item.Title = "UNI STRUT COMBINATIONS 2"
			}
			if itr == 6 && i == 5 {
				item.Title = "UNI STRUT COMBINATIONS 3"
			}
			item.ImageSrcs = append(item.ImageSrcs, srcs...)
			srcs = []string{}
			list.Items = append(list.Items, item)
		}
	})

	return list
}

func getProductListData8(list productList) productList {
	fmt.Println(list.Url)

	// convert URL from big5 to utf8
	doc, err := NewDocumentFromNonUtf8Url(list.Url, "big5")
	if err != nil {
		panic(err)
	}

	// get links of final product
	table := doc.Find(`table[height="204"]`).First()
	// one iteration get the link of one final product
	table.Find("tr").Each(func(i int, tr *goquery.Selection) {
		fmt.Println(i, "@@@")
		list = get2ItemDataANGLE(list, tr, i)
	})

	return list
}
