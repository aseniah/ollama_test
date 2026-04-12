package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func parseName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	firstName := parts[0]
	lastName := ""
	if len(parts) > 1 {
		lastName = parts[len(parts)-1]
	}
	return firstName, lastName
}

func parseRelative(fullName, relationship string) *Relative {
	if fullName == "" || fullName == "null" || fullName == "NULL" {
		return nil
	}
	firstName, lastName := parseName(fullName)
	return &Relative{
		FirstName:    firstName,
		LastName:     lastName,
		Relationship: relationship,
	}
}

func calculateAge(birthdayStr string) int {
	layout := "1/2/2006"
	birthday, err := time.Parse(layout, birthdayStr)
	if err != nil {
		return 0
	}

	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)
	age := refDate.Year() - birthday.Year()

	// Adjust if birthday hasn't occurred yet in the reference year
	if refDate.YearDay() < birthday.YearDay() {
		age--
	}
	
	// Handle edge case where Month/Day logic needs more precision than YearDay
	// YearDay is 0-365, but let's double check month/day specifically
	if age < 0 {
		age = 0
	}
	
	// More precise age calculation
	anniversary := time.Date(refDate.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)
	if refDate.Before(anniversary) {
		age--
	}

	return age
}

func formatBirthdate(birthdayStr string) string {
	// Input format: M/D/YYYY or MM/D/YYYY or M/DD/YYYY
	// Target format: YYYY-MM-DD
	
	// We need to parse the specific input format
	// The input format doesn't have leading zeros consistently
	layout := "1/2/2006"
	t, err := time.Parse(layout, birthdayStr)
	if err != nil {
		return birthdayStr
	}
	return t.Format("2006-01-02")
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

	if len(records) < 2 {
		fmt.Println("[]")
		return
	}

	// Skip header
	dataRows := records[1:]
	var people []Person

	for _, row := range dataRows {
		if len(row) < 7 {
			continue
		}

		name := row[0]
		birthdayRaw := row[1]
		father := row[3]
		mother := row[4]
		brother := row[5]
		sister := row[6]

		firstName, lastName := parseName(name)
		birthdayFormatted := formatBirthdate(birthdayRaw)
		age := calculateAge(birthdayRaw)

		var relatives []Relative
		if rel := parseRelative(father, "Father"); rel != nil {
			relatives = append(relatives, *rel)
		}
		if rel := parseRelative(mother, "Mother"); rel != nil {
			relatives = append(relatives, *rel)
		}
		if rel := parseRelative(brother, "Brother"); rel != nil {
			relatives = append(relatives, *rel)
		}
		if rel := parseRelative(sister, "Sister"); rel != nil {
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

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(people); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}