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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		for _, f := range fields {
			// Convert to lowercase and keep only letters
			var b strings.Builder
			for _, ch := range strings.ToLower(f) {
				if unicode.IsLetter(ch) {
					b.WriteRune(ch)
				}
			}
			w := b.String()
			if w != "" {
				words = append(words, w)
			}
		}
	}

	// Count frequencies
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	// Build slice of key-value pairs
	type kv struct {
		word  string
		count int
	}
	var pairs []kv
	for w, c := range counts {
		pairs = append(pairs, kv{w, c})
	}

	// Sort: count descending, then word alphabetically ascending
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})

	// Output
	for _, p := range pairs {
		fmt.Printf("%s: %d\n", p.word, p.count)
	}
}