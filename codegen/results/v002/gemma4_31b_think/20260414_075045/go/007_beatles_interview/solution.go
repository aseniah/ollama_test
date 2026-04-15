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

	// Header: Name, Birthday, Died, Father, Mother, Brother, Sister
	header := records[0]
	data := records[1:]

	var people []Person
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	for _, row := range data {
		rowMap := make(map[string]string)
		for i, val := range row {
			rowMap[header[i]] = val
		}

		nameParts := strings.Fields(rowMap["Name"])
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		birthStr := rowMap["Birthday"]
		birthDate, err := time.Parse("1/2/2006", birthStr)
		if err != nil {
			continue
		}

		// Determine target date for age calculation (Died date or July 1, 2025)
		targetDate := refDate
		diedStr := rowMap["Died"]
		if diedStr != "" && diedStr != "null" {
			if d, err := time.Parse("1/2/2006", diedStr); err == nil {
				targetDate = d
			}
		}

		age := calculateAge(birthDate, targetDate)

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthDate.Format("2006-01-02"),
			Age:       age,
			Relatives: []Relative{},
		}

		// Process relatives
		relativeCols := []string{"Father", "Mother", "Brother", "Sister"}
		for _, col := range relativeCols {
			relName := rowMap[col]
			if relName == "" || relName == "null" {
				continue
			}

			relParts := strings.Fields(relName)
			if len(relParts) > 0 {
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    relParts[0],
					LastName:     relParts[len(relParts)-1],
					Relationship: col,
				})
			}
		}

		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(output))
}

func calculateAge(birth, target time.Time) int {
	age := target.Year() - birth.Year()
	// Adjust if the birthday hasn't occurred yet in the target year
	if target.Month() < birth.Month() || (target.Month() == birth.Month() && target.Day() < birth.Day()) {
		age--
	}
	return age
}