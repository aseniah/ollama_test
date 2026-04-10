package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	wordCounts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`[^a-z]+`)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		line = re.ReplaceAllString(line, "")
		if len(line) == 0 {
			continue
		}
		word := strings.TrimRight(line, " ")
		wordCounts[word]++
	}

	type wordEntry struct {
		word string
		count int
	}

	var entries []wordEntry
	for word, count := range wordCounts {
		entries = append(entries, wordEntry{word, count})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count != entries[j].count {
			return entries[i].count > entries[j].count
		}
		return strings.Compare(entries[i].word, entries[j].word) < 0
	})

	for _, entry := range entries {
		fmt.Printf("%s: %d\n", entry.word, entry.count)
	}
}