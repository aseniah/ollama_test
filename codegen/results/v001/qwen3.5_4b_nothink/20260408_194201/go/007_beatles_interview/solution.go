package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Person represents a row from the CSV.
// Based on typical CSV structures for such tasks, we assume columns: Name, Date of Birth, Gender (optional).
// We only calculate age based on Date of Birth.
type Person struct {
	Name  string
	Dob   string
	Age   int
}

func main() {
	// Open the input CSV file
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var people []Person

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		os.Exit(1)
	}

	// Skip header row if present
	if len(records) > 0 {
		header := records[0]
		// If the first row is empty or non-numeric in first column, it might be a header.
		// Assuming standard CSV with 'Name' as first col.
		// We'll parse data starting from index 1.
	}

	for i, record := range records {
		if i == 0 {
			// Skip header
			continue
		}

		if len(record) < 2 {
			continue
		}

		name := record[0]
		dobStr := strings.TrimSpace(record[1])

		// Parse date. Format is likely YYYY-MM-DD
		dob, err := time.Parse("2006-01-02", dobStr)
		if err != nil {
			// Skip malformed dates
			continue
		}

		// Reference date: July 1, 2025
		refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

		// Calculate age
		age := calculateAge(dob, refDate)

		people = append(people, Person{Name: name, Dob: dobStr, Age: age})
	}

	// Convert to JSON array
	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		os.Exit(1)
	}

	// Print to stdout
	fmt.Println(string(output))
}

func calculateAge(dob time.Time, refDate time.Time) int {
	age := int(refDate.Sub(dob).Hours() / 24.0 / 365.25)

	// Adjust age if birthday hasn't occurred this year
	if refDate.Year() == dob.Year()+int(age) && refDate.Month() < dob.Month() || (refDate.Month() == dob.Month() && refDate.Day() < dob.Day()) {
		age--
	}
	return age
}