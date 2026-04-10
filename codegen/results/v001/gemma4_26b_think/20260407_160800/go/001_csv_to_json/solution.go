package main

import (
	"encoding/csv"
	"encoding/json"
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
	// Open the target CSV file
	file, err := os.Open("input/data.csv")
	if err != nil {
		// If the file cannot be opened, we exit without outputting to stdout
		// to avoid violating the "Output must be valid JSON" constraint.
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	// If there are no records or only a header, output an empty array
	if len(records) <= 1 {
		json.NewEncoder(os.Stdout).Encode([]Person{})
		return
	}

	// Skip the header (index 0) and process data rows
	dataRows := records[1:]
	var people []Person

	for _, row := range dataRows {
		// Ensure the row has enough columns to prevent index out of range
		if len(row) < 4 {
			continue
		}

		// Parse Age (string to int)
		age, err := strconv.Atoi(row[1])
		if err != nil {
			age = 0
		}

		// Parse Score (string to float64)
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			score = 0.0
		}

		// Construct the Person object
		person := Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}

		people = append(people, person)
	}

	// Encode the slice of Person objects directly to stdout as a JSON array
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(people)
	if err != nil {
		return
	}
}