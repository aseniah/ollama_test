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
	f, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	freq := make(map[string]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for _, field := range strings.Fields(line) {
			var b strings.Builder
			for _, r := range field {
				if unicode.IsLetter(r) {
					b.WriteRune(unicode.ToLower(r))
				}
			}
			word := b.String()
			if word != "" {
				freq[word]++
			}
		}
	}

	type entry struct {
		word  string
		count int
	}
	entries := make([]entry, 0, len(freq))
	for w, c := range freq {
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
