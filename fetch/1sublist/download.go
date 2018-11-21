package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func isExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		// exist
		return true
	}
	// not exist
	return false
}

func download(url, filepath string) (err error) {
	fmt.Println("Downloading ", url, " to ", filepath)
	if isExist(filepath) {
		fmt.Println(filepath, "already exists!")
		return
	}
	createDirIfNotExist(path.Dir(filepath))

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return
}

func downloadImage(url string) {
	download(url, getLocalImagePathFromUrl(url))
}
