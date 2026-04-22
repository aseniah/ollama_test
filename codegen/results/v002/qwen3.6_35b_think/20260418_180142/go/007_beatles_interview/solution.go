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
	FirstName string      `json:"FirstName"`
	LastName  string      `json:"LastName"`
	Birthday  string      `json:"Birthday"`
	Age       int         `json:"Age"`
	Relatives []Relative  `json:"Relatives"`
}

func main() {
	f, err := os.Open("input/input.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	if len(records) < 2 {
		return
	}

	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	var result []Person

	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}
		if len(record) < 2 {
			continue
		}

		// Parse First & Last name
		nameParts := strings.Fields(record[0])
		if len(nameParts) < 2 {
			continue
		}
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse Birthday & format as YYYY-MM-DD
		birthdayStr := strings.TrimSpace(record[1])
		birthDate, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			panic(err)
		}
		birthdayFormatted := fmt.Sprintf("%04d-%02d-%02d", birthDate.Year(), birthDate.Month(), birthDate.Day())

		// Calculate Age as of July 1, 2025
		age := refDate.Year() - birthDate.Year()
		if birthDate.Month() > refDate.Month() || (birthDate.Month() == refDate.Month() && birthDate.Day() > refDate.Day()) {
			age--
		}

		// Parse Relatives (Father, Mother, Brother, Sister columns)
		var relatives []Relative
		for j := 3; j < len(record); j++ {
			relName := strings.TrimSpace(record[j])
			if relName == "" || relName == "null" {
				continue
			}
			relParts := strings.Fields(relName)
			if len(relParts) < 2 {
				continue
			}
			
			var relKey string
			switch j {
			case 3: relKey = "Father"
			case 4: relKey = "Mother"
			case 5: relKey = "Brother"
			case 6: relKey = "Sister"
			}

			if relKey != "" {
				relatives = append(relatives, Relative{
					FirstName:    relParts[0],
					LastName:     relParts[len(relParts)-1],
					Relationship: relKey,
				})
			}
		}

		result = append(result, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		})
	}

	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}