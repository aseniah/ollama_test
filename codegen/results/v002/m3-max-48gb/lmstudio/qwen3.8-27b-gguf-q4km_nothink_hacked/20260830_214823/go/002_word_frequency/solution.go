package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type wordCount struct {
	Word  string
	Count int
}

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	frequencies := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := strings.Fields(line)
		for _, w := range words {
			// Strip all non-letter characters
			var cleaned strings.Builder
			for _, ch := range w {
				if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
					cleaned.WriteRune(ch)
				}
			}
			cleanedWord := cleaned.String()
			if cleanedWord != "" {
				frequencies[cleanedWord]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Convert to slice for sorting
	var result []wordCount
	for word, count := range frequencies {
		result = append(result, wordCount{Word: word, Count: count})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(result, func(i, j int) bool {
		if result[i].Count != result[j].Count {
			return result[i].Count > result[j].Count
		}
		return result[i].Word < result[j].Word
	})

	// Output
	for _, wc := range result {
		fmt.Printf("%s: %d\n", wc.Word, wc.Count)
	}
}