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
		fmt.Println(err)
		return
	}
	defer file.Close()

	wordCount := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		runes := []rune(line)
		word := ""
		for _, r := range runes {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				word += string(r)
			} else if word != "" {
				wordCount[word]++
				word = ""
			}
		}
		if word != "" {
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	type wordEntry struct {
		word  string
		count int
	}

	var entries []wordEntry
	for word, count := range wordCount {
		entries = append(entries, wordEntry{word, count})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count == entries[j].count {
			return entries[i].word < entries[j].word
		}
		return entries[i].count > entries[j].count
	})

	for _, entry := range entries {
		fmt.Printf("%s: %d\n", entry.word, entry.count)
	}
}