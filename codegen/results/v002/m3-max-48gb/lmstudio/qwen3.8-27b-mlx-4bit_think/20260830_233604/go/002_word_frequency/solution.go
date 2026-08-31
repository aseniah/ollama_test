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
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	counts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, w := range words {
			// Convert to lowercase
			w = strings.ToLower(w)
			// Strip all non-letter characters (keep only letters)
			var b strings.Builder
			for _, r := range w {
				if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
					b.WriteRune(r)
				}
			}
			cleaned := b.String()
			if cleaned != "" {
				counts[cleaned]++
			}
		}
	}

	type entry struct {
		word  string
		count int
	}

	entries := make([]entry, 0, len(counts))
	for w, c := range counts {
		entries = append(entries, entry{w, c})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count != entries[j].count {
			return entries[i].count > entries[j].count
		}
		return entries[i].word < entries[j].word
	})

	for _, e := range entries {
		fmt.Printf("%s: %d\n", e.word, e.count)
	}
}