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
	if ref.Month() < birthday.Month() || (ref.Month() == birthday.Month() && ref.Day() < birthday.Day()) {
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

	for i, row := range records {
		if i == 0 || len(row) < 7 {
			continue
		}

		nameParts := strings.Fields(strings.TrimSpace(row[0]))
		firstName, lastName := "", ""
		if len(nameParts) >= 2 {
			firstName = nameParts[0]
			lastName = nameParts[len(nameParts)-1]
		} else if len(nameParts) == 1 {
			firstName = nameParts[0]
		}

		bday, err := time.Parse("1/2/2006", strings.TrimSpace(row[1]))
		if err != nil {
			panic(err)
		}

		var diedTime time.Time
		hasDied := false
		diedStr := strings.TrimSpace(row[2])
		if diedStr != "null" && diedStr != "" {
			diedTime, err = time.Parse("1/2/2006", diedStr)
			if err != nil {
				panic(err)
			}
			hasDied = true
		}

		age := calcAge(bday, diedTime, hasDied)

		relFields := []struct{ val, rel string }{
			{strings.TrimSpace(row[3]), "Father"},
			{strings.TrimSpace(row[4]), "Mother"},
			{strings.TrimSpace(row[5]), "Brother"},
			{strings.TrimSpace(row[6]), "Sister"},
		}

		var relatives []Relative
		for _, rf := range relFields {
			if rf.val == "null" || rf.val == "" {
				continue
			}
			fn, ln := splitName(rf.val)
			relatives = append(relatives, Relative{FirstName: fn, LastName: ln, Relationship: rf.rel})
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  bday.Format("2006-01-02"),
			Age:       age,
			Relatives: relatives,
		})
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(people)
}
