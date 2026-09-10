package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Person represents the structure of the data required in the JSON output.
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

	// Check if the file is empty or only contains a header
	if len(records) <= 1 {
		fmt.Println("[]")
		return
	}

	var people []Person

	// Iterate through records, skipping the header row (index 0)
	for i := 1; i < len(records); i++ {
		row := records[i]
		
		// Ensure the row has the expected number of columns
		if len(row) < 4 {
			continue
		}

		// Parse Age as integer
		age, err := strconv.Atoi(row[1])
		if err != nil {
			continue
		}

		// Parse Score as float64
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
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

	// Marshal the slice of people into JSON
	jsonData, err := json.Marshal(people)
	if err != nil {
		return
	}

	// Output the JSON to stdout
	fmt.Println(string(jsonData))
}