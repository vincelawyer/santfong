package main

import (
	"fmt"
	"testing"
)

func TestGetAllImageRst(t *testing.T) {
	enrstpath := "../../content/pages/en/product/conduit-outlet-boxes/tb-type-t-type.rst"
	zhrstpath := getChineseRstPath(enrstpath)

	url := getUrlInRst(enrstpath)
	imgurls := getAllImageUrlFromWebpage(url)
	rstImgs := createImageRstFromUrl(imgurls)

	fmt.Println()
	fmt.Println(url)
	fmt.Println(imgurls)
	fmt.Println(rstImgs)
	AppendStringToFile(enrstpath, rstImgs)

	zhurl := getUrlInRst(zhrstpath)
	zhimgurls := getAllImageUrlFromWebpage(zhurl)
	zhrstImgs := createImageRstFromUrl(zhimgurls)
	AppendStringToFile(zhrstpath, zhrstImgs)
}
