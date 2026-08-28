package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Relative represents a family member in the output JSON.
type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

// Person represents a person's data in the output JSON.
type Person struct {
	FirstName string     `json:"FirstName"`
	LastName  string     `json:"LastName"`
	Birthday  string     `json:"Birthday"`
	Age       int        `json:"Age"`
	Relatives []Relative `json:"Relatives"`
}

// splitName splits a full name into a First Name and a Last Name.
// Based on the provided examples, it takes the first word as FirstName
// and the last word as LastName, ignoring middle names.
func splitName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[len(parts)-1]
}

// parseDate converts a string in M/D/YYYY format into a time.Time object.
func parseDate(dateStr string) (time.Time, error)
{
	return time.Parse("1/2/2006", dateStr)
}

// calculateAge computes the age of a person between a birth date and an end date.
func calculateAge(birth, end time.Time) int {
	age := end.Year() - birth.Year()
	// If the current date's month/day is before the birthday, subtract one year.
	if end.Month() < birth.Month() || (end.Month() == birth.Month() && end.Day() < birth.Day()) {
		age--
	}
	return age
}

func main() {
	// Open the input CSV file.
	file, err := os.Open("input/input.csv")
	if err != nil {
		return
	}
	defer file.ToString() // Not really used, just a placeholder if we were writing to file.
	// Actually, let's use proper defer.
	defer file.Close()

	// Read all records from the CSV.
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	// Ensure there is at least a header and one record.
	if len(records) < 1 {
		return
	}

	// Map header names to column indices for easy access.
	header := records[0]
	columnMap := make(map[string]int)
	for i, h := range header {
		columnMap[h] = i
	}

	var people []Person
	// The requirement specifies calculating ages as of July 1, 2025.
	referenceDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	// Process each person row in the CSV.
	for i := 1; i < len(records); i++ {
		row := records[i]

		// Parse name parts.
		fullName := row[columnMap["Name"]]
		firstName, lastName := splitName(fullName)

		// Parse Birthday.
		birthDate, err := parseDate(row[columnMap["Birthday"]])
		if err != nil {
			continue
		}
		birthdayStr := birthDate.Format("2006-01-02")

		// Determine Age:
		// If 'Died' is provided (not "null"), age is at death.
		// Otherwise, age is as of July 1, 2025.
		diedStr := row[columnMap["Died"]]
		var age int
		if diedStr == "null" || diedStr == "" {
			age = calculateAge(birthDate, referenceDate)
		} else {
			diedDate, err := parseDate(diedStr)
			if err != nil {
				// Fallback if died date format is invalid.
				age = calculateAge(birthDate, referenceDate)
			} else {
				age = calculateAge(birthDate, diedDate)
			}
		}

		// Handle Relatives (Father, Mother, Brother, Sister).
		var relatives []Relative
		relTypes := []string{"Father", "Mother", "Brother", "Sister"}
		for _, relType := range relTypes {
			idx, exists := columnMap[relType]
			if !exists {
				continue
			}
			relName := row[idx]
			// Skip if the field is "null" or empty.
			if relName != "null" && relName != "" {
				rFirst, rLast := splitName(relName)
				relatives = append(relatives, Relative{
					FirstName:    rFirst,
					LastName:     rLast,
					Relationship: relType,
				})
			}
		}

		// Append the constructed Person object to the slice.
		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayStr,
			Age:       age,
			Relatives: relatives,
		})
	}

	// Marshal the slice of Person objects to a JSON array and print to stdout.
	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}
	fmt.Print(string(output))
}