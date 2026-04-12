package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strings"
	"time"
)

type Relative struct {
	FirstName    string
	LastName     string
	Relationship string
}

type Person struct {
	FirstName string
	LastName  string
	Birthday  string
	Age       int
	Relatives []Relative
}

func splitName(full string) (string, string) {
	parts := strings.Fields(full)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[len(parts)-1]
}

func calcAge(birthday time.Time, died time.Time, hasDied bool) int {
	ref := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	if hasDied && died.Before(ref) {
		ref = died
	}
	age := ref.Year() - birthday.Year()
	if ref.YearDay() < birthday.YearDay() {
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

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	var people []Person

	for _, row := range records[1:] {
		name := row[0]
		birthdayStr := row[1]
		diedStr := row[2]
		father := row[3]
		mother := row[4]
		brother := row[5]
		sister := row[6]

		nameParts := strings.Fields(name)
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		birthday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			birthday, err = time.Parse("01/02/2006", birthdayStr)
			if err != nil {
				panic(err)
			}
		}

		var died time.Time
		hasDied := false
		if diedStr != "null" && diedStr != "" {
			died, err = time.Parse("1/2/2006", diedStr)
			if err != nil {
				died, err = time.Parse("01/02/2006", diedStr)
				if err != nil {
					panic(err)
				}
			}
			hasDied = true
		}

		age := calcAge(birthday, died, hasDied)

		var relatives []Relative
		for _, rel := range []struct {
			val  string
			kind string
		}{
			{father, "Father"},
			{mother, "Mother"},
			{brother, "Brother"},
			{sister, "Sister"},
		} {
			if rel.val == "null" || rel.val == "" {
				continue
			}
			fn, ln := splitName(rel.val)
			relatives = append(relatives, Relative{
				FirstName:    fn,
				LastName:     ln,
				Relationship: rel.kind,
			})
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthday.Format("2006-01-02"),
			Age:       age,
			Relatives: relatives,
		})
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(people)
}
