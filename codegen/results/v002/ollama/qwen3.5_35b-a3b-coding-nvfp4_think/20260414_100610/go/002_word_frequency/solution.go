package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	// Read the file
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// Count word frequencies
	wordCounts := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			// Convert to lowercase and strip punctuation
			cleanWord := strings.ToLower(word)
			cleanWord = strings.Map(func(r rune) rune {
				if unicode.IsLetter(r) {
					return r
				}
				return -1
			}, cleanWord)

			if cleanWord != "" {
				wordCounts[cleanWord]++
			}
		}
	}

	// Convert to slice for sorting
	type wordCount struct {
		word  string
		count int
	}

	var wcList []wordCount
	for word, count := range wordCounts {
		wcList = append(wcList, wordCount{word, count})
	}

	// Sort by count descending, then alphabetically ascending for ties
	sort.Slice(wcList, func(i, j int) bool {
		if wcList[i].count != wcList[j].count {
			return wcList[i].count > wcList[j].count
		}
		return wcList[i].word < wcList[j].word
	})

	// Output results
	for _, wc := range wcList {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}