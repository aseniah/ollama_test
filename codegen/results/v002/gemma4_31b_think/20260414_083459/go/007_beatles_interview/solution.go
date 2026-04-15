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

func splitName(full := string) (string, string) {
	parts := strings.Fields(full)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[len(parts)-1]
}

func calculateAge(birth, death, ref time.Time) int {
	end := ref
	if !death.IsZero() {
		end = death
	}

	age := end.Year() - birth.Year()
	if end.Month() < birth.Month() || (end.Month() == birth.Month() && end.Day() < birth.Day()) {
		age--
	}
	return age
}

func parseCSVDate(s string) time.Time {
	if s == "null" || s == "" {
		return time.Time{}
	}
	// Input format is M/D/YYYY
	t, err := time.Parse("1/2/2006", s)
	if err != nil {
		return time.Time{}
	}
	return t
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

	// Reference date: July 1, 2025
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	// Skip header row
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 7 {
			continue
		}

		name := row[0]
		birthdayStr := row[1]
		diedStr := row[2]
		fatherStr := row[3]
		motherStr := row[4]
		brotherStr := row[5]
		sisterStr := row[6]

		firstName, lastName := splitName(name)
		birthDate := parseCSVDate(birthdayStr)
		deathDate := parseCSVDate(diedStr)

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthDate.Format("2006-01-02"),
			Age:       calculateAge(birthDate, deathDate, refDate),
			Relatives: []Relative{},
		}

		// Process relatives
		relData := []struct {
			val  string
			rel  string
		}{
			{fatherStr, "Father"},
			{motherStr, "Mother"},
			{brotherStr, "Brother"},
			{sisterStr, "Sister"},
		}

		for _, rd := range relData {
			if rd.val != "null" && rd.val != "" {
				fName, lName := splitName(rd.val)
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    fName,
					LastName:     lName,
					Relationship: rd.rel,
				})
			}
		}

		people = append(people, person)
	}

	jsonOutput, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}

	fmt.Println(string(jsonOutput))
}