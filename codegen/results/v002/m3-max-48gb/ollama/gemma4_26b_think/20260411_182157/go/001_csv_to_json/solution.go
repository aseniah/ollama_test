package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Person represents the structure of each row in the CSV.
// Using the exact field names provided in the requirements.
type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	// Open the input file
	filePath := "input/data.csv"
	file, err := os.Open(filePath)
	if err != nil {
		// If the file cannot be opened, we exit without printing anything to avoid
		// polluting the stdout with non-JSON error messages.
		return
	}
	defer file.Close()

	// Initialize CSV reader
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	// Ensure there is data to process (at least header + 1 row)
	if len(records) <= 1 {
		fmt.Print("[]")
		return
	}

	var people []Person

	// Iterate through records, skipping the first row (the header)
	for i := 1; i < len(records); i++ {
		row := records[i]

		// Ensure the row has the expected number of columns
		if len(row) < 4 {
			continue
		}

		// Parse Age (integer)
		age, err := strconv.Atoi(row[1])
		if err != nil {
			age = 0 // Fallback for invalid integers
		}

		// Parse Score (float)
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			score = 0.0 // Fallback for invalid floats
		}

		// Create Person object and add to slice
		person := Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}
		people = append(people, person)
	}

	// Convert the slice of Person objects to a JSON array
	// Using MarshalIndent to provide a clean, readable JSON output
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}

	// Output the valid JSON to stdout
	fmt.Print(string(jsonData))
}