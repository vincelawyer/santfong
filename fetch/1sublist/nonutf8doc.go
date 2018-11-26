package main

import (
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

func NonUtf8UrlToLines(url, charset string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	utfBody, err := iconv.NewReader(resp.Body, charset, "utf-8")
	if err != nil {
		return nil, err
	}

	return LinesFromReader(utfBody)
}

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

func NewDocumentFromNonUtf8File(fpath, charset string) (doc *goquery.Document, err error) {
	f, err := os.Open(fpath)
	if err != nil {
		return
	}
	defer f.Close()

	utfBody, err := iconv.NewReader(f, charset, "utf-8")
	if err != nil {
		return
	}

	doc, err = goquery.NewDocumentFromReader(utfBody)
	return
}
