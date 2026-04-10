package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// Person represents the structure of the output JSON objects.
// You may need to adjust field names based on your specific expected_format.json.
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Reference date for age calculation: July 1, 2025
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	// Open the input CSV file
	filePath := "input/input.csv"
	file, err := os.Open(filePath)
	if err != nil {
		// If file doesn't exist in this environment, we exit silently 
		// or handle as per requirement.
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read header row to identify columns
	header, err := reader.Read()
	if err != nil {
		return
	}

	// Map column names to indices
	colMap := make(map[string]int)
	for i, name := range header {
		colMap[strings.TrimSpace(name)] = i
	}

	var results []Person

	// Iterate through CSV rows
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		// Extract values using the map
		// Assuming columns are named "name" and "birth_date"
		nameIdx, nameExists := colMap["name"]
		dobIdx, dobExists := colMap["birth_date"]

		if !nameExists || !dobExists {
			break
		}

		name := record[nameIdx]
		dobStr := record[dobIdx]

		// Parse date (assuming format YYYY-MM-DD, adjust if your CSV differs)
		dob, err := time.Parse("2006-01-02", strings.TrimSpace(dobStr))
		if err != nil {
			continue // Skip rows with invalid dates
		}

		// Calculate age
		age := calculateAge(dob, refDate)

		results = append(results, Person{
			Name: name,
			Age:  age,
		})
	}

	// Encode the slice to JSON and write to stdout
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(results); err != nil {
		log.Fatal(err)
	}
}

// calculateAge computes the age based on a birth date and a reference date.
func calculateAge(birthDate, refDate time.Time) int {
	years := refDate.Year() - birthDate.Year()

	// Adjust if the birthday hasn't occurred yet in the reference year
	if refDate.Month() < birthDate.Month() || 
	   (refDate.Month() == birth:// birthDate.Month() && refDate.Day() < birthDate.Day()) {
		years--
	}

	return years
}