package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	data, err := os.ReadFile("input/text.txt")
	if err != nil {
		panic(err)
	}

	// Convert entire content to lowercase first
	content := strings.ToLower(string(data))

	// Extract sequences of letters only
	re := regexp.MustCompile(`[a-z]+`)
	words := re.FindAllString(content, -1)

	// Count frequencies
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	// Prepare for sorting
	type wc struct {
		word  string
		count int
	}
	var result []wc
	for w, c := range counts {
		result = append(result, wc{word: w, count: c})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(result, func(i, j int) bool {
		if result[i].count != result[j].count {
			return result[i].count > result[j].count
		}
		return result[i].word < result[j].word
	})

	// Output results
	for _, r := range result {
		fmt.Printf("%s: %d\n", r.word, r.count)
	}
}