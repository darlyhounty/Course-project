package main

import (
	"fmt"
	"strings"
)

func frequencySort(s string) string {
	charMap := make(map[rune]int)
	arr := make([]string, len(s))
	for _, c := range s {
		charMap[c]++
	}
	for k, v := range charMap {
		arr[v-1] += strings.Repeat(string(k), v)
	}
	r := ""
	for _, c := range arr {
		r = c + r
	}

	return r
}
func main() {
	var s string
	fmt.Scan(&s)
	frequencySort(s)
	fmt.Print(frequencySort(s))
}
