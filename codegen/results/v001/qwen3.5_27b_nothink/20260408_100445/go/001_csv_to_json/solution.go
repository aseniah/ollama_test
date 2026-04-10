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
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// Skip header (first row)
	if len(records) < 2 {
		// If no data rows, output empty JSON array
		jsonBytes, _ := json.Marshal([]Person{})
		fmt.Println(string(jsonBytes))
		return
	}

	var people []Person

	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 4 {
			// Skip malformed rows or handle as needed
			continue
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			// Handle invalid age gracefully (skip or default)
			continue
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			// Handle invalid score gracefully
			continue
		}

		person := Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}
		people = append(people, person)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(people)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}