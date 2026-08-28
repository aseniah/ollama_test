package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

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

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// Skip header (first row)
	if len(records) == 0 {
		fmt.Println("[]")
		return
	}

	records = records[1:]

	// Parse records into persons
	var persons []Person
	for _, record := range records {
		if len(record) < 4 {
			continue
		}

		name := record[0]
		age, _ := strconv.Atoi(record[1])
		email := record[2]
		score, _ := strconv.ParseFloat(record[3], 64)

		person := Person{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		}
		persons = append(persons, person)
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(persons)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Output to stdout
	fmt.Println(string(jsonData))
}