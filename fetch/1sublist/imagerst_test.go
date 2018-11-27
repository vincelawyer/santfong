package main

import (
	"fmt"
	"testing"
)

func TestGetAllImageRst(t *testing.T) {
	enrstpath := "../../content/pages/en/product/explosion-proof-round-outlet-boxes/gb-type-gd-type.rst"
	zhrstpath := getChineseRstPath(enrstpath)

	url := getUrlInRst(enrstpath)
	imgurls := getAllImageUrlFromWebpage(url)
	rstImgs := createImageRstFromUrl(imgurls)
	rstTables := parseTypeType(url)

	fmt.Println()
	fmt.Println(url)
	fmt.Println(imgurls)
	fmt.Println(rstImgs)
	AppendStringToFile(enrstpath, rstImgs)
	AppendStringToFile(enrstpath, rstTables)

	zhurl := getUrlInRst(zhrstpath)
	zhimgurls := getAllImageUrlFromWebpage(zhurl)
	zhrstImgs := createImageRstFromUrl(zhimgurls)
	zhrstTables := parseTypeType(zhurl)
	AppendStringToFile(zhrstpath, zhrstImgs)
	AppendStringToFile(zhrstpath, zhrstTables)
}
