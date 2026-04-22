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
	f, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	var people []Person
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	for i, record := range records {
		if i == 0 || len(record) == 0 || strings.TrimSpace(record[0]) == "" {
			continue
		}

		name := record[0]
		bdayStr := record[1]
		father := record[3]
		mother := record[4]
		brother := record[5]
		sister := record[6]

		// Extract first and last name from full name
		parts := strings.Fields(name)
		if len(parts) == 0 {
			continue
		}
		fName := parts[0]
		lName := parts[len(parts)-1]

		// Parse birthday (MM/DD/YYYY)
		bday, err := time.Parse("1/2/2006", bdayStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing date %s: %v\n", bdayStr, err)
			continue
		}

		// Calculate age as of July 1, 2025
		age := refDate.Year() - bday.Year()
		if refDate.Month() < bday.Month() || (refDate.Month() == bday.Month() && refDate.Day() < bday.Day()) {
			age--
		}

		birthdayFormatted := bday.Format("2006-01-02")

		// Build relatives list
		var relatives []Relative
		addRel := func(n, rel string) {
			if n != "" && n != "null" {
				p := strings.Fields(n)
				if len(p) > 0 {
					relatives = append(relatives, Relative{
						FirstName:    p[0],
						LastName:     p[len(p)-1],
						Relationship: rel,
					})
				}
			}
		}

		addRel(father, "Father")
		addRel(mother, "Mother")
		addRel(brother, "Brother")
		addRel(sister, "Sister")

		people = append(people, Person{
			FirstName: fName,
			LastName:  lName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		})
	}

	out, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}