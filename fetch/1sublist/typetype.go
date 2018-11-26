package main

import (
	"fmt"
)

func parseTypeType(url string) {
	lines, err := NonUtf8UrlToLines(url, "big5")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
