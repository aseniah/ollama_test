package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Person struct {
	FirstName string      `json:"FirstName"`
	LastName  string      `json:"LastName"`
	Birthday  string      `json:"Birthday"`
	Age       int         `json:"Age"`
	Relatives []Relative  `json:"Relatives"`
}

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("1/2/2006", dateStr)
}

func calculateAge(birthday time.Time) int {
	targetDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	age := targetDate.Year() - birthday.Year()
	if targetDate.Month() < birthday.Month() ||
		(targetDate.Month() == birthday.Month() && targetDate.Day() < birthday.Day()) {
		age--
	}
	return age
}

func splitName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], strings.Join(parts[1:], " ")
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	var people []Person

	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) < 6 {
			continue
		}

		firstName, lastName := splitName(record[0])
		birthday, err := parseDate(record[1])
		if err != nil {
			continue
		}

		age := calculateAge(birthday)
		birthdayStr := birthday.Format("2006-01-02")

		var relatives []Relative

		relationshipTypes := []struct {
			index int
			rel   string
		}{
			{3, "Father"},
			{4, "Mother"},
			{5, "Brother"},
			{6, "Sister"},
		}

		for _, rt := range relationshipTypes {
			if rt.index < len(record) && record[rt.index] != "" && record[rt.index] != "null" {
				relFirstName, relLastName := splitName(record[rt.index])
				relatives = append(relatives, Relative{
					FirstName:    relFirstName,
					LastName:     relLastName,
					Relationship: rt.rel,
				})
			}
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayStr,
			Age:       age,
			Relatives: relatives,
		}
		people = append(people, person)
	}

	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
