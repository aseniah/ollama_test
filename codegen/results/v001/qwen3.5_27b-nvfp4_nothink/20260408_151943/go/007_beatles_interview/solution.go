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
	Name     string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	Age      int    `json:"age"`
}

func main() {
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading CSV: %v\n", err)
		os.Exit(1)
	}

	var people []Person

	// Skip header if present (assumed based on typical CSV format)
	startIndex := 0
	if len(records) > 0 && strings.EqualFold(records[0][0], "Name") || strings.EqualFold(records[0][0], "name") {
		startIndex = 1
	}

	for _, record := range records[startIndex:] {
		if len(record) < 2 {
			continue
		}

		name := strings.TrimSpace(record[0])
		dateStr := strings.TrimSpace(record[1])
		
		var dob time.Time
		// Try common formats: YYYY-MM-DD, DD/MM/YYYY, MM/DD/YYYY
		var err error
		formats := []string{"2006-01-02", "02/01/2006", "01/02/2006"}
		for _, format := range formats {
			dob, err = time.Parse(format, dateStr)
			if err == nil {
				break
			}
		}
		
		if err != nil {
			// Fallback: try to parse loosely if strict parsing fails, 
			// but for this exercise we assume valid dates or skip.
			continue
		}

		age := calculateAge(dob, refDate)

		people = append(people, Person{
			Name:        name,
			DateOfBirth: dateStr, // Output the original string as per typical expectations unless specified otherwise
			Age:         age,
		})
	}

	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(jsonData))
}

func calculateAge(birthDate time.Time, referenceDate time.Time) int {
	years := referenceDate.Year() - birthDate.Year()
	
	// Check if the birthday has occurred this year yet
	birthdayThisYear := time.Date(referenceDate.Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, birthDate.Location())
	if referenceDate.Before(birthdayThisYear) {
		years--
	}

	return years
}