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

func splitFullName(fullName string) (string, string) {
	words := strings.Fields(fullName)
	if len(words) == 1 {
		return words[0], ""
	}
	return words[0], words[len(words)-1]
}

func calculateAge(birthDate time.Time, referenceDate time.Time) int {
	age := referenceDate.Year() - birthDate.Year()
	if referenceDate.Month() < birthDate.Month() ||
		(referenceDate.Month() == birthDate.Month() && referenceDate.Day() < birthDate.Day()) {
		age--
	}
	return age
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

	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	for i, record := range records {
		if i == 0 {
			continue // skip header
		}

		if len(record) < 8 {
			continue
		}

		name, birthdayStr, diedStr, father, mother, brother, sister := record[0], record[1], record[2], record[3], record[4], record[5], record[6]

		firstName, lastName := splitFullName(name)

		// Parse birthday: MM/DD/YYYY
		birthday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			birthday, err = time.Parse("1/2/2006", birthdayStr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing birthday %s: %v\n", birthdayStr, err)
				continue
			}
		}

		birthdayFormatted := fmt.Sprintf("%04d-%02d-%02d", birthday.Year(), int(birthday.Month()), birthday.Day())

		// Calculate age
		var age int
		if diedStr == "null" {
			age = calculateAge(birthday, referenceDate)
		} else {
			died, err := time.Parse("1/2/2006", diedStr)
			if err != nil {
				died, err = time.Parse("1/2/2006", diedStr)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing died date %s: %v\n", diedStr, err)
					age = calculateAge(birthday, referenceDate)
				} else {
					age = calculateAge(birthday, died)
				}
			} else {
				age = calculateAge(birthday, died)
			}
		}

		// Build relatives
		var relatives []Relative

		relativesData := []struct {
			name string
			rel  string
		}{
			{father, "Father"},
			{mother, "Mother"},
			{brother, "Brother"},
			{sister, "Sister"},
		}

		for _, rd := range relativesData {
			if rd.name == "null" || rd.name == "" {
				continue
			}
			fName, lName := splitFullName(rd.name)
			relatives = append(relatives, Relative{
				FirstName:    fName,
				LastName:     lName,
				Relationship: rd.rel,
			})
		}

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

	fmt.Println(string(output))
}