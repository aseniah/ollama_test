package main

import (
	"encoding/csv"
	"encoding/json"
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

func splitName(full string) (first, last string) {
	parts := strings.Fields(full)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], parts[0]
	}
	return parts[0], parts[len(parts)-1]
}

func calcAge(birthday time.Time, died time.Time, asOf time.Time) int {
	end := asOf
	if !died.IsZero() {
		end = died
	}
	years := end.Year() - birthday.Year()
	// subtract 1 if birthday hasn't occurred yet in the end year
	bDay := time.Date(end.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)
	if end.Before(bDay) {
		years--
	}
	return years
}

func main() {
	f, err := os.Open("input/input.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	asOf := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	layout := "1/2/2006"
	isoLayout := "2006-01-02"

	// header: Name,Birthday,Died,Father,Mother,Brother,Sister
	header := records[0]
	colIndex := map[string]int{}
	for i, h := range header {
		colIndex[h] = i
	}

	relCols := []string{"Father", "Mother", "Brother", "Sister"}

	var people []Person

	for _, row := range records[1:] {
		name := row[colIndex["Name"]]
		firstName, lastName := splitName(name)

		birthdayRaw := row[colIndex["Birthday"]]
		birthday, err := time.Parse(layout, birthdayRaw)
		if err != nil {
			panic(err)
		}

		diedRaw := row[colIndex["Died"]]
		var died time.Time
		if diedRaw != "" && diedRaw != "null" {
			died, err = time.Parse(layout, diedRaw)
			if err != nil {
				panic(err)
			}
		}

		age := calcAge(birthday, died, asOf)

		var relatives []Relative
		for _, rel := range relCols {
			idx, ok := colIndex[rel]
			if !ok {
				continue
			}
			val := row[idx]
			if val == "" || val == "null" {
				continue
			}
			rf, rl := splitName(val)
			relatives = append(relatives, Relative{
				FirstName:    rf,
				LastName:     rl,
				Relationship: rel,
			})
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthday.Format(isoLayout),
			Age:       age,
			Relatives: relatives,
		})
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(people)
}
