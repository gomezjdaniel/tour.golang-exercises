package main

import (
	"strings"
	"tour/wc"
)

func WordCount(s string) map[string]int {

	subStrings := strings.Fields(s)
	m := make(map[string]int)

	for _, value := range subStrings {
		m[value]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
