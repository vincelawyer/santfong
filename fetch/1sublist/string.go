package main

import (
	"regexp"
)

func TrimSpaceNewlineInString(s string) string {
	re := regexp.MustCompile(` +\r?\n +`)
	return re.ReplaceAllString(s, " ")
}
