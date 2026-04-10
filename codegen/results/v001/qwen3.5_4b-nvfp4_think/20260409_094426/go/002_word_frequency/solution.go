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
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	wordCountMap := make(map[string]int)
	var wordSlice []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches, _ := regexp.FindAllString(line, -1)
		wordSlice = append(wordSlice, matches...)
	}

	pattern := regexp.MustCompile(`[^a-z]`)

	for _, word := range wordSlice {
		w := strings.ToLower(pattern.ReplaceAllString(word, ""))
		if len(w) > 0 {
			wordCountMap[w]++
		}
	}

	type WordPair struct {
		Name  string
		Count int
	}

	var result []WordPair
	for word, count := range wordCountMap {
		result = append(result, WordPair{Name: word, Count: count})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Count != result[j].Count {
			return result[i].Count > result[j].Count
		}
		return result[i].Name < result[j].Name
	})

	for _, item := range result {
		fmt.Printf("%s: %d\n", item.Name, item.Count)
	}
}