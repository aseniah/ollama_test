package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type WordCount struct {
	word  string
	count int
}

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	wordCounts := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(word)
			reg, _ := regexp.Compile("[^a-z]+")
			cleanWord := reg.ReplaceAllString(word, "")
			if cleanWord != "" {
				wordCounts[cleanWord]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var wordCountSlice []WordCount
	for word, count := range wordCounts {
		wordCountSlice = append(wordCountSlice, WordCount{word: word, count: count})
	}

	sort.Slice(wordCountSlice, func(i, j int) bool {
		if wordCountSlice[i].count == wordCountSlice[j].count {
			return wordCountSlice[i].word < wordCountSlice[j].word
		}
		return wordCountSlice[i].count > wordCountSlice[j].count
	})

	for _, wc := range wordCountSlice {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}