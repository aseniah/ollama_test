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

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	// Check if file has data beyond the header
	if len(records) <= 1 {
		fmt.Println("[]")
		return
	}

	var people []Person

	// Process records skipping the header (index 0)
	for i := 1; i < len(records); i++ {
		row := records[i]
		
		// Ensure row has the expected number of columns
		if len(row) < 4 {
			continue
		}

		// Parse Age to integer
		age, _ := strconv.Atoi(row[1])
		
		// Parse Score to float
		score, _ := strconv.ParseFloat(row[3], 64)

		person := Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}
		people = append(people, person)
	}

	// Convert the slice to JSON
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}

	// Output valid JSON to stdout
	fmt.Println(string(jsonData))
}