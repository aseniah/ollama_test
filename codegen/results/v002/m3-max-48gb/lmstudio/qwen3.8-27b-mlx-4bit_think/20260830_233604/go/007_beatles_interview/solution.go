package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func parseDate(s string) (month, day, year int) {
	parts := strings.Split(s, "/")
	month, _ = strconv.Atoi(parts[0])
	day, _ = strconv.Atoi(parts[1])
	year, _ = strconv.Atoi(parts[2])
	return
}

func calculateAge(birthY, birthM, birthD, refY, refM, refD int) int {
	age := refY - birthY
	if refM < birthM || (refM == birthM && refD < birthD) {
		age--
	}
	return age
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

	for i, record := range records {
		if i == 0 { // skip header
			continue
		}

		name := strings.TrimSpace(record[0])
		nameParts := strings.Fields(name)
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse birthday: MM/DD/YYYY
		birthM, birthD, birthY := parseDate(strings.TrimSpace(record[1]))
		birthdayStr := fmt.Sprintf("%04d-%02d-%02d", birthY, birthM, birthD)

		// Determine reference date for age
		refY, refM, refD := 2025, 7, 1
		diedField := strings.TrimSpace(record[2])
		if diedField != "null" && diedField != "" {
			diedM, diedD, diedY := parseDate(diedField)
			refY, refM, refD = diedY, diedM, diedD
		}

		age := calculateAge(birthY, birthM, birthD, refY, refM, refD)

		var relatives []Relative

		relFields := []struct {
			index        int
			relationship string
		}{
			{3, "Father"},
			{4, "Mother"},
			{5, "Brother"},
			{6, "Sister"},
		}

		for _, rf := range relFields {
			val := strings.TrimSpace(record[rf.index])
			if val != "null" && val != "" {
				parts := strings.Fields(val)
				relFirstName := parts[0]
				relLastName := parts[len(parts)-1]
				relatives = append(relatives, Relative{
					FirstName:    relFirstName,
					LastName:     relLastName,
					Relationship: rf.relationship,
				})
			}
		}

		if relatives == nil {
			relatives = []Relative{}
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayStr,
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