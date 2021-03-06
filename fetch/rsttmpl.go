package main

func rstMeta(title, slug, tags, summary, lang, order, source string) string {
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

	return l1 + l2 + l3 + l4 + l5 + l6 + l7 + l8 + "\n\n"
}

func rstListItem(text, url string) string {
	return "- `" + text + " <" + url + ">`_\n"
}

func rstImage(url string) string {
	l1 := "\n"
	l2 := "  .. image:: " + url + "\n"
	l3 := "     :alt: product\n"
	l4 := "     :class: product-image-thumbnail\n"

	return l1 + l2 + l3 + l4
}
