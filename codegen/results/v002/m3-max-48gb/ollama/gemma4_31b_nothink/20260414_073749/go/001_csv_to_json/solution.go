package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Person represents the structure of each row in the CSV
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

	// Read the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// If the file is empty or only contains the header, return an empty array
	if len(records) <= 1 {
		fmt.Println("[]")
		return
	}

	var people []Person

	// Iterate through records, skipping the header (index 0)
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 4 {
			continue
		}

		// Parse Age as integer
		age, err := strconv.Atoi(row[1])
		if err != nil {
			age = 0
		}

		// Parse Score as float
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			score = 0.0
		}

		people = append(people, Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	// Marshal the slice to JSON
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}

	// Output the JSON to stdout
	fmt.Println(string(jsonData))
}