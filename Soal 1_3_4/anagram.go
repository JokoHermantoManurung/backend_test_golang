package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagram(arr []string) {
	list := make(map[string][]string)

	for _, word := range arr {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		for _, w := range words {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func main() {
	var array_string = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	anagram(array_string)
}
