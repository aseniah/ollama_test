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
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var people []Person
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	for i, record := range records {
		if i == 0 {
			continue // skip header
		}
		if len(record) < 7 {
			continue
		}

		name := strings.TrimSpace(record[0])
		nameParts := strings.Fields(name)
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		birthdayStr := strings.TrimSpace(record[1])
		bday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			panic(err)
		}

		age := refDate.Year() - bday.Year()
		if bday.Month() > refDate.Month() || (bday.Month() == refDate.Month() && bday.Day() > refDate.Day()) {
			age--
		}

		bdayFormatted := bday.Format("2006-01-02")

		var relatives []Relative
		cols := []string{"Father", "Mother", "Brother", "Sister"}
		for j, col := range cols {
			val := strings.TrimSpace(record[2+j])
			if val == "" || strings.EqualFold(val, "null") {
				continue
			}
			parts := strings.Fields(val)
			if len(parts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    parts[0],
					LastName:     parts[len(parts)-1],
					Relationship: col,
				})
			}
		}
		if relatives == nil {
			relatives = []Relative{}
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  bdayFormatted,
			Age:       age,
			Relatives: relatives,
		})
	}

	out, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}