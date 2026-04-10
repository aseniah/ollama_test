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
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Extract only letters (a-z, A-Z) and spaces; replace others with space or ignore non-letter sequences
		// Using a regex to keep only letters and join them if they are separated by non-letters
		pattern := regexp.MustCompile(`[^a-zA-Z]+`)
		matches := pattern.Split(line, -1)
		for _, match := range matches {
			if match == "" || strings.Index(match[:len(match)-1], " ") == len(match)-1 {
				continue
			}
			words = append(words, match)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	wordMap := make(map[string]int)
	for _, word := range words {
		lowerWord := strings.ToLower(word)
		wordMap[lowerWord]++
	}

	type wordEntry struct {
		word string
		count int
	}

	var entries []wordEntry
	for w, c := range wordMap {
		entries = append(entries, wordEntry{word: w, count: c})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count != entries[j].count {
			return entries[i].count > entries[j].count
		}
		return entries[i].word < entries[j].word
	})

	outFile, err := os.Stdout.WriteString("")
	fmt.Fprintf(os.Stdout, "")
	for _, e := range entries {
		fmt.Fprintf(os.Stdout, "%s: %d\n", e.word, e.count)
	}
}