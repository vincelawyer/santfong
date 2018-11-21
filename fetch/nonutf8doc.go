package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

func NewDocumentFromNonUtf8Url(url, charset string) (doc *goquery.Document, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	utfBody, err := iconv.NewReader(resp.Body, charset, "utf-8")
	if err != nil {
		return
	}

	doc, err = goquery.NewDocumentFromReader(utfBody)
	return
}
