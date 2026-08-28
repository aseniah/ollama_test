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
	FirstName   string     `json:"FirstName"`
	LastName    string     `json:"LastName"`
	Birthday    string     `json:"Birthday"`
	Age         int        `json:"Age"`
	Relatives   []Relative `json:"Relatives"`
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
			continue
		}
		if len(record) < 7 {
			continue
		}

		nameParts := strings.Fields(record[0])
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		bday, err := time.Parse("1/2/2006", strings.TrimSpace(record[1]))
		if err != nil {
			panic(err)
		}
		bdayFormatted := bday.Format("2006-01-02")

		age := refDate.Year() - bday.Year()
		if refDate.Month() < bday.Month() || (refDate.Month() == bday.Month() && refDate.Day() < bday.Day()) {
			age--
		}

		relatives := make([]Relative, 0)
		relPairs := []struct {
			Rel string
			Idx int
		}{
			{"Father", 3},
			{"Mother", 4},
			{"Brother", 5},
			{"Sister", 6},
		}

		for _, rp := range relPairs {
			val := strings.TrimSpace(record[rp.Idx])
			if val == "null" {
				continue
			}
			relParts := strings.Fields(val)
			if len(relParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    relParts[0],
					LastName:     relParts[len(relParts)-1],
					Relationship: rp.Rel,
				})
			}
		}

		people = append(people, Person{
			FirstName:   firstName,
			LastName:    lastName,
			Birthday:    bdayFormatted,
			Age:         age,
			Relatives:   relatives,
		})
	}

	jsonBytes, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(jsonBytes))
}