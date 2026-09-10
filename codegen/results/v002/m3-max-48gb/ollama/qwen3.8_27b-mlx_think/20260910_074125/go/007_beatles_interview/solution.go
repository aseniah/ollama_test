package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

func calcAge(birthDate time.Time, refDate time.Time) int {
	age := refDate.Year() - birthDate.Year()
	if refDate.Month() < birthDate.Month() ||
		(refDate.Month() == birthDate.Month() && refDate.Day() < birthDate.Day()) {
		age--
	}
	return age
}

func main() {
	f, err := os.Open("input/input.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var people []Person

	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		// Parse Name: first word = FirstName, last word = LastName
		nameParts := strings.Fields(row[0])
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse Birthday M/D/YYYY -> YYYY-MM-DD
		birthdayParts := strings.Split(row[1], "/")
		month, _, _ := parseInt(birthdayParts[0])
		day, _, _ := parseInt(birthdayParts[1])
		year, _, _ := parseInt(birthdayParts[2])
		birthDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		birthdayStr := fmt.Sprintf("%04d-%02d-%02d", year, month, day)

		// Calculate age
		var age int
		died := row[3] // "Died" column
		if strings.ToLower(died) != "null" && died != "" {
			diedParts := strings.Split(died, "/")
			dMonth, _, _ := parseInt(diedParts[0])
			dDay, _, _ := parseInt(diedParts[1])
			dYear, _, _ := parseInt(diedParts[2])
			deathDate := time.Date(dYear, time.Month(dMonth), dDay, 0, 0, 0, 0, time.UTC)
			age = calcAge(birthDate, deathDate)
		} else {
			refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)
			age = calcAge(birthDate, refDate)
		}

		// Build relatives
		relatives := []Relative{}

		relOrder := []struct {
			ColIndex    int
			Relationship string
		}{
			{4, "Father"},
			{5, "Mother"},
			{6, "Brother"},
			{7, "Sister"},
		}

		for _, rel := range relOrder {
			val := row[rel.ColIndex]
			if strings.ToLower(val) == "null" || val == "" {
				continue
			}
			relParts := strings.Fields(val)
			relFirst := relParts[0]
			relLast := relParts[len(relParts)-1]
			relatives = append(relatives, Relative{
				FirstName:    relFirst,
				LastName:     relLast,
				Relationship: rel.Relationship,
			})
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayStr,
			Age:       age,
			Relatives: relatives,
		})
	}

	// Marshal to JSON
	output, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}

	// Pretty print
	var buf strings.Builder
	json.Indent(&buf, output, "", "  ")
	// The expected format uses specific indentation; let's just print with 1-space indent style
	// Actually let's just use MarshalIndent
	indented, _ := json.MarshalIndent(people, "", "  ")
	fmt.Println(string(indented))
	_ = buf
	_ = sort.IntSlice{} // avoid unused import
}

func parseInt(s string) (int, string, error) {
	var n int
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	return n, "", nil
}