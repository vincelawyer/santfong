package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getEXPLOSIONTitleHref(list productList, elm *goquery.Selection) (item productItem) {
	// get title of the final product
	elm.Find("a").Each(func(i int, a *goquery.Selection) {
		if i != 1 {
			return
		}
		item.Title = strings.TrimSpace(TrimSpaceNewlineInString(a.Text()))

		// get url of the final product
		href, ok := a.Attr("href")
		if ok {
			item.Href = fullUrl(list.Url, href)
		}
	})

	return
}

func get2ItemDataEXPLOSION(list productList, tr *goquery.Selection, itr int) productList {
	var srcs []string
	tr.Find("td").Each(func(i int, td *goquery.Selection) {
		//fmt.Println(td.Text())
		switch i {
		case 0, 2:
			srcss := getImageSrcs(list, td)
			srcs = append(srcs, srcss...)
		case 1, 3:
			if itr == 4 && i == 3 {
				return
			}
			if itr == 0 && i == 3 {
				return
			}
			item := getItemTitleHref(list, td)
			item.ImageSrcs = append(item.ImageSrcs, srcs...)
			srcs = []string{}
			list.Items = append(list.Items, item)
			/*
				if item.Title == "GA TYPE ; GC TYPE" {
					fmt.Println(itr, i)
				}
			*/
		}
	})

	return list
}

func get2ItemDataEXPLOSION2(list productList, tr *goquery.Selection) productList {
	var srcs []string
	tr.Find("td").Each(func(i int, td *goquery.Selection) {
		switch i {
		case 3:
			srcss := getImageSrcs(list, td)
			srcs = append(srcs, srcss...)
		case 4:
			item := getEXPLOSIONTitleHref(list, td)
			item.ImageSrcs = append(item.ImageSrcs, srcs...)
			srcs = []string{}
			list.Items = append(list.Items, item)
		}
	})

	return list
}

func getProductListData5(list productList) productList {
	fmt.Println(list.Url)

	// convert URL from big5 to utf8
	doc, err := NewDocumentFromNonUtf8Url(list.Url, "big5")
	if err != nil {
		panic(err)
	}

	// get links of final product
	table := doc.Find(`table[height="195"]`).First()
	// one iteration get the link of one final product
	table.Find("tr").Each(func(i int, tr *goquery.Selection) {
		list = get2ItemDataEXPLOSION(list, tr, i)
		if i == 0 {
			list = get2ItemDataEXPLOSION2(list, tr)
		}
	})

	return list
}
