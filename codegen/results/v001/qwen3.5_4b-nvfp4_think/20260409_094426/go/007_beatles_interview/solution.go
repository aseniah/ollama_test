package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Person struct {
	Name     string        `json:"name"`
	DateOfBirth string     `json:"date_of_birth"`
	Age      int           `json:"age"`
}

func main() {
	// Read input.csv and process to calculate ages as of July 1, 2025
	inputPath := filepath.Join("input", "input.csv")
	expectedFormat := filepath.Join("input", "expected_format.json")

	// Open the CSV file
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var records []Person

	// Read all rows from the CSV
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// Parse records (skip header row at index 0)
	for i := 1; i < len(rows); i++ {
		row, err := reader.Read()
		if err != nil {
			continue // skip invalid rows
		}
		
		if len(row) >= 2 {
			record := parsePersonRow(row)
			records = append(records, record)
		}
	}

	// Reference expected format (for output validation)
	expectedFormatBytes, err := os.ReadFile(expectedFormat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading expected format: %v\n", err)
		// Continue anyway to process records
	}

	// Output JSON array to stdout
	output, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func parsePersonRow(row []string) Person {
	// Expected CSV columns: name, date_of_birth (YYYY-MM-DD)
	name := row[0]
	dateOfBirth := row[1]

	// Calculate age as of July 1, 2025
	now := time.Date(2025, 7, 1, 0, 0, 0, 0, time.Local)
	doB, err := time.Parse("2006-01-02", dateOfBirth)
	if err != nil {
		// Use a reasonable fallback if parsing fails
		nameAge := Person{
			Name:     name,
			DateOfBirth: dateOfBirth,
			Age:      0,
		}
		return nameAge
	}

	yearsAgo := now.Sub(doB).YearsSince()
	age := int(yearsAgo)

	return Person{
		Name:     name,
		DateOfBirth: dateOfBirth,
		Age:      age,
	}
}