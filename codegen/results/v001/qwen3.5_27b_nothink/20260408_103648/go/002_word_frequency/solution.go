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
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase
		line = strings.ToLower(line)

		// Remove punctuation and split into words
		var sb strings.Builder
		for _, r := range line {
			if unicode.IsLetter(r) {
				sb.WriteRune(r)
			} else {
				// Treat non-letters as delimiters by adding a space
				sb.WriteRune(' ')
			}
		}

		words := strings.Fields(sb.String())
		for _, w := range words {
			if w != "" {
				wordCounts[w]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Convert map to slice of key-value pairs for sorting
	type kv struct {
		Word  string
		Count int
	}
	var sortedList []kv
	for word, count := range wordCounts {
		sortedList = append(sortedList, kv{Word: word, Count: count})
	}

	// Sort by count descending, then by word ascending
	sort.Slice(sortedList, func(i, j int) bool {
		if sortedList[i].Count != sortedList[j].Count {
			return sortedList[i].Count > sortedList[j].Count
		}
		return sortedList[i].Word < sortedList[j].Word
	})

	// Output results
	for _, item := range sortedList {
		fmt.Printf("%s: %d\n", item.Word, item.Count)
	}
}