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

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Strip punctuation: keep only letters
		cleaned := strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) {
				return r
			}
			return ' '
		}, line)
		words := strings.Fields(cleaned)
		for _, w := range words {
			w = strings.ToLower(w)
			if w != "" {
				wordCount[w]++
			}
		}
	}

	type wordEntry struct {
		word  string
		count int
	}

	entries := make([]wordEntry, 0, len(wordCount))
	for w, c := range wordCount {
		entries = append(entries, wordEntry{w, c})
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