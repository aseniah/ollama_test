package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
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

// Person represents the transformed data for each person in the input CSV.
type Person struct {
	FirstName string     `json:"FirstName"`
	LastName  string     `json:"LastName"`
	Birthday  string     `json:"Birthday"`
	Age       int        `json	: "Age"`
	Relatives []Relative `json:"Relatives"`
}

// splitName extracts the first and last names from a full name string.
// As inferred from the example, middle names are omitted.
func splitName(full string) (string, string) {
	parts := strings.Fields(full)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[len(parts)-1]
}

// parseDate parses a date string in M/D/YYYY format.
func parseDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" || strings.ToLower(dateStr) == "null" {
		return time.Time{}, nil
	}

	parts := strings.Split(dateStr, "/")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
	}

	var month, day, year int
	if _, err := fmt.Sscanf(parts[0], "%d", &month); err != nil {
		return time.Time{}, err
	}
	if _, err := fmt.Sscanf(parts[1], "%d", &day); err != nil {
		return time.Time{}, err
	}
	if _, err := fmt.Sscanf(parts[2], "%d", &year); err != nil {
		return time.Time{}, err
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

// calculateAge determines age based on a target date. 
// If diedDate is provided, age is calculated at death; otherwise, as of July 1, 2025.
func calculateAge(birth, target time.Time) int {
	if birth.IsZero() {
		return 0
	}
	age := target.Year() - birth.Year()
	// If the birthday hasn't occurred yet in the target year, subtract one year.
	if target.Month() < birth.Month() || (target.Month() == birth.Month() && target.Day() < birth.Day()) {
		age--
	}
	return age
}

func main() {
	// Open the input file
	file, err := os.Open("input/input.csv")
	if err !=
		!= nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip the CSV header
	if _, err := reader.Read(); err != nil {
		return
	}

	var people []Person
	targetDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	relTypes := []string{"Father", "Mother", "Brother", "Sister"}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		// CSV Columns: 0:Name, 1:Birthday, 2:Died, 3:Father, 4:Mother, 5:Brother, 6:Sister
		fullName := record[0]
		firstName, lastName := splitName(fullName)

		birthDate, err := parseDate(record[1])
		if err != nil {
			continue
		}

		diedDate, err := parseDate(record[2])
		if err != nil {
			continue
		}

		var age int
		if !diedDate.IsZero() {
			age = calculateAge(birthDate, diedDate)
		} else {
			age = calculateAge(birthDate, targetDate)
		}

		relatives := []Relative{}
		// Process relative columns 3 through 6
		for i := 3; i <= 6; i++ {
			relVal := record[i]
			if relVal != "" && strings.ToLower(relVal) != "null" {
				rFirst, rLast := splitName(relVal)
				relatives = append(relatives, Relative{
					FirstName:    rFirst,
					LastName:     rLast,
					Relationship: relTypes[i-3],
				})
			}
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthDate.Format("2006-01-02"),
			Age:       age,
			Relatives: relatives,
		})
	}

	// Output the resulting array as JSON to stdout
	jsonOutput, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(jsonOutput))
}