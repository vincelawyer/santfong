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

func titleImageToRstList(title, href, imgUrl string) (rst string) {
	rst += "\n"
	rst += fmt.Sprintf("- `%s <%s>`_\n\n", title, getRstFinalProductLocalLink(title))
	rst += fmt.Sprintf("  .. image:: %s\n", getRstImagePath(imgUrl))
	rst += fmt.Sprintf("     :name: %s\n", imgUrl)
	rst += fmt.Sprintf("     :alt: product\n")
	rst += fmt.Sprintf("     :class: product-image-thumbnail\n")
	return
}
