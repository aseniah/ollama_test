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
		return
	}
	defer file.Close()

	// Initialize CSV reader
	reader := csv.NewReader(file)
	
	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	// Slice to hold the parsed Person objects
	var people []Person

	// Iterate through records, skipping the header (index 0)
	for i, record := range records {
		if i == 0 {
			continue
		}

		// Parse Age to integer
		age, err := strconv.Atoi(record[1])
		if err != nil {
			continue
		}

		// Parse Score to float64
		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			continue
		}

		// Create Person object and add to slice
		people = append(people, Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		})
	}

	// Convert the slice of Person objects to JSON
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}

	// Print the JSON output to stdout
	fmt.Println(string(jsonData))
}