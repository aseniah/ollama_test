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

func splitFullName(name string) (string, string) {
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	firstName := parts[0]
	lastName := parts[len(parts)-1]
	return firstName, lastName
}

func convertBirthday(birthday string) (string, time.Time) {
	// Input format: MM/DD/YYYY
	t, err := time.Parse("1/2/2006", birthday)
	if err != nil {
		// Try with leading zeros
		t, err = time.Parse("1/2/2006", birthday)
		if err != nil {
			return "", time.Time{}
		}
	}
	return t.Format("2006-01-02"), t
}

func calculateAge(birthday time.Time) int {
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	age := refDate.Year() - birthday.Year()
	// If birthday hasn't occurred yet this year (as of July 1), subtract 1
	if refDate.Month() < birthday.Month() || (refDate.Month() == birthday.Month() && refDate.Day() < birthday.Day()) {
		age--
	}
	return age
}

func addRelative(relatives *[]Relative, name string, relationship string) {
	if name == "" || name == "null" {
		return
	}
	firstName, lastName := splitFullName(name)
	*relatives = append(*relatives, Relative{
		FirstName:    firstName,
		LastName:     lastName,
		Relationship: relationship,
	})
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
		fmt.Print("[]\n")
		return
	}

	header := records[0]
	var people []Person

	for _, record := range records[1:] {
		if len(record) < 7 {
			continue
		}
		// Map fields by header index
		nameField := header[0]  // Name
		_ = nameField

		fullName := record[0]
		birthdayStr := record[1]
		_ = record[2] // Died - not used
		father := record[3]
		mother := record[4]
		brother := record[5]
		sister := record[6]

		firstName, lastName := splitFullName(fullName)
		birthdayFormatted, birthdayTime := convertBirthday(birthdayStr)
		age := calculateAge(birthdayTime)

		var relatives []Relative
		addRelative(&relatives, father, "Father")
		addRelative(&relatives, mother, "Mother")
		addRelative(&relatives, brother, "Brother")
		addRelative(&relatives, sister, "Sister")

		// If no relatives, output empty array not null
		if relatives == nil {
			relatives = []Relative{}
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		})
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(output))
	_ = io.Discard // ensure io is used if needed
	_ = strconv.Itoa(0) // ensure strconv is used if needed
}