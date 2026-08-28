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
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	if len(records) == 0 {
		fmt.Print("[]")
		return
	}

	// Reference date: July 1, 2025
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	// Skip header row
	for i := 1; i < len(records); i++ {
		row := records[i]
		
		// Parse Full Name
		fullName := row[0]
		nameParts := strings.Fields(fullName)
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse Birthday (M/D/YYYY)
		bdayStr := row[1]
		bday, _ := time.Parse("1/2/2006", bdayStr)
		
		// Calculate Age as of July 1, 2025
		age := refDate.Year() - bday.Year()
		if refDate.Month() < bday.Month() || (refDate.Month() == bday.Month() && refDate.Day() < bday.Day()) {
			age--
		}

		// Handle death date: If they died, age is fixed at time of death
		diedStr := row[2]
		if diedStr != "null" && diedStr != "" {
			deathDate, err := time.Parse("1/2/2006", diedStr)
			if err == nil {
				age = deathDate.Year() - bday.Year()
				if deathDate.Month() < bday.Month() || (deathDate.Month() == bday.Month() && deathDate.Day() < bday.Day()) {
					age--
				}
			}
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  bday.Format("2006-01-02"),
			Age:       age,
			Relatives: []Relative{},
		}

		// Map relative columns
		relCols := []struct {
			index int
			rel   string
		}{
			{3, "Father"},
			{4, "Mother"},
			{5, "Brother"},
			{6, "Sister"},
		}

		for _, col := range relCols {
			val := row[col.index]
			if val != "null" && val != "" {
				parts := strings.Fields(val)
				if len(parts) >= 2 {
					person.Relatives = append(person.Relatives, Relative{
						FirstName:    parts[0],
						LastName:     parts[len(parts)-1],
						Relationship: col.rel,
					})
				}
			}
		}

		people = append(people, person)
	}

	output, _ := json.MarshalIndent(people, "", "  ")
	fmt.Println(string(output))
}