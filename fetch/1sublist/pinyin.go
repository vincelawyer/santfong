package main

import (
	"unicode"

	"github.com/mozillazg/go-pinyin"
)

var a = pinyin.NewArgs()

func zhCharToPinyin(p string) (s string) {
	for _, r := range p {
		if unicode.Is(unicode.Han, r) {
			s += string(pinyin.Pinyin(string(r), a)[0][0])
		} else {
			s += string(r)
		}
	}
	return
}
