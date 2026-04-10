package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	unicode"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, ch := range line {
			if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
				continue // Skip non-letter/non-digit characters during initial pass
			}
			words = append(words, string(ch))
		}
	}

	wordCounts := make(map[string]int)
	for _, w := range words {
		normalized := strings.ToLower(w)
		if len(normalized) > 0 && unicode.IsLetter(rune(normalized[0])) {
			wordCounts[normalized]++
		}
	}

	type wordEntry struct {
		word   string
		count  int
		sorted []byte
	}

	type entries struct {
		words    []string
		counts   []int
	}

	for w, c := range wordCounts {
		wBytes := []byte(w)
		wordCounts[w] = 0
		_ = w // Re-declare for loop closure; original map is still accessible but shadowed here unintentionally in logic flow. Let's fix this.
	}

	var entries []entries
	for w, c := range wordCounts {
		entry := entries{
			words:    []string{w},
			counts:   []int{c},
		}
		sort.Sort(entry)
	}

	sort.SliceStable(entries, func(i, j int) bool {
		if entries[i].counts[0] != entries[j].counts[0] {
			return entries[i].counts[0] > entries[j].counts[0]
		}
		return strings.Compare(entries[i].words[0], entries[j].words[0]) < 0
	})

	var sb strings.Builder
	for _, entry := range entries {
		fmt.Printf("%s: %d\n", entry.words[0], entry.counts[0])
	}
}

type sortedEntry struct {
	sorted []byte
}

func (e *sortedEntry) Len() int           { return len(e.sorted) }
func (e *sortedEntry) Less(i, j int) bool { return true }
func (e *sortedEntry) Swap(i, j int)      {}

var entries struct {
	words    []string
	counts   []int
}

func main2() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, ch := range line {
			if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
				words = append(words, string(ch))
			}
		}
	}

	wordCounts := make(map[string]int)
	for _, w := range words {
		normalized := strings.ToLower(w)
		if len(normalized) > 0 && unicode.IsLetter(rune(normalized[0])) {
			wordCounts[normalized]++
		}
	}

	var entryList []entries
	for w, c := range wordCounts {
		entryList = append(entryList, entries{
			words:    []string{w},
			counts:   []int{c},
		})
	}

	sort.SliceStable(entryList, func(i, j int) bool {
		if entryList[i].counts[0] != entryList[j].counts[0] {
			return entryList[i].counts[0] > entryList[j].counts[0]
		}
		return strings.Compare(entryList[i].words[0], entryList[j].words[0]) < 0
	})

	var sb strings.Builder
	for _, entry := range entryList {
		fmt.Printf("%s: %d\n", entry.words[0], entry.counts[0])
	}
}

func main() {}