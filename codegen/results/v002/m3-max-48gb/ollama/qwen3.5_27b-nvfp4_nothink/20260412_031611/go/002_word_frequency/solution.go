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

	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			cleaned := cleanWord(word)
			if len(cleaned) > 0 {
				wordCounts[cleaned]++
			}
		}
	}

	type entry struct {
		Word string
		Cnt  int
	}

	entries := make([]entry, 0, len(wordCounts))
	for w, c := range wordCounts {
		entries = append(entries, entry{Word: w, Cnt: c})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Cnt != entries[j].Cnt {
			return entries[i].Cnt > entries[j].Cnt // Descending by count
		}
		return entries[i].Word < entries[j].Word   // Ascending by word
	})

	for _, e := range entries {
		fmt.Printf("%s: %d\n", e.Word, e.Cnt)
	}
}

func cleanWord(s string) string {
	s = strings.ToLower(s)
	var sb strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}