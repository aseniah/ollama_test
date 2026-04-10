package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Normalize case
		line = strings.ToLower(line)

		// Split into runes to handle punctuation stripping manually
		words := splitAndClean(line)
		for _, w := range words {
			if len(w) > 0 {
				wordCounts[w]++
			}
		}
	}

	// Convert map to slice for sorting
	type wordCount struct {
		word string
		count int
	}
	
	var wcList []wordCount
	for w, c := range wordCounts {
		wcList = append(wcList, wordCount{w, c})
	}

	// Sort: descending by count, then ascending alphabetically for ties
	sortWords(wcList)

	// Output results
	for _, item := range wcList {
		fmt.Printf("%s: %d\n", item.word, item.count)
	}
}

func splitAndClean(s string) []string {
	var result []string
	var current strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) {
			current.WriteRune(r)
		} else {
			if current.Len() > 0 {
				result = append(result, current.String())
				current.Reset()
			}
		}
	}

	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result
}

func sortWords(list []wordCount) {
	// Since Go doesn't have a built-in sort for this specific dual criteria in stdlib slices 
	// without importing "sort" and defining a custom less function, we do it manually or use sort.Slice.
	// We will implement a simple bubble/insertion sort for clarity or just use the sort package.
	// Let's use the sort package to be idiomatic.
	
	// We need to import "sort", so let's add it implicitly via the logic below.
	// However, to keep the code self-contained without modifying imports inside this function scope:
	// I will perform a manual sort here to ensure no extra imports are needed outside main block 
	// (though 'sort' is standard, adding imports in the middle of execution is messy).
	// Actually, let's just write a simple bubble sort for this specific requirement since performance 
	// on text files usually isn't critical enough to require introspection sort.
	
	n := len(list)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			// Compare counts: if list[j].count > list[minIdx].count, j is "smaller" (higher priority)
			if list[j].count > list[minIdx].count {
				minIdx = j
			} else if list[j].count == list[minIdx].count && list[j].word < list[minIdx].word {
				minIdx = j
			}
		}
		list[i], list[minIdx] = list[minIdx], list[i]
	}
}