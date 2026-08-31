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
		return
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.FieldsFunc(line, func(r rune) bool {
			return !unicode.IsLetter(r)
		})

		for _, word := range words {
			word = strings.ToLower(word)
			if word != "" {
				wordCount[word]++
			}
		}
	}

	type wordFreq struct {
		word  string
		count int
	}

	var words []wordFreq
	for word, count := range wordCount {
		words = append(words, wordFreq{word, count})
	}

	sort.Sort(sort.Reverse(&wordFreqSlice{words}))

	for _, wf := range words {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}

type wordFreqSlice struct {
	freqs []wordFreq
}

func (w *wordFreqSlice) Len() int {
	return len(w.freqs)
}

func (w *wordFreqSlice) Less(i, j int) bool {
	if w.freqs[i].count == w.freqs[j].count {
		return w.freqs[i].word < w.freqs[j].word
	}
	return w.freqs[i].count < w.freqs[j].count
}

func (w *wordFreqSlice) Swap(i, j int) {
	w.freqs[i], w.freqs[j] = w.freqs[j], w.freqs[i]
}