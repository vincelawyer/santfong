package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Get info of one final product
func getProductItemDataTRAY(list productList, tr *goquery.Selection) productList {
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

	if item.Title == "" {
		return list
	}

	list.Items = append(list.Items, item)
	return list
}

func getProductListData9(list productList) productList {
	fmt.Println(list.Url)

	// convert URL from big5 to utf8
	doc, err := NewDocumentFromNonUtf8Url(list.Url, "big5")
	if err != nil {
		panic(err)
	}

	// get links of final product
	table := doc.Find(`table[height="237"]`).First()
	// one iteration get the link of one final product
	table.Find("tr").Each(func(i int, tr *goquery.Selection) {
		fmt.Println(i, "@@@")
		list = getProductItemDataTRAY(list, tr)
	})

	fmt.Println("item number: ", len(list.Items))
	fmt.Println(checkReplicate(list))

	return list
}

func LoadCABLETRAYSJSON() productList {
	b, err := ioutil.ReadFile("cable-trays.json")
	if err != nil {
		panic(err)
	}

	list := productList{}
	err = json.Unmarshal(b, &list)
	if err != nil {
		panic(err)
	}

	return list
}

func SetCABLETRAYSChinese(enlist, zhlist productList) productList {
	for i, item := range enlist.Items {
		zhlist.Items[i].EnTitle = item.Title
		zhlist.Items[i].ImageSrcs = item.ImageSrcs
	}

	return zhlist
}
