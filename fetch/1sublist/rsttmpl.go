package main

import (
	"fmt"
)

func rstMeta(title, slug, tags, summary, lang, order, source, ogImage string) string {
	l1 := ":title: " + title + "\n"
	l2 := ":slug: " + slug + "\n"
	l3 := ":tags: " + tags + "\n"
	l4 := ":summary: " + summary + "\n"
	l5 := ":lang: " + lang + "\n"

	l6 := ":order: " + order + "\n"
	if order == "" {
		l6 = ""
	}

	l7 := ""
	if order == "" {
		l7 = ":status: hidden\n"
	}

	l8 := ":source: " + source + "\n"
	l9 := ":og_image: " + ogImage + "\n"

	return l1 + l2 + l3 + l4 + l5 + l6 + l7 + l8 + l9
}

func (i productItem) ToRstList() (rst string) {
	rst += "\n"

	frstpath := ""
	if i.EnTitle == "" {
		frstpath = getRstFinalProductLocalLink(i.Title)
	} else {
		frstpath = getRstFinalProductLocalLink(i.EnTitle)
	}
	rst += fmt.Sprintf("- `%s <%s>`_\n", i.Title, frstpath)

	for _, imgUrl := range i.ImageSrcs {
		rst += "\n"
		rst += fmt.Sprintf("  .. image:: %s\n", getRstImagePath(imgUrl))
		rst += fmt.Sprintf("     :name: %s\n", imgUrl)
		rst += fmt.Sprintf("     :alt: product\n")
		rst += fmt.Sprintf("     :class: product-image-thumbnail\n")
	}
	return
}
