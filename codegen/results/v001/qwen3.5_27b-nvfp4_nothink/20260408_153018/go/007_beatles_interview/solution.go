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
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func calculateAge(birthDateString string, referenceDate time.Time) (int, error) {
	// Try common formats: YYYY-MM-DD, MM/DD/YYYY, DD/MM/YYYY, etc.
	var birthDate time.Time
	formats := []string{
		"2006-01-02", // ISO
		"01/02/2006", // US/Canada
		"02/01/2006", // EU/Aus
		"2006/01/02", // Slash separated year first
		"02-01-2006", // Dash separated day first (ambiguous without locale, assuming EU style)
	}

	for _, format := range formats {
		parsed, err := time.Parse(format, strings.TrimSpace(birthDateString))
		if err == nil {
			birthDate = parsed
			break
		}
	}

	if birthDate.IsZero() {
		return 0, fmt.Errorf("unable to parse date: %s", birthDateString)
	}

	age := referenceDate.Year() - birthDate.Year()
	if referenceDate.YearDay() < birthDate.YearDay() {
		age--
	}
	return age, nil
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		fmt.Fprintf(os.Stderr, "Invalid or empty CSV file")
		os.Exit(1)
	}

	// Determine column indices based on header (assuming case-insensitive match for Name and DOB/BirthDate)
	headers := records[0]
	nameIdx := -1
	dobIdx := -1

	for i, h := range headers {
		hClean := strings.ToLower(strings.TrimSpace(h))
		if hClean == "name" || hClean == "full_name" {
			nameIdx = i
		} else if hClean == "date_of_birth" || hClean == "dob" || hClean == "birth_date" || hClean == "birthdate" {
			dobIdx = i
		}
	}

	if nameIdx == -1 || dobIdx == -1 {
		fmt.Fprintf(os.Stderr, "CSV must contain 'name' and 'date_of_birth' (or similar) columns")
		os.Exit(1)
	}

	var people []Person
	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	// Process data rows
	for i := 1; i < len(records); i++ {
		row := records[i]
		if nameIdx >= len(row) || dobIdx >= len(row) {
			continue
		}

		name := strings.TrimSpace(row[nameIdx])
		dobStr := strings.TrimSpace(row[dobIdx])

		if name == "" {
			continue
		}

		age, err := calculateAge(dobStr, referenceDate)
		if err != nil {
			// If date parsing fails, we might skip or set -1. 
			// Given the prompt implies a transformation, skipping invalid rows is safer for clean JSON.
			continue
		}

		people = append(people, Person{
			Name: name,
			Age:  age,
		})
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
	
	// Suppress the strconv unused error if it occurs in different Go versions/contexts by using a dummy usage
	_ = strconv.Itoa(len(records))
}