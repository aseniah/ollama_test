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

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading CSV:", err)
		os.Exit(1)
	}

	// Parse header
	header := records[0]
	// header: Name,Birthday,Died,Father,Mother,Brother,Sister

	people := make([]Person, 0)

	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) < 7 {
			continue
		}

		fullName := record[0]
		birthdayStr := record[1]
		// died := record[2] // not used for age calculation, age is as of 2025-07-01

		// Split name into first and last
		parts := strings.Split(strings.TrimSpace(fullName), " ")
		var firstName, lastName string
		if len(parts) == 1 {
			firstName = parts[0]
			lastName = ""
		} else {
			firstName = parts[0]
			lastName = strings.Join(parts[1:], " ")
		}

		// Parse birthday
		birthday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			birthday, err = time.Parse("1/2/2006", birthdayStr)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing birthday:", err)
			continue
		}

		// Calculate age as of July 1, 2025
		refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
		age := refDate.Year() - birthday.Year()
		// Check if birthday hasn't occurred yet in 2025
		if refDate.Month() < birthday.Month() || (refDate.Month() == birthday.Month() && refDate.Day() < birthday.Day()) {
			age--
		}

		// Format birthday as ISO date
		birthdayISO := birthday.Format("2006-01-02")

		// Build relatives list
		relatives := make([]Relative, 0)

		if father := strings.TrimSpace(record[3]); father != "" && father != "null" {
			parts := strings.Split(strings.TrimSpace(father), " ")
			if len(parts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    parts[0],
					LastName:     strings.Join(parts[1:], " "),
					Relationship: "Father",
				})
			}
		}

		if mother := strings.TrimSpace(record[4]); mother != "" && mother != "null" {
			parts := strings.Split(strings.TrimSpace(mother), " ")
			if len(parts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    parts[0],
					LastName:     strings.Join(parts[1:], " "),
					Relationship: "Mother",
				})
			}
		}

		if brother := strings.TrimSpace(record[5]); brother != "" && brother != "null" {
			parts := strings.Split(strings.TrimSpace(brother), " ")
			if len(parts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    parts[0],
					LastName:     strings.Join(parts[1:], " "),
					Relationship: "Brother",
				})
			}
		}

		if sister := strings.TrimSpace(record[6]); sister != "" && sister != "null" {
			parts := strings.Split(strings.TrimSpace(sister), " ")
			if len(parts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    parts[0],
					LastName:     strings.Join(parts[1:], " "),
					Relationship: "Sister",
				})
			}
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayISO,
			Age:       age,
			Relatives: relatives,
		}

		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}