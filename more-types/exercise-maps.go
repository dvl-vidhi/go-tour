package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)

	s_arr := strings.Split(s, " ")

	for _, val := range s_arr {
		/**
		 * When a key doesn't exist in the map, then it will return 0
		 * and as we are incrementing by 1, the count will start from 1.
		 */
		wordCount[val] += 1
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
