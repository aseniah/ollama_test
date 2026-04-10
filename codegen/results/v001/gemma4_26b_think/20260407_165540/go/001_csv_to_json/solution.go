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
	// Open the input file
	file, err := os.Open("input/data.csv")
	if err != nil {
		// If the file cannot be opened, we exit. 
		// We do not print anything to stdout to comply with "Do not output anything else".
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

	var people []Person

	// Iterate through records, skipping the header (index 0)
	for i := 1; i < len(records); i++ {
		row := records[i]
		
		// Ensure the row has enough columns to prevent index out of bounds
		if len(row) < 4 {
			continue
		}

		// Parse Age as integer
		age, err := strconv.Atoi(row[1])
		if err != nil {
			age = 0
		}

		// Parse Score as float64
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			score = 0.0
		}

		// Create Person object
		person := Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}

		people = append(people, person)
	}

	// Encode the slice of Person objects directly to stdout as JSON
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(people)
	if err != nil {
		return
	}
}