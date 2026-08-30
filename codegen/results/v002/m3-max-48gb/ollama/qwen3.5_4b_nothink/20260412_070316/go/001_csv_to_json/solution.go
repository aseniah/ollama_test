package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Email string
	Score float64
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	var persons []Person
	for i, record := range records {
		// Skip header row
		if i == 0 {
			continue
		}

		p := Person{
			Name:  record[0],
			Age:   0,
			Email: record[2],
			Score: 0,
		}

		if len(record) >= 3 {
			age, err := fmt.Sscanf(record[1], "%d", &p.Age)
			if err != nil || age != 1 {
				fmt.Fprintf(os.Stderr, "Error parsing age for %s: %v\n", record[0], err)
				os.Exit(1)
			}
		}

		if len(record) >= 4 {
			score, err := fmt.Sscanf(record[3], "%f", &p.Score)
			if err != nil || score != 1 {
				fmt.Fprintf(os.Stderr, "Error parsing score for %s: %v\n", record[0], err)
				os.Exit(1)
			}
		}

		persons = append(persons, p)
	}

	output, err := json.MarshalIndent(persons, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}