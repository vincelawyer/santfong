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

func fullUrl(url, rel string) string {
	dir := path.Dir(url)
	tmp := path.Join(dir, rel)
	return strings.Replace(tmp, "http:/s", "http://s", 1)
}

func getRstImagePath(url string) string {
	p1 := getLocalImagePathFromUrl(url)
	p2 := strings.TrimPrefix(p1, "../../content")
	return "{filename}" + p2
}

func getRstFinalProductLocalLink(title string) string {
	slug := titleToSlug(title)
	return "{filename}" + slug + ".rst"
}

func getFinalProductRstPath(parentRstPath, title string) string {
	dir := path.Dir(parentRstPath)
	return path.Join(dir, titleToSlug(title)+".rst")
}

func getChineseRstPath(rstpath string) string {
	return strings.Replace(rstpath, "/en/", "/zh/", 1)
}
