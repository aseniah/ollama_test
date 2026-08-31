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
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Process each line
		words := strings.FieldsFunc(line, func(r rune) bool {
			return !unicode.IsLetter(r)
		})
		
		for _, word := range words {
			// Convert to lowercase
			lowerWord := strings.ToLower(word)
			// Remove any remaining punctuation
			cleanWord := strings.TrimFunc(lowerWord, func(r rune) bool {
				return !unicode.IsLetter(r)
			})
			
			if cleanWord != "" {
				wordCounts[cleanWord]++
			}
		}
	}

	// Convert map to slice for sorting
	type wordCount struct {
		word  string
		count int
	}
	
	var words []wordCount
	for word, count := range wordCounts {
		words = append(words, wordCount{word, count})
	}

	// Sort by count descending, then by word ascending
	sort.Slice(words, func(i, j int) bool {
		if words[i].count == words[j].count {
			return words[i].word < words[j].word
		}
		return words[i].count > words[j].count
	})

	// Output results
	for _, wc := range words {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}