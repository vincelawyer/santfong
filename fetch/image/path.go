package main

import (
	"path"
	"regexp"
	"strings"
)

var imgRootDir = "../../content/images"

func checkPath(p string) {
	re := regexp.MustCompile("[a-z0-9-/.]+")
	pp := re.ReplaceAllString(p, "")
	if pp != "" {
		panic(pp)
	}
}

func cleanPath(pypath string) (cpath string) {
	cpath = strings.ToLower(pypath)
	cpath = strings.Replace(cpath, "%", "-", -1)
	cpath = strings.Replace(cpath, "(", "-", -1)
	cpath = strings.Replace(cpath, ")", "-", -1)
	return
}

func getLocalImagePathFromUrl(url string) (imgPath string) {
	ll := strings.Split(url, ".com.tw/")
	relpath := ll[1]
	pypath := zhCharToPinyin(relpath)
	cpath := cleanPath(pypath)
	checkPath(cpath)
	imgPath = path.Join(imgRootDir, cpath)
	return
}
