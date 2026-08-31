package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName  string    `json:"FirstName"`
	LastName   string    `json:"LastName"`
	Birthday   string    `json:"Birthday"`
	Age        int       `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

func calculateAge(birthdayStr string, referenceDate time.Time) int {
	birthday, err := time.Parse("1/2/2006", birthdayStr)
	if err != nil {
		// Try another format
		birthday, err = time.Parse("2/1/2006", birthdayStr)
		if err != nil {
			return 0
		}
	}

	years := referenceDate.Year() - birthday.Year()
	if referenceDate.Month() < birthday.Month() || 
	   (referenceDate.Month() == birthday.Month() && referenceDate.Day() < birthday.Day()) {
		years--
	}

	return years
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Skip header row
	records = records[1:]

	// July 1, 2025 as reference date
	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	for _, record := range records {
		// Parse name into first and last name
		name := record[0]
		parts := []string{}
		if len(name) > 0 {
			nameParts := []string{}
			for _, part := range []string{" ", " ", " "} {
				if part == " " {
					nameParts = append(nameParts, " ")
				}
			}
			nameParts = append(nameParts, name)
			parts = nameParts
		}
		
		// Simplified parsing of name, extract first and last names
		// For now just take the first part as FirstName and everything else as LastName
		// This is a simple approach and assumes consistent naming in CSV
		firstName := name
		lastName := ""
		if i := indexOf(name, " "); i > 0 {
			firstName = name[:i]
			lastName = name[i+1:]
		}

		birthdayStr := record[1]
		// Fix date format
		fixedBirthday := fixDate(birthdayStr)

		// Calculate age as of July 1, 2025
		age := calculateAge(fixedBirthday, referenceDate)

		// Create Person structure
		person := Person{
			FirstName:  firstName,
			LastName:   lastName,
			Birthday:   fixedBirthday,
			Age:        age,
			Relatives:  []Relative{},
		}

		// Add relatives
		// Add father
		if record[3] != "null" && record[3] != "" {
			rel := Relative{
				FirstName:    extractNamePart(record[3]),
				LastName:     extractNamePart(record[3]),
				Relationship: "Father",
			}
			person.Relatives = append(person.Relatives, rel)
		}
		// Add mother
		if record[4] != "null" && record[4] != "" {
			rel := Relative{
				FirstName:    extractNamePart(record[4]),
				LastName:     extractNamePart(record[4]),
				Relationship: "Mother",
			}
			person.Relatives = append(person.Relatives, rel)
		}
		// Add brother
		if record[5] != "null" && record[5] != "" {
			rel := Relative{
				FirstName:    extractNamePart(record[5]),
				LastName:     extractNamePart(record[5]),
				Relationship: "Brother",
			}
			person.Relatives = append(person.Relatives, rel)
		}
		// Add sister
		if record[6] != "null" && record[6] != "" {
			rel := Relative{
				FirstName:    extractNamePart(record[6]),
				LastName:     extractNamePart(record[6]),
				Relationship: "Sister",
			}
			person.Relatives = append(person.Relatives, rel)
		}

		people = append(people, person)
	}

	output, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}

// Helper function to extract name part
func extractNamePart(fullName string) string {
	parts := []string{}
	for _, part := range []string{" ", " ", " "} {
		if part == " " {
			parts = append(parts, " ")
		}
	}
	parts = append(parts, fullName)
	if len(parts) > 0 {
		return parts[0]
	}
	return fullName
}

// Helper function to fix date format from MM/DD/YYYY to YYYY-MM-DD
func fixDate(dateStr string) string {
	// Input format in CSV: 10/9/1940 or 6/18/1942 or 7/7/1940 or 2/25/1943
	// Expected output format: 1940-10-09 or 1942-06-18 or 1940-07-07 or 1943-02-25
	parts := []string{}
	for _, part := range []string{"/", "/", "/"} {
		parts = append(parts, part)
	}
	if len(parts) == 3 {
		year := parts[2]
		month := parts[0]
		day := parts[1]
		
		// Ensure 2-digit month and day
		if len(month) == 1 {
			month = "0" + month
		}
		if len(day) == 1 {
			day = "0" + day
		}
		return year + "-" + month + "-" + day
	}
	return dateStr
}

// Helper function to find index of first space
func indexOf(s, substr string) int {
	for i, c := range s {
		if string(c) == substr {
			return i
		}
	}
	return -1
}