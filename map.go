package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	splitWords := strings.Fields(s)
	dic := make(map[string]int, len(splitWords))
	
	for _, v := range splitWords{
		dic[v] += 1
	}
	
	return dic
}

func main() {
	wc.Test(WordCount)
}

