package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name  string
	Birth string `json:"birthdate"`
	Age   int    `json:"age"`
}

func main() {
	// Read CSV file
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading CSV:", err)
		os.Exit(1)
	}

	var persons []Person

	// Skip header row if it exists
	if len(records) > 0 {
		reader.Reset()
	}

	for _, record := range records {
		if len(record) < 2 {
			continue
		}

		name := record[0]
		birthDate := strings.TrimSpace(record[1])

		// Calculate age as of July 1, 2025
		referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

		// Parse birth date - try multiple formats
		var birthTime time.Time
		parsedDate, err := time.Parse("2006-01-02", birthDate)
		if err != nil {
			birthTime = parsedDate
		} else {
			birthTime = parsedDate
		}

		age := calculateAge(birthTime, referenceDate)

		person := Person{
			Name:  name,
			Birth: birthDate,
			Age:   age,
		}

		persons = append(persons, person)
	}

	// Sort by name
	sort.Slice(persons, func(i, j int) bool {
		return strings.ToLower(persons[i].Name) < strings.ToLower(persons[j].Name)
	})

	// Output JSON array
	output, err := json.Marshal(persons)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func calculateAge(birthTime time.Time, referenceDate time.Time) int {
	// Ensure both dates are in UTC
	birthTime = birthTime.UTC()
	referenceDate = referenceDate.UTC()

	// Calculate age
	age := int(referenceDate.Sub(birthTime).Hours() / 24.0)

	// Adjust if birthday has passed in the reference year
	if birthTime.After(referenceDate.AddDate(0, 0, -1)) {
		age--
	}

	return age
}