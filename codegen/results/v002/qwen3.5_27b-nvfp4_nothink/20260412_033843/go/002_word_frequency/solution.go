package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase immediately
		line = strings.ToLower(line)

		// Split into words based on non-letter characters
		words := strings.FieldsFunc(line, func(r rune) bool {
			return !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'))
		})

		for _, w := range words {
			if len(w) > 0 {
				wordCount[w]++
			}
		}
	}

	// Convert map to slice of pairs for sorting
	type wordPair struct {
		Word  string
		Count int
	}
	pairs := make([]wordPair, 0, len(wordCount))
	for w, c := range wordCount {
		pairs = append(pairs, wordPair{w, c})
	}

	// Sort: count descending, then word ascending
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Count != pairs[j].Count {
			return pairs[i].Count > pairs[j].Count // Descending count
		}
		return pairs[i].Word < pairs[j].Word // Ascending word
	})

	// Output
	for _, p := range pairs {
		fmt.Printf("%s: %d\n", p.Word, p.Count)
	}
}