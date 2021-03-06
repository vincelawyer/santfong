package main

import (
	"regexp"
	"strings"
)

func checkSlug(s string) {
	re := regexp.MustCompile("[a-z-]+")
	if re.ReplaceAllString(s, "") != "" {
		panic("bad slug: " + s)
	}
}

func titleToSlug(s string) (slug string) {
	slug = strings.TrimSpace(s)
	slug = strings.ToLower(slug)

	patterns := []string{` *& *`, ` *、 *`, " +"}
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		slug = re.ReplaceAllString(slug, "-")
	}
	checkSlug(slug)
	return
}

func zhSummaryToTitleSlug(summary string) (title, slug string) {
	ll := strings.Split(summary, " / ")
	title = strings.TrimSpace(ll[0])
	slug = titleToSlug(ll[1])
	return
}
