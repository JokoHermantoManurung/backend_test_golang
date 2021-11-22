package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	indexOpenBracket := strings.Index(str, "(")

	if indexOpenBracket < 0 {
		return ""
	}
	indexOpenBracket++

	indexClosingBracket := strings.Index(str[indexOpenBracket:], ")")

	if indexClosingBracket < 0 {
		return ""
	}

	return str[indexOpenBracket : indexOpenBracket+indexClosingBracket]
}

func main() {
	var a = findFirstStringInBracket("joko")
	fmt.Println(a)
}
