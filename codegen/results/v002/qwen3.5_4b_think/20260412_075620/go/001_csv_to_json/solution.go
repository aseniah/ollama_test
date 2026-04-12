package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Email string `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var people []Person
	seenHeader := false

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
			os.Exit(1)
		}

		if len(record) < 4 {
			continue
		}

		if !seenHeader {
			seenHeader = true
			continue
		}

		name := record[0]
		age, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing age for %s: %v\n", name, err)
			os.Exit(1)
		}
		email := record[2]
		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing score for %s: %v\n", name, err)
			os.Exit(1)
		}

		people = append(people, Person{Name: name, Age: age, Email: email, Score: score})
	}

	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}