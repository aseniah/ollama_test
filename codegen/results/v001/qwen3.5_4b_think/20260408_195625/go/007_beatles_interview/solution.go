package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

// Person represents a person from the CSV data
type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Birthdate *time.Time `json:"birthdate"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
}

// CalculateAge calculates the age as of July 1, 2025
func CalculateAge(birthTime time.Time, referenceTime time.Time) int {
	years := referenceTime.Year() - birthTime.Year()
	daysInYear := 365
	if referenceTime.Year() == birthTime.Year() && daysInYear != 365 {
		daysInYear = 366
	}
	age := years * daysInYear - birthTime.Year() + 1
	if birthTime.Year() != referenceTime.Year() {
		age = birthTime.Year() - referenceTime.Year() + 1
	} else if birthTime.Month() < referenceTime.Month() || (birthTime.Month() == referenceTime.Month() && birthTime.Day() < referenceTime.Day()) {
		age = birthTime.Year() - referenceTime.Year() - 1
	}
	return age
}

// ParseAge parses an age from birthdate to reference date
func ParseAge(birthTime time.Time, referenceTime time.Time) int {
	years := referenceTime.Year() - birthTime.Year()
	if birthTime.Year() == referenceTime.Year() && birthTime.Month() >= referenceTime.Month() {
		return years
	}
	return years - 1
}

// Person represents a person from the CSV data
type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Birthdate *time.Time `json:"birthdate"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
}

// calculateAge calculates the age as of July 1, 2025
func calculateAge(birthTime time.Time) int {
	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	years := referenceDate.Year() - birthTime.Year()
	if years < 0 {
		return -1
	}
	if birthTime.Year() == referenceDate.Year() {
		if birthTime.Month() > referenceDate.Month() || (birthTime.Month() == referenceDate.Month() && birthTime.Day() > referenceDate.Day()) {
			years--
		}
	}
	return years
}

// parseCSV reads and parses the input CSV file
func parseCSV(filename string) ([]Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var persons []Person

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(records) == 0 {
		return persons, nil
	}

	// Assuming CSV has headers: ID, Name, DOB, Gender
	// Adjust headers as needed based on your actual input
	// Example: ID, Name, BirthDate, Gender
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		if len(record) < 2 {
			continue
		}

		var person Person

		// Parse ID (assumed to be numeric, may not be needed)
		id := i + 1
		person.ID = id

		// Parse Name
		name := record[1]
		person.Name = name

		// Parse Birthdate
		birthStr := record[2]
		if len(birthStr) == 0 {
			continue
		}

		birthTime, err := time.Parse("2006-01-02", birthStr)
		if err != nil {
			fmt.Printf("Skipping record %d: invalid birthdate format: %s\n", i+1, birthStr)
			continue
		}
		person.Birthdate = &birthTime

		// Parse Age (as of July 1, 2025)
		person.Age = calculateAge(*birthTime)

		// Parse Gender (optional, may not exist in some files)
		if len(record) > 3 {
			person.Gender = record[3]
		}

		persons = append(persons, person)
	}

	return persons, nil
}

// main reads the CSV file and outputs the JSON array
func main() {
	persons, err := parseCSV("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing CSV: %v\n", err)
		os.Exit(1)
	}

	// Output to stdout as JSON array
	output, err := json.Marshal(persons)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}