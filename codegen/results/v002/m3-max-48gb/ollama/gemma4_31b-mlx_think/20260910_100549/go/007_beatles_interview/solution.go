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
	FirstName string     `json:"FirstName"`
	LastName  string     `json:"LastName"`
	Birthday  string     `json:"Birthday"`
	Age       int        `json:"Age"`
	Relatives []Relative `json:"Relatives"`
}

func calculateAge(birthDate, referenceDate time.Time) int {
	age := referenceDate.Year() - birthDate.Year()
	if referenceDate.Month() < birthDate.Month() || (referenceDate.Month() == birthDate.Month() && referenceDate.Day() < birthDate.Day()) {
		age--
	}
	return age
}

func parseName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[len(parts)-1]
}

func parseDate(dateStr string) (time.Time, bool) {
	if dateStr == "null" || dateStr == "" {
		return time.Time{}, false
	}
	// Input format is M/D/YYYY
	t, err := time.Parse("1/2/2006", dateStr)
	if err != nil {
		return time.Time{}, false
	}
	return t, true
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	if len(records) == 0 {
		fmt.Println("[]")
		return
	}

	// Assume first row is header
	header := records[0]
	data := records[1:]

	targetDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)
	var people []Person

	for _, row := range data {
		rowMap := make(map[string]string)
		for i, val := range row {
			if i < len(header) {
				rowMap[header[i]] = val
			}
		}

		firstName, lastName := parseName(rowMap["Name"])
		birthDate, okB := parseDate(rowMap["Birthday"])
		if !okB {
			continue
		}

		deathDate, okD := parseDate(rowMap["Died"])
		effectiveDate := targetDate
		if okD {
			effectiveDate = deathDate
		}

		age := calculateAge(birthDate, effectiveDate)

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthDate.Format("2006-01-02"),
			Age:       age,
			Relatives: []Relative{},
		}

		relationships := []string{"Father", "Mother", "Brother", "Sister"}
		for _, relType := range relationships {
			relName := rowMap[relType]
			if relName == "null" || relName == "" {
				continue
			}
			fName, lName := parseName(relName)
			person.Relatives = append(person.Relatives, Relative{
				FirstName:    fName,
				LastName:     lName,
				Relationship: relType,
			})
		}

		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(output))
}