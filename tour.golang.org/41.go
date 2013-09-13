package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	mWordCount := make(map[string]int)

	asWords := strings.Split(s, " ")
	for _, sWord := range asWords {
		mWordCount[sWord]++
	}

	return mWordCount
}

func main() {
	wc.Test(WordCount)
}
