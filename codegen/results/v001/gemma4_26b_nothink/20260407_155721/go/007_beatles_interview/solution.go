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

// TargetDate represents the reference date for age calculation
var TargetDate = time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

// Person represents the structure inferred from the transformation rules
type Person struct {
	Name string `json:"name"`
	Age  int    `json    "age"`
}

func main() {
	// 1. Open the input CSV file
	csvFile, err := os.Open("input/input.csv")
	if err != nil {
		log.Fatalf("failed to open csv: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	// 2. Read header to map columns
	header, err := reader.Read()
	if err != nil {
		log.Fatalf("failed to read header: %v", err)
	}

	// Map column names to indices
	colMap := make(map[string]int)
	for i, name := range header {
		colMap[strings.ToLower(strings.TrimSpace(name))] = i
	}

	var results []Person

	// 3. Process rows
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error reading row: %v", err)
		}

		// We assume the CSV has 'name' and 'birth_date' columns 
		// based on the requirement to calculate age.
		nameIdx, hasName := colMap["name"]
		dobIdx, hasDob := colMap["birth_date"]

		if !hasName || !hasDob {
			// If column names differ, this logic would be adjusted 
			// to match the specific mapping in your expected_format.json
			log.Fatalf("CSV must contain 'name' and 'birth_date' columns")
		}

		name := record[nameIdx]
		dobStr := record[dobIdx]

		// Parse Date (supporting common ISO format YYYY-MM-DD)
		dob, err := time.Parse("2006-01-02", strings.TrimSpace(dobStr))
		if err != nil {
			continue // Skip rows with invalid dates
		}

		// 4. Calculate Age as of July 1, 2025
		age := calculateAge(dob, TargetDate)

		results = append(results, Person{
			Name: name,
			Age:  age,
		})
	}

	// 5. Output to stdout as JSON array
	encoder := json.NewEncoder(os.Stdout)
	// Set indent for readability if needed, though standard JSON output is fine
	// encoder.SetIndent("", "  ") 
	if err := encoder.Encode(results); err != nil {
		log.Fatalf("failed to encode json: %v", err)
	}
}

// calculateAge computes the age in years between birth date and target date
func calculateAge(birthDate, targetDate time.Time) int {
	years := targetDate.Year() - birthDate.Year()

	// Adjust if the birthday hasn't occurred yet in the target year
	if targetDate.Month() < birthDate.Month() || 
	   (targetDate.Month() == birthDate.Month() && targetDate.Day() < birthDate.Day()) {
		years--
	}
	return years
}