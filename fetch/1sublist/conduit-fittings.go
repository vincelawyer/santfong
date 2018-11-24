package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func get2ItemDataCONDUITFITTINGS(list productList, tr *goquery.Selection, itr int) productList {
	var srcs []string
	tr.Find("td").Each(func(i int, td *goquery.Selection) {
		//fmt.Println(td.Html())
		switch i {
		case 1, 4:
			//fmt.Println(i, "@@@")
			srcss := getImageSrcs(list, td)
			srcs = append(srcs, srcss...)
		case 2, 5:
			if itr == 24 && i == 5 {
				return
			}
			//fmt.Println(i, "@@@")
			item := getItemTitleHref(list, td)
			if item.Title == "90° \n                  ELBOWS \u0026 45° ELBOWS" {
				item.Title = "90° ELBOWS \u0026 45° ELBOWS"
			}
			item.ImageSrcs = append(item.ImageSrcs, srcs...)
			srcs = []string{}
			list.Items = append(list.Items, item)
		}
	})

	return list
}

func getProductListData6(list productList) productList {
	fmt.Println(list.Url)

	// convert URL from big5 to utf8
	doc, err := NewDocumentFromNonUtf8Url(list.Url, "big5")
	if err != nil {
		panic(err)
	}

	// get links of final product
	table := doc.Find(`table[height="199"]`).First()
	// one iteration get the link of one final product
	table.Find("tr").Each(func(i int, tr *goquery.Selection) {
		//fmt.Println(i, "@@@")
		if i == 1 || i == 4 || i == 7 || i == 11 || i == 14 || i == 17 || i == 21 || i == 24 {
			list = get2ItemDataCONDUITFITTINGS(list, tr, i)
		}
	})

	return list
}
