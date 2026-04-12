package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	Birthday  string    `json:"Birthday"`
	Age       int       `json:"Age"`
	Relatives []Relative `json:"Relatives"`
}

func parseDate(dateStr string) (time.Time, error) {
	formats := []string{
		"1/2/2006",   // M/D/YYYY
		"1/02/2006",  // M/DD/YYYY
		"01/2/2006",  // MM/D/YYYY
		"01/02/2006", // MM/DD/YYYY
	}

	for _, format := range formats {
		t, err := time.Parse(format, dateStr)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}

func calculateAge(birthDate time.Time, asOfDate time.Time) int {
	years := asOfDate.Year() - birthDate.Year()
	
	// Check if birthday has occurred in the current year
	if asOfDate.Month() < birthDate.Month() || 
	   (asOfDate.Month() == birthDate.Month() && asOfDate.Day() < birthDate.Day()) {
		years--
	}
	
	return years
}

func parseFullName(name string) (string, string) {
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "", ""
	}
	firstName := parts[0]
	
	if len(parts) == 1 {
		return firstName, ""
	}
	
	lastName := strings.Join(parts[1:], " ")
	return firstName, lastName
}

func parseRelative(relStr string, relationship string) *Relative {
	if relStr == "null" || relStr == "" {
		return nil
	}
	
	firstName, lastName := parseFullName(relStr)
	
	return &Relative{
		FirstName:    firstName,
		LastName:     lastName,
		Relationship: relationship,
	}
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

	// Skip header row
	people := []Person{}
	asOfDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	for _, record := range records[1:] {
		if len(record) < 6 {
			continue
		}

		name := record[0]
		birthdayStr := strings.TrimSpace(record[1])
		diedStr := strings.TrimSpace(record[2])
		fatherStr := strings.TrimSpace(record[3])
		motherStr := strings.TrimSpace(record[4])
		brotherStr := strings.TrimSpace(record[5])
		sisterStr := strings.TrimSpace(record[6])

		firstName, lastName := parseFullName(name)

		birthDate, err := parseDate(birthdayStr)
		if err != nil {
			continue
		}

		birthdayFormatted := birthDate.Format("2006-01-02")

		var age int
		if diedStr != "null" && strings.TrimSpace(diedStr) != "" {
			diedDate, err := parseDate(diedStr)
			if err == nil {
				age = calculateAge(birthDate, diedDate)
			} else {
				age = calculateAge(birthDate, asOfDate)
			}
		} else {
			age = calculateAge(birthDate, asOfDate)
		}

		relatives := []Relative{}
		
		if rel := parseRelative(fatherStr, "Father"); rel != nil {
			relatives = append(relatives, *rel)
		}
		if rel := parseRelative(motherStr, "Mother"); rel != nil {
			relatives = append(relatives, *rel)
		}
		if rel := parseRelative(brotherStr, "Brother"); rel != nil {
			relatives = append(relatives, *rel)
		}
		if rel := parseRelative(sisterStr, "Sister"); rel != nil {
			relatives = append(relatives, *rel)
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		}

		people = append(people, person)
	}

	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}