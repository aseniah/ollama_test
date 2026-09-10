package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Person represents the data structure for each row in the CSV
type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	// Open the CSV file
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Initialize CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// If the file is empty or only has a header, output an empty array
	if len(records) <= 1 {
		fmt.Println("[]")
		return
	}

	var people []Person

	// Iterate through records, skipping the header (index 0)
	for i := 1; i < len(records); i++ {
		row := records[i]
		
		// Parse Age to integer
		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing age for %s: %v\n", row[0], err)
			continue
		}

		// Parse Score to float64
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing score for %s: %v\n", row[0], err)
			continue
		}

		// Create Person object and add to slice
		people = append(people, Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	// Encode the slice of Person objects to JSON and write to stdout
	encoder := json.NewEncoder(os.Stdout)
	// SetIndent is optional, but for standard JSON output, we just Encode
	if err := encoder.Encode(people); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}