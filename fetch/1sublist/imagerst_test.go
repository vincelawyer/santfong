package main

import (
	"fmt"
	"testing"
)

func TestGetAllImageRst(t *testing.T) {
	enrstpath := "../../content/pages/en/product/cast-device-boxes/fl-type-fr-type.rst"
	zhrstpath := getChineseRstPath(enrstpath)

	url := getUrlInRst(enrstpath)
	parseTypeType(url)
	return

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
