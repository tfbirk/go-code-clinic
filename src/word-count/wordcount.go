// my solution to Tour of Go's maps exercise

package main

import (
	"strings"
	
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// return map[string]int{"x": 1}
	count := map[string]int{}
	
	allWords := strings.Fields(s);
	
	for _, val := range allWords {
		count[val] += 1
	}
	
	return count
}

func main() {
	wc.Test(WordCount)
}
