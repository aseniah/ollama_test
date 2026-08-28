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
		panic(err)
	}
	defer f.Close()

	counts := make(map[string]int)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		raw := scanner.Text()
		var b strings.Builder
		for _, r := range strings.ToLower(raw) {
			if unicode.IsLetter(r) {
				b.WriteRune(r)
			}
		}
		word := b.String()
		if word != "" {
			counts[word]++
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
