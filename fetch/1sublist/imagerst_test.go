package main

import (
	"fmt"
	"testing"
)

const ulimg = `
.. image:: {filename}/images/ul-mark.png
   :alt: UL LISTED
   :class: img-fluid ul-max-width
`

func TestGetAllImageRst(t *testing.T) {
	enrstpath := "../../content/pages/en/product/conduit-fittings/90-m-f-elbows-90-m-m-elbows.rst"
	zhrstpath := getChineseRstPath(enrstpath)

	url := getUrlInRst(enrstpath)
	imgurls := getAllImageUrlFromWebpage(url)
	rstImgs := createImageRstFromUrl(imgurls)
	//rstImgs += ulimg
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
	//zhrstImgs += ulimg
	zhrstTables := parseTypeType(zhurl)
	AppendStringToFile(zhrstpath, zhrstImgs)
	AppendStringToFile(zhrstpath, zhrstTables)
}
