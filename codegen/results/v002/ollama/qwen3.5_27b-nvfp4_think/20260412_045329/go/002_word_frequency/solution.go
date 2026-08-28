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
		panic(err)
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := strings.Fields(line)
		for _, word := range words {
			cleanedWord := cleanWord(word)
			if cleanedWord != "" {
				wordCount[cleanedWord]++
			}
		}
	}

	type wordFreq struct {
		word  string
		count int
	}

	var freqList []wordFreq
	for word, count := range wordCount {
		freqList = append(freqList, wordFreq{word: word, count: count})
	}

	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].count != freqList[j].count {
			return freqList[i].count > freqList[j].count
		}
		return freqList[i].word < freqList[j].word
	})

	for _, wf := range freqList {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}

func cleanWord(word string) string {
	var builder strings.Builder
	for _, char := range word {
		if unicode.IsLetter(char) {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}