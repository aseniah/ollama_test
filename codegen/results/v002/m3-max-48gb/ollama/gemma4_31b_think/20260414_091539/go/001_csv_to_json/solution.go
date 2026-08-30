package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Person represents the structure of the data in the CSV
type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	// Open the input file
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// If the file is empty or only contains a header, return an empty array
	if len(records) <= 1 {
		fmt.Println("[]")
		return
	}

	var people []Person

	// Iterate through records starting from index 1 to skip the header
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 4 {
			continue
		}

		// Parse Age as integer
		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing age on row %d: %v\n", i, err)
			continue
		}

		// Parse Score as float
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing score on row %d: %v\n", i, err)
			continue
		}

		// Create Person object and append to slice
		person := Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}
		people = append(people, person)
	}

	// Marshal the slice to JSON
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling to JSON: %v\n", err)
		os.Exit(1)
	}

	// Output the final JSON array to stdout
	fmt.Println(string(jsonData))
}