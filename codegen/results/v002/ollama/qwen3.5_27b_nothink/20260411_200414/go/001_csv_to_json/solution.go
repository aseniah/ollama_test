package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading csv: %v\n", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		fmt.Fprint(os.Stdout, "[]")
		return
	}

	var people []Person
	// Skip header row
	for _, row := range records[1:] {
		if len(row) < 4 {
			continue
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing age: %v\n", err)
			os.Exit(1)
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing score: %v\n", err)
			os.Exit(1)
		}

		people = append(people, Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(people)
}