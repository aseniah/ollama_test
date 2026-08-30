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
	// Open the file
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		os.Exit(1)
	}

	// Skip header (first row)
	if len(records) < 2 {
		// Output empty JSON array if no data rows
		fmt.Println("[]")
		return
	}

	var people []Person

	// Iterate over data rows (starting from index 1)
	for _, row := range records[1:] {
		if len(row) < 4 {
			continue // Skip malformed rows
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Println("Error parsing Age:", err)
			continue
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Println("Error parsing Score:", err)
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

	// Marshal to JSON with indentation
	jsonOutput, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		os.Exit(1)
	}

	// Write to stdout
	fmt.Println(string(jsonOutput))
}