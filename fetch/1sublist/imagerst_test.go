package main

import (
	"fmt"
	"testing"
)

func TestGetAllImageRst(t *testing.T) {
	rstpath := "../../content/pages/zh/product/cd-cf-flexible-conduit-and-fittings/cd-cf-flexible-conduit-fittings.rst"
	url := getUrlInRst(rstpath)
	imgurls := getAllImageUrlFromWebpage(url)
	rstImgs := createImageRstFromUrl(imgurls)

	fmt.Println()
	fmt.Println(url)
	fmt.Println(imgurls)
	fmt.Println(rstImgs)
	AppendStringToFile(rstpath, rstImgs)
}
