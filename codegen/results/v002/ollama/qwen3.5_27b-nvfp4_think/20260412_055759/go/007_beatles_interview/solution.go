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
	FirstName string   `json:"FirstName"`
	LastName  string   `json:"LastName"`
	Birthday  string   `json:"Birthday"`
	Age       int      `json:"Age"`
	Relatives []Relative `json:"Relatives"`
}

func parseName(fullName string) (firstName, lastName string) {
	parts := strings.Fields(fullName)
	if len(parts) == 1 {
		firstName = parts[0]
		lastName = ""
	} else {
		firstName = parts[0]
		lastName = parts[len(parts)-1]
	}
	return
}

func parseDate(dateStr string) (time.Time, error) {
	formats := []string{
		"1/2/2006",
		"1/2/06",
		"M/D/YYYY",
	}
	for _, f := range formats {
		if t, err := time.Parse(f, dateStr); err == nil {
			return t, nil
		}
	}
	// Try YYYY-MM-DD format
	t, err := time.Parse("2006-01-02", dateStr)
	if err == nil {
		return t, nil
	}
	return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}

func calculateAge(birthDate, referenceDate time.Time) int {
	age := referenceDate.Year() - birthDate.Year()
	
	// Check if birthday has occurred this year
	birthdayThisYear := time.Date(referenceDate.Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, referenceDate.Location())
	if referenceDate.Before(birthdayThisYear) {
		age--
	}
	
	return age
}

func processRelatives(name string, relationship string) Relative {
	firstName, lastName := parseName(name)
	return Relative{
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

	var people []Person
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}

		fullName, _ := parseName(record[0])
		
		birthDate, err := parseDate(record[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing birthday for %s: %v\n", fullName, err)
			continue
		}

		// Format birthday as YYYY-MM-DD
		birthdayStr := birthDate.Format("2006-01-02")

		// Determine reference date for age calculation
		var ref time.Time
		if record[2] == "null" || strings.TrimSpace(record[2]) == "" {
			ref = refDate
		} else {
			diedDate, err := parseDate(record[2])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing death date for %s: %v\n", fullName, err)
				continue
			}
			ref = diedDate
		}

		age := calculateAge(birthDate, ref)

		var relatives []Relative
		
		// Process Father
		if record[3] != "null" && strings.TrimSpace(record[3]) != "" {
			relatives = append(relatives, processRelatives(record[3], "Father"))
		}
		
		// Process Mother
		if record[4] != "null" && strings.TrimSpace(record[4]) != "" {
			relatives = append(relatives, processRelatives(record[4], "Mother"))
		}
		
		// Process Brother
		if record[5] != "null" && strings.TrimSpace(record[5]) != "" {
			relatives = append(relatives, processRelatives(record[5], "Brother"))
		}
		
		// Process Sister
		if record[6] != "null" && strings.TrimSpace(record[6]) != "" {
			relatives = append(relatives, processRelatives(record[6], "Sister"))
		}

		person := Person{
			FirstName: firstName,
			LastName:  _,
			Birthday:  birthdayStr,
			Age:       age,
			Relatives: relatives,
		}

		if len(person.Relatives) == 0 {
			person.Relatives = []Relative{} // Ensure empty array, not null
		}

		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", " ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}