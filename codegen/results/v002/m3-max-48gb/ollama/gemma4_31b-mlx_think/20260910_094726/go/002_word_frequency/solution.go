package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// wordCount stores a word and its corresponding frequency.
type wordCount struct {
	word  string
	count int
}

func main() {
	// Read the file content
	content, err := os.ReadFile("input/text.txt")
	if err != nil {
		return
	}

	// Convert all words to lowercase
	text := strings.ToLower(string(content))

	// Use a regular expression to find all sequences of letters.
	// This effectively strips all punctuation and numbers, keeping only letters.
	re := regexp.MustCompile(`[a-z]+`)
	words := re.FindAllString(text, -1)

	// Count the frequency of each word using a map
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	// Convert the map to a slice of wordCount structs for sorting
	var sortedList []wordCount
	for w, c := range counts {
		sortedList = append(sortedList, wordCount{w, c})
	}

	// Sort by count descending; for ties, sort alphabetically ascending
	sort.Slice(sortedList, func(i, j int) bool {
		if sortedList[i].count != sortedList[j].count {
			return sortedList[i].count > sortedList[j].count
		}
		return sortedList[i].word < sortedList[j].word
	})

	// Output the results in the format word: count
	for _, wc := range sortedList {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}