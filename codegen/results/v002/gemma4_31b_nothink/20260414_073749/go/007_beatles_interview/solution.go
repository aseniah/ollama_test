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

func calculateAge(birthdayStr string, referenceDate time.Time) int {
	// Input format from CSV seems to be M/D/YYYY (e.g., 10/9/1940)
	t, err := time.Parse("1/2/2006", birthdayStr)
	if err != nil {
		return 0
	}

	age := referenceDate.Year() - t.Year()
	// Check if birthday has occurred yet in the reference year
	if referenceDate.YearDay() < t.YearDay() {
		// This is a simplification; for production we'd handle leap years more precisely
		// but for standard age calculation:
		if referenceDate.Month() < t.Month() || (referenceDate.Month() == t.Month() && referenceDate.Day() < t.Day()) {
			age--
		}
	}
	return age
}

func parseName(fullName string) (string, string) {
	if fullName == "null" || fullName == "" {
		return "", ""
	}
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	first := parts[0]
	last := ""
	if len(parts) > 1 {
		last = parts[len(parts)-1]
	}
	return first, last
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

	if len(records) < 1 {
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

		// Parse Main Person
		fullPersonName := row[0]
		fName, lName := parseName(fullPersonName)
		
		// Format Birthday to YYYY-MM-DD
		bDayRaw := row[1]
		t, _ := time.Parse("1/2/2006", bDayRaw)
		bDayFormatted := t.Format("2006-01-02")

		person := Person{
			FirstName: fName,
			LastName:  lName,
			Birthday:  bDayFormatted,
			Age:       calculateAge(bDayRaw, refDate),
			Relatives: []Relative{},
		}

		// Relatives mapping: Father, Mother, Brother, Sister
		relCols := []struct {
			index int
			label string
		}{
			{3, "Father"},
			{4, "Mother"},
			{5, "Brother"},
			{6, "Sister"},
		}

		for _, relCol := range relCols {
			relName := row[relCol.index]
			if relName != "null" && relName != "" {
				rfName, rlName := parseName(relName)
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    rfName,
					LastName:     rlName,
					Relationship: relCol.label,
				})
			}
		}

		people = append(people, person)
	}

	output, _ := json.MarshalIndent(people, "", "  ")
	fmt.Println(string(output))
}