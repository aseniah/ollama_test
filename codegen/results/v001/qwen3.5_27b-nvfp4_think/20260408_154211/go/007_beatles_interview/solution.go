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
	
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	calculateDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)
	
	var people []Person
	rowNum := 1
	
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading record %d: %v\n", rowNum, err)
			os.Exit(1)
		}

		person := Person{}

		for i, name := range header {
			switch strings.ToLower(strings.TrimSpace(name)) {
			case "name":
				if i < len(record) {
					person.Name = strings.TrimSpace(record[i])
				}
			case "date_of_birth", "dob", "birth_date", "birthday":
				if i < len(record) && record[i] != "" {
					date, err := parseDate(record[i])
					if err == nil {
						age := calculateAge(date, calculateDate)
						person.Age = age
					}
				}
			case "email":
			case "city":
			default:
			}
		}

		if person.Name != "" {
			people = append(people, person)
		}

		rowNum++
	}

	jsonBytes, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonBytes))
}

func parseDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)

	formats := []string{
		"2006-01-02",
		"01/02/2006",
		"02/01/2006",
		"Jan 02, 2006",
		"January 2, 2006",
		"02-Jan-2006",
		"2 Jan 2006",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t, nil
		}
	}

	ageYears, err := strconv.Atoi(dateStr)
	if err == nil && ageYears > 0 && ageYears < 120 {
		birthYear := 2025 - ageYears
		return time.Date(birthYear, time.July, 1, 0, 0, 0, 0, time.UTC), nil
	}

	return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}

func calculateAge(birthDate time.Time, referenceDate time.Time) int {
	age := referenceDate.Year() - birthDate.Year()
	if referenceDate.Month() < birthDate.Month() ||
		(referenceDate.Month() == birthDate.Month() && referenceDate.Day() < birthDate.Day()) {
		age--
	}
	return age
}