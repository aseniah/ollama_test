package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"strings"
	"time"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	DOB   string `json:"dob"`
	Email string `json:"email,omitempty"`
}

func calculateAge(birthDate time.Time, referenceDate time.Time) int {
	years := referenceDate.Year() - birthDate.Year()
	if referenceDate.YearDay() < birthDate.YearDay() ||
		(referenceDate.YearDay() == birthDate.YearDay() && referenceDate.Month() < birthDate.Month()) {
		years--
	}
	return years
}

func parseDate(dateStr string) time.Time {
	formats := []string{
		"2006-01-02",
		"2006/01/02",
		"01/02/2006",
		"02/01/2006",
		"2006-01-02T00:00:00Z",
		"2006-01-02T15:04:05",
	}
	for _, format := range formats {
		if t, err := time.Parse(format, strings.TrimSpace(dateStr)); err == nil {
			return t
		}
	}
	return time.Time{}
}

func main() {
	// Reference date: July 1, 2025
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	// Open CSV file
	file, err := os.Open("input/input.csv")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read header
	header, err := reader.Read()
	if err != nil {
		os.Exit(1)
	}

	// Normalize headers to lowercase for comparison
	for i, h := range header {
		header[i] = strings.ToLower(strings.TrimSpace(h))
	}

	// Find column indices
	nameIdx := -1
	dobIdx := -1
	emailIdx := -1

	for i, h := range header {
		switch h {
		case "name":
			nameIdx = i
		case "dob", "date_of_birth", "dateofbirth":
			dobIdx = i
		case "email":
			emailIdx = i
		}
	}

	// Validate required columns
	if nameIdx == -1 || dobIdx == -1 {
		os.Exit(1)
	}

	var people []Person

	// Read all records
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			os.Exit(1)
		}

		person := Person{}

		// Extract name
		if nameIdx < len(record) {
			person.Name = strings.TrimSpace(record[nameIdx])
		}

		// Extract and parse DOB
		if dobIdx < len(record) && record[dobIdx] != "" {
			person.DOB = strings.TrimSpace(record[dobIdx])
			birthDate := parseDate(person.DOB)
			if !birthDate.IsZero() {
				person.Age = calculateAge(birthDate, refDate)
			} else {
				person.Age = 0
			}
		}

		// Extract email if column exists
		if emailIdx >= 0 && emailIdx < len(record) && record[emailIdx] != "" {
			person.Email = strings.TrimSpace(record[emailIdx])
		}

		people = append(people, person)
	}

	// Output JSON array to stdout (compact format)
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(people); err != nil {
		os.Exit(1)
	}
}