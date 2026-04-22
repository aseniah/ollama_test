package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("input/text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	counts := make(map[string]int)
	words := strings.Fields(string(data))
	for _, w := range words {
		w = strings.ToLower(w)
		w = strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) {
				return r
			}
			return -1
		}, w)
		if w == "" {
			continue
		}
		counts[w]++
	}

	sortedWords := make([]string, 0, len(counts))
	for w := range counts {
		sortedWords = append(sortedWords, w)
	}

	sort.Slice(sortedWords, func(i, j int) bool {
		ci := counts[sortedWords[i]]
		cj := counts[sortedWords[j]]
		if ci != cj {
			return ci > cj
		}
		return sortedWords[i] < sortedWords[j]
	})

	for _, w := range sortedWords {
		fmt.Printf("%s: %d\n", w, counts[w])
	}
}