package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Email string `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	// Open the CSV file
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Skip the header row
	var people []Person
	for i, record := range records {
		if i == 0 {
			continue // Skip the header
		}

		// Parse each field and create a Person object
		person := Person{
			Name:  record[0],
			Age:   parseInt(record[1]),
			Email: record[2],
			Score: parseFloat(record[3]),
		}
		people = append(people, person)
	}

	// Marshal the people slice to JSON
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	// Output the JSON data to stdout
	fmt.Println(string(jsonData))
}

func parseInt(value string) int {
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Error parsing integer value:", err)
		return 0
	}
	return parsedValue
}

func parseFloat(value string) float64 {
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Error parsing float value:", err)
		return 0.0
	}
	return parsedValue
}