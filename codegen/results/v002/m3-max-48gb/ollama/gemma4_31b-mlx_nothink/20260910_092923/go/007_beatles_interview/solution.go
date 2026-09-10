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

func calculateAge(birthdayStr string, referenceDate time.Time) int {
	// Birthday in CSV is M/D/YYYY
	bday, err := time.Parse("1/2/2006", birthdayStr)
	if err != nil {
		return 0
	}

	age := referenceDate.Year() - bday.Year()
	// If the birthday hasn't occurred yet in the reference year
	if referenceDate.YearDay() < bday.YearDay() {
		// This is a simplification; for precise leap year handling:
		if referenceDate.Month() < bday.Month() || (referenceDate.Month() == bday.Month() && referenceDate.Day() < bday.Day()) {
			age--
		}
	} else {
		if referenceDate.Month() < bday.Month() || (referenceDate.Month() == bday.Month() && referenceDate.Day() < bday.Day()) {
			age--
		}
	}
	return age
}

func parseName(fullName string) (string, string) {
	if fullName == "null" || fullName == "" {
		return "", ""
	}
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	// Treat the last part as the last name, everything before as first name
	firstName := strings.Join(parts[:len(parts)-1], " ")
	lastName := parts[len(parts)-1]
	return firstName, lastName
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

	if len(records) == 0 {
		fmt.Println("[]")
		return
	}

	referenceDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)
	var people []Person

	// Skip header (index 0)
	for i := 1; i < len(records); i++ {
		row := records[i]
		// Name, Birthday, Died, Father, Mother, Brother, Sister
		fullName := row[0]
		bdayStr := row[1]
		
		fName, lName := parseName(fullName)
		
		// Format birthday to YYYY-MM-DD
		t, _ := time.Parse("1/2/2006", bdayStr)
		formattedBday := t.Format("2006-01-02")

		person := Person{
			FirstName: fName,
			LastName:  lName,
			Birthday:  formattedBday,
			Age:       calculateAge(bdayStr, referenceDate),
			Relatives: []Relative{},
		}

		// Relatives columns: 3=Father, 4=Mother, 5=Brother, 6=Sister
		relLabels := []string{"Father", "Mother", "Brother", "Sister"}
		for j, label := range relLabels {
			colIdx := j + 3
			if colIdx < len(row) {
				relName := row[colIdx]
				if relName != "null" && relName != "" {
					rfName, rlName := parseName(relName)
					person.Relatives = append(person.Relatives, Relative{
						FirstName:    rfName,
						LastName:     rlName,
						Relationship: label,
					})
				}
			}
		}
		people = append(people, person)
	}

	output, _ := json.MarshalIndent(people, "", "  ")
	fmt.Println(string(output))
}