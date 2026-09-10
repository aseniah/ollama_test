package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// User represents the structure of each row in the CSV
type User struct {
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

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// If file is empty or only has header, return empty array
	if len(records) <= 1 {
		fmt.Println("[]")
		return
	}

	var users []User

	// Iterate through records skipping the header (index 0)
	for i := 1; i < len(records); i++ {
		row := records[i]
		
		// Basic validation to ensure row has 4 columns
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

		users = append(users, User{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	// Marshal the slice to JSON
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Output to stdout
	fmt.Println(string(jsonData))
}