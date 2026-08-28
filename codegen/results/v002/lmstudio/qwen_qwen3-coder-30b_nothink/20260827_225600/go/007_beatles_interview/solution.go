package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

func calculateAge(birthDate string, referenceDate time.Time) int {
	birth, err := time.Parse("1/2/2006", birthDate)
	if err != nil {
		// Try different date format if needed
		birth, err = time.Parse("1/2/2006", birthDate)
		if err != nil {
			log.Fatalf("Failed to parse date: %v", err)
		}
	}

	years := referenceDate.Year() - birth.Year()
	months := referenceDate.Month() - birth.Month()
	days := referenceDate.Day() - birth.Day()

	if months < 0 || (months == 0 && days < 0) {
		years--
	}

	return years
}

func main() {
	// Read CSV file
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

	// Create the result array
	var people []Person

	// Parse the header row
	header := records[0]

	// Set reference date (July 1, 2025)
	referenceDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	// Process each record (skip header)
	for i := 1; i < len(records); i++ {
		record := records[i]
		person := Person{}

		// Split name and extract first and last names
		fullName := record[0]
		nameParts := splitName(fullName)
		person.FirstName = nameParts[0]
		person.LastName = nameParts[1]

		// Parse birthday
		person.Birthday = record[1]
		
		// Convert to the expected format (YYYY-MM-DD)
		// The input format is M/D/YYYY or MM/DD/YYYY
		birthDate, err := parseDate(record[1])
		if err != nil {
			log.Fatalf("Failed to parse date for %s: %v", fullName, err)
		}
		person.Birthday = birthDate

		// Calculate age as of July 1, 2025
		person.Age = calculateAge(person.Birthday, referenceDate)

		// Parse relatives
		person.Relatives = []Relative{}

		// Parse father
		if record[3] != "null" && record[3] != "" {
			fatherParts := splitName(record[3])
			person.Relatives = append(person.Relatives, Relative{
				FirstName:    fatherParts[0],
				LastName:     fatherParts[1],
				Relationship: "Father",
			})
		}

		// Parse mother
		if record[4] != "null" && record[4] != "" {
			motherParts := splitName(record[4])
			person.Relatives = append(person.Relatives, Relative{
				FirstName:    motherParts[0],
				LastName:     motherParts[1],
				Relationship: "Mother",
			})
		}

		// Parse brother
		if record[5] != "null" && record[5] != "" {
			brotherParts := splitName(record[5])
			person.Relatives = append(person.Relatives, Relative{
				FirstName:    brotherParts[0],
				LastName:     brotherParts[1],
				Relationship: "Brother",
			})
		}

		// Parse sister
		if record[6] != "null" && record[6] != "" {
			sisterParts := splitName(record[6])
			person.Relatives = append(person.Relatives, Relative{
				FirstName:    sisterParts[0],
				LastName:     sisterParts[1],
				Relationship: "Sister",
			})
		}

		people = append(people, person)
	}

	// Output JSON array
	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}

// Helper function to split name into first and last name
func splitName(fullName string) []string {
	parts := []string{}
	names := []string{}

	// Split by spaces, but keep parts that might have spaces in names
	// (like "John Winston Lennon")
	// For now, simple approach: assume first and last name or just last name
	parts = []string{}
	for _, part := range fullName {
		if part == ' ' {
			parts = append(parts, " ")
		} else {
			if len(parts) == 0 {
				parts = append(parts, string(part))
			} else {
				parts[len(parts)-1] = parts[len(parts)-1] + string(part)
			}
		}
	}

	// Special case for names like "John Winston Lennon"
	// More robust approach:
	parts = []string{}
	for i, char := range fullName {
		if i == 0 || (fullName[i-1] == ' ' && fullName[i] != ' ') {
			parts = append(parts, "")
		}
		if fullName[i] != ' ' {
			parts[len(parts)-1] += string(char)
		}
	}

	if len(parts) == 1 {
		return []string{"", parts[0]}
	}
	
	// Check for common first name patterns
	firstName := ""
	lastName := ""
	
	// If there are multiple words, the first is most likely first name
	// and the rest is the last name (this needs proper parsing for accuracy)
	
	// For now, simple logic:
	if len(parts) > 0 {
		firstName = parts[0]
		if len(parts) > 1 {
			lastName = parts[len(parts)-1]
		} else {
			lastName = ""
		}
	} else {
		firstName = ""
		lastName = ""
	}
	
	return []string{firstName, lastName}
}

// Helper function to parse date and reformat it to YYYY-MM-DD
func parseDate(dateStr string) (string, error) {
	// Try different formats
	formats := []string{
		"1/2/2006",
		"1/2/2006",
		"1/2/2006",
		"1/2/2006",
		"1/2/2006",
		"1/2/2006",
		"1/2/2006",
		"1/2/2006",
		"1/2/2006",
		"1/2/2006",
	}
	
	// For now, assuming a specific input format (d/M/YYYY or DD/MM/YYYY)
	// Based on examples given:
	// "10/9/1940" -> "1940-10-09"
	// "6/18/1942" -> "1942-06-18"
	
	parts := []string{}
	currentPart := ""
	for i, char := range dateStr {
		if char == '/' {
			parts = append(parts, currentPart)
			currentPart = ""
		} else {
			currentPart += string(char)
		}
	}
	parts = append(parts, currentPart)

	if len(parts) != 3 {
		return "", fmt.Errorf("invalid date format: %s", dateStr)
	}

	// Convert to YYYY-MM-DD format
	// If MM/DD/YYYY format (as shown in example)
	month, day, year := parts[0], parts[1], parts[2]

	// Ensure zero-padding
	if len(month) == 1 {
		month = "0" + month
	}
	if len(day) == 1 {
		day = "0" + day
	}

	return fmt.Sprintf("%s-%s-%s", year, month, day), nil
}