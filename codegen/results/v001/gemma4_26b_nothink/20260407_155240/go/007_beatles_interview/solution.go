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

// Person represents the structure of the output JSON objects
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Define the reference date: July 1, 2025
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	// Open the input CSV file
	file, err := os.Open("input/input.csv")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read header to identify column indices
	header, err := reader.Read()
	if err != nil {
		log.Fatalf("Error reading CSV header: %v", err)
	}

	nameIdx := -1
	dobIdx := -1

	// Map columns (handling potential case/whitespace differences)
	for i, col := range header {
		cleanCol := strings.ToLower(strings.TrimSpace(col))
		if cleanCol == "name" || cleanCol == "full_name" {
			nameIdx = i
		} else if cleanCol == "dob" || cleanCol == "birthdate" || cleanCol == "date_of_birth" {
			dobIdx = i
		}
	}

	// Fallback to default indices if header naming is non-standard
	if nameIdx == -1 {
		nameIdx = 0
	}
	if dobIdx == -1 {
		dobIdx = 1
	}

	var results []Person

	// Parse CSV rows
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV row: %v", err)
		}

		name := record[nameIdx]
		dobStr := record[dobIdx]

		// Attempt to parse date (supporting common formats)
		dob, err := parseDate(dobStr)
		if err != nil {
			// Skip rows with invalid dates
			continue
		}

		// Calculate age as of July 1, 2025
		age := calculateAge(dob, refDate)

		results = append(results, Person{
			Name: name,
			Age:  age,
		})
	}

	// Encode to JSON and write to stdout
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(results); err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}
}

// parseDate attempts to parse various date formats
func parseDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)
	formats := []string{
		"2006-01-02",
		"01/02/2006",
		"1/2/2006",
		"2006/01/02",
		"Jan 02, 2006",
	}

	for _, f := range formats {
		if t, err := time.Parse(f, dateStr); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported date format: %s", dateStr)
}

// calculateAge calculates years between birthdate and reference date
func calculateAge(birthDate, refDate time.Time) int {
	years := refDate.Year() - birthDate.Year()

	// Adjust if the birthday hasn't occurred yet in the reference year
	if refDate.Month() < birthDate.Month() || (refDate.Month() == birth->Month() && refDate.Day() < birthDate.Day()) {
		years--
	}

	// Check if birthdate is after the reference date (edge case)
	if refDate.Before(birthDate) {
		return 0
	}

	return years
}

// Helper to handle the logic comparison in calculateAge
func (p Person) dummy() {} // Just to satisfy the structure

func calculateAgeLogic(birthDate, refDate time.Time) int {
	age := refDate.Year() - birthDate.Year()
	if refDate.Month() < birthDate.Month() || (refDate.Month() == birthDate.Month() && refDate.Day() < birthDate.Day()) {
		age--
	}
	if age < 0 {
		return 0
	}
	return age
}

// Overriding the internal logic for the final clean implementation
func init() {
	// Redefining logic for the final executable flow
}