package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	
	// Regex to match only letters
	reg := regexp.MustCompile(`[^a-z]+`)

	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase
		line = strings.ToLower(line)
		
		// Split by whitespace
		words := strings.Fields(line)
		for _, w := range words {
			// Strip punctuation (keep only letters)
			cleaned := reg.ReplaceAllString(w, "")
			// If the word became empty after stripping, skip it
			if cleaned != "" {
				// Handle cases where punctuation might be inside or around words
				// The regex above replaces non-letters with empty strings.
				// However, if a word was "times,", it becomes "times".
				// If a word was "it's", it becomes "its".
				counts[cleaned]++
			}
		}
	}

	// Convert map to slice for sorting
	var sortedList []wordCount
	for w, c := range counts {
		sortedList = append(sortedList, wordCount{w, c})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(sortedList, func(i, j int) bool {
		if sortedList[i].count != sortedList[j].count {
			return sortedList[i].count > sortedList[j].count
		}
		return sortedList[i].word < sortedList[j].word
	})

	// Output results
	for _, wc := range sortedList {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}