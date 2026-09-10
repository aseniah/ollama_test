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
		fmt.Println("[]")
		return
	}

	// CSV Header: Name,Birthday,Died,Father,Mother,Brother,Sister
	var people []Person
	referenceDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		nameParts := strings.Split(row[0], " ")
		firstName := nameParts[0]
		lastName := ""
		if len(nameParts) > 1 {
			// If there are middle names, combine them into last name or just take the last part?
			// Based on expected_format.json: "John Winston Lennon" -> "John" "Lennon"
			lastName = nameParts[len(nameParts)-1]
		}

		bdayStr := row[1] // format M/D/YYYY
		bday, _ := time.Parse("1/2/2006", bdayStr)
		
		diedStr := row[2]
		var diedDate time.Time
		hasDied := diedStr != "null" && diedStr != ""
		if hasDied {
			diedDate, _ = time.Parse("1/2/2006", diedStr)
		}

		// Calculate Age
		var age int
		endDate := referenceDate
		if hasDied {
			endDate = diedDate
		}
		
		age = endDate.Year() - bday.Year()
		if endDate.YearDay() < bday.YearDay() {
			age--
		}

		p := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  bday.Format("2006-01-02"),
			Age:       age,
			Relatives: []Relative{},
		}

		// Relatives mapping
		relColumns := []struct {
			index int
			label string
		}{
			{3, "Father"},
			{4, "Mother"},
			{5, "Brother"},
			{6, "Sister"},
		}

		for _, col := range relColumns {
			val := row[col.index]
			if val == "null" || val == "" {
				continue
			}
			relParts := strings.Split(val, " ")
			fName := relParts[0]
			lName := ""
			if len(relParts) > 1 {
				lName = relParts[len(relParts)-1]
			}
			p.Relatives = append(p.Relatives, Relative{
				FirstName:    fName,
				LastName:     lName,
				Relationship: col.label,
			})
		}
		people = append(people, p)
	}

	output, _ := json.MarshalIndent(people, "", "  ")
	fmt.Println(string(output))
}