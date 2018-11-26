package main

import (
	"fmt"
	"strings"
)

func isUnitLine(lineno int, line string) bool {
	if strings.HasPrefix(line, `<p align="right"`) {
		return true
	}
	return false
}

func isTableEnd(lineno int, line string) bool {
	if line == `</table>` {
		return true
	}
	return false
}

func extractTablesHtml(lines []string) (tables []string) {
	table := ""
	isInTable := false
	for i, line := range lines {
		if isUnitLine(i, line) {
			isInTable = true
		}
		if isInTable {
			table = table + line + "\n"
		}
		if isTableEnd(i, line) {
			if isInTable == true {
				isInTable = false
				tables = append(tables, table)
				table = ""
			}
		}
	}
	return
}

func createTablesRst(lines []string) (rst string) {
	tables := extractTablesHtml(lines)
	//fmt.Println(tables)
	//fmt.Println(len(tables))

	for _, table := range tables {
		rst += ".. raw:: html\n\n"

		r := strings.NewReader(table)
		lines2, err := LinesFromReader(r)
		if err != nil {
			panic(err)
		}

		for _, line2 := range lines2 {
			rst = rst + "  " + line2 + "\n"
		}

		rst += "\n\n"
	}

	return
}

func parseTypeType(url string) {
	lines, err := NonUtf8UrlToLines(url, "big5")
	if err != nil {
		panic(err)
	}

	rst := createTablesRst(lines)
	fmt.Println(rst)
}
