package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func replaceLines(filepath string, repl func(string) string) (linesAfter []string, err error) {
	lines, err := FileToLines(filepath)
	if err != nil {
		return
	}

	for _, line := range lines {
		linesAfter = append(linesAfter, repl(line))
	}

	return
}

func replaceLinesAndWrite(filepath string, repl func(string) string) (linesAfter []string, err error) {
	lines, err := FileToLines(filepath)
	if err != nil {
		return
	}

	for _, line := range lines {
		linesAfter = append(linesAfter, repl(line))
	}

	err = WriteLinesToFile(linesAfter, filepath)
	return
}

func FileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func WriteLinesToFile(lines []string, filename string) (err error) {
	return ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")+"\n"), 0644)
}
