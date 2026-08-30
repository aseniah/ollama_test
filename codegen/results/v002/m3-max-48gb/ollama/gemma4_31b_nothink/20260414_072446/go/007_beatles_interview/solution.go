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
	t, err := time.Parse("1/2/2006", birthdayStr)
	if err != nil {
		return 0
	}

	age := referenceDate.Year() - t.Year()
	if referenceDate.YearDay() < t.YearDay() {
		age--
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

	// Header: Name, Birthday, Died, Father, Mother, Brother, Sister
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)
	var people []Person

	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		fullName := row[0]
		bDayRaw := row[1]
		
		fName, lName := parseName(fullName)
		
		// Transform date from M/D/YYYY to YYYY-MM-DD
		t, _ := time.Parse("1/2/2006", bDayRaw)
		formattedBday := t.Format("2006-01-02")

		age := calculateAge(bDayRaw, refDate)

		var relatives []Relative
		relCols := []struct {
			index int
			label string
		}{
			{3, "Father"},
			{4, "Mother"},
			{5, "Brother"},
			{6, "Sister"},
		}

		for _, col := range relCols {
			if col.index < len(row) {
				val := row[col.index]
				if val != "null" && val != "" {
					rfName, rlName := parseName(val)
					relatives = append(relatives, Relative{
						FirstName:    rfName,
						LastName:     rlName,
						Relationship: col.label,
					})
				}
			}
		}

		people = append(people, Person{
			FirstName: fName,
			LastName:  lName,
			Birthday:  formattedBday,
			Age:       age,
			Relatives: relatives,
		})
	}

	output, _ := json.MarshalIndent(people, "", "  ")
	fmt.Println(string(output))
}