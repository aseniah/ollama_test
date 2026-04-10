package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name  string `json:"name"`
	Birth int    `json:"birth_year"`
}

func main() {
	// Open the CSV file
	csvFile, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer csvFile.Close()

	// Read the CSV data
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Process each record to calculate age and format as JSON
	var people []Person

	for i, record := range records {
		if i == 0 { // Skip the header row
			continue
		}

		name := strings.TrimSpace(record[0])
		birthYear, err := strconv.Atoi(strings.TrimSpace(record[1]))
		if err != nil {
			fmt.Println("Error parsing birth year:", err)
			return
		}

		today := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
		age := today.Year() - birthYear

		people = append(people, Person{
			Name:  name,
			Birth: birthYear,
		})
	}

	// Marshal the people slice to JSON
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))
}