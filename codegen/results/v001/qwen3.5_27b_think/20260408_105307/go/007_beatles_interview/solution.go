package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) == 0 {
		fmt.Println("[]")
		return
	}

	// Detect column positions
	header := records[0]
	nameIdx, dateIdx := -1, -1
	for i, col := range header {
		colLower := strings.ToLower(strings.TrimSpace(col))
		if colLower == "name" || colLower == "first_name" || colLower == "full_name" || colLower == "person" {
			nameIdx = i
		}
		if colLower == "date_of_birth" || colLower == "birthdate" || colLower == "dob" || colLower == "birth_date" || colLower == "birth" {
			dateIdx = i
		}
	}

	// If not found, assume default positions
	if nameIdx == -1 {
		nameIdx = 0
	}
	if dateIdx == -1 {
		dateIdx = 1
	}

	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	// Process data rows (skip header)
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) <= nameIdx || len(row) <= dateIdx {
			continue
		}

		name := strings.TrimSpace(row[nameIdx])
		dateStr := strings.TrimSpace(row[dateIdx])

		if name == "" {
			continue
		}

		// Parse birth date - try multiple formats
		var birthDate time.Time
		var parsed bool

		formats := []string{
			"2006-01-02",
			"2006/01/02",
			"01/02/2006",
			"02/01/2006",
			"January 2, 2006",
			"Jan 2, 2006",
			"2 January 2006",
			"2 Jan 2006",
			"2006-1-2",
			"2006-01-2",
			"2-1-2006",
		}

		for _, format := range formats {
			birthDate, parsed = time.Parse(format, dateStr)
			if parsed {
				break
			}
		}

		if !parsed {
			continue
		}

		// Calculate age as of July 1, 2025
		age := referenceDate.Year() - birthDate.Year()
		if referenceDate.YearDay() < birthDate.YearDay() ||
			(referenceDate.YearDay() == birthDate.YearDay() &&
				(referenceDate.Month() < birthDate.Month() ||
					(referenceDate.Month() == birthDate.Month() && referenceDate.Day() < birthDate.Day()))) {
			age--
		}

		people = append(people, Person{
			Name: name,
			Age:  age,
		})
	}

	// Output JSON to stdout
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(people)
}