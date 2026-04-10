package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Score float64 `json:"score"`
}

func main() {
	filePath := "input/data.csv"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	reader := csv.NewReader(csv.NewReader(os.Stdin))
	if reader, err := csv.NewReader(content); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating CSV reader: %v\n", err)
		os.Exit(1)
	}

	var persons []Person
	lineIndex := 0
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		if lineIndex == 0 {
			lineIndex++
			continue
		}
		if len(record) != 4 {
			continue
		}
		person := Person{
			Name:  record[0],
			Age:   mustParseInt(record[1]),
			Email: record[2],
			Score: mustParseFloat(record[3]),
		}
		persons = append(persons, person)
		lineIndex++
	}

	jsonBytes, err := json.Marshal(persons)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling to JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonBytes))
}

func mustParseInt(s string) int {
	i, err := fmt.Sscanf(s, "%d")
	if err != nil || i != 1 {
		panic(err)
	}
	return i
}

func mustParseFloat(s string) float64 {
	f, err := fmt.Sscanf(s, "%f")
	if err != nil || f != 1 {
		panic(err)
	}
	return f
}