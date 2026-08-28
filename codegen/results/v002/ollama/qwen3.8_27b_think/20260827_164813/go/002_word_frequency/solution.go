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

	scanner := bufio.NewScanner(file)
	wordCount := make(map[string]int)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := strings.Fields(line)
		for _, word := range words {
			cleaned := make([]rune, 0, len(word))
			for _, r := range word {
				if unicode.IsLetter(r) {
					cleaned = append(cleaned, r)
				}
			}
			c := string(cleaned)
			if c != "" {
				wordCount[c]++
			}
		}
	}

	type entry struct {
		word  string
		count int
	}
	entries := make([]entry, 0, len(wordCount))
	for w, c := range wordCount {
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