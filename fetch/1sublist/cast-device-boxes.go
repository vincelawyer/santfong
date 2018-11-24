package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func getImageSrcs(list productList, elm *goquery.Selection) []string {
	var srcs []string

	elm.Find("img").Each(func(i int, img *goquery.Selection) {
		src, ok := img.Attr("src")
		if ok {
			downloadImage(fullUrl(list.Url, src))
			srcs = append(srcs, fullUrl(list.Url, src))
		}
	})

	return srcs
}

func get2ItemDataCAST(list productList, tr *goquery.Selection) productList {
	var srcs []string
	tr.Find("td").Each(func(i int, td *goquery.Selection) {
		//fmt.Println(td.Text())
		switch i {
		case 1, 4:
			srcss := getImageSrcs(list, td)
			srcs = append(srcs, srcss...)
		case 2, 5:
			item := getItemTitleHref(list, td)
			item.ImageSrcs = append(item.ImageSrcs, srcs...)
			srcs = []string{}
			list.Items = append(list.Items, item)
		}
	})

	return list
}

func get2ItemDataCAST2(list productList, tr *goquery.Selection) productList {
	var srcs []string
	tr.Find("td").Each(func(i int, td *goquery.Selection) {
		//fmt.Println(td.Text())
		switch i {
		case 1, 3:
			srcss := getImageSrcs(list, td)
			srcs = append(srcs, srcss...)
		case 2, 4:
			item := getItemTitleHref(list, td)
			item.ImageSrcs = append(item.ImageSrcs, srcs...)
			srcs = []string{}
			list.Items = append(list.Items, item)
		}
	})

	return list
}

func getProductListData3(list productList) productList {
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
		if i == 6 {
			list = get2ItemDataCAST2(list, tr)
			return
		}
		list = get2ItemDataCAST(list, tr)
		/*
			if list.Items[len(list.Items)-1].Title == "FC TYPE ; FEC TYPE" {
				fmt.Println(i, "@@@")
			}
		*/
	})

	return list
}
