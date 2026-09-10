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
		os.Exit(1)
	}
	defer file.Close()

	counts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		for _, field := range fields {
			// Strip non-letter characters, keep only letters
			var sb strings.Builder
			for _, ch := range field {
				if unicode.IsLetter(ch) {
					sb.WriteRune(unicode.ToLower(ch))
				}
			}
			word := sb.String()
			if word != "" {
				counts[word]++
			}
		}
	}

	// Build slice of (word, count) pairs
	type wc struct {
		word  string
		count int
	}
	pairs := make([]wc, 0, len(counts))
	for w, c := range counts {
		pairs = append(pairs, wc{w, c})
	}

	// Sort by count descending, then alphabetically ascending
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