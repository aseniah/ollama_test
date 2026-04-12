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

func calcAge(bday time.Time, died time.Time, hasDied bool) int {
	ref := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	if hasDied && died.Before(ref) {
		ref = died
	}
	age := ref.Year() - bday.Year()
	if ref.Month() < bday.Month() || (ref.Month() == bday.Month() && ref.Day() < bday.Day()) {
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

	headers := records[0]
	colIndex := make(map[string]int)
	for i, h := range headers {
		colIndex[h] = i
	}

	relCols := []string{"Father", "Mother", "Brother", "Sister"}

	var people []Person

	for _, row := range records[1:] {
		fullName := row[colIndex["Name"]]
		firstName, lastName := splitName(fullName)

		bdayStr := row[colIndex["Birthday"]]
		bday, err := time.Parse("1/2/2006", bdayStr)
		if err != nil {
			panic(err)
		}

		diedStr := row[colIndex["Died"]]
		var diedTime time.Time
		hasDied := false
		if diedStr != "null" && diedStr != "" {
			diedTime, err = time.Parse("1/2/2006", diedStr)
			if err != nil {
				panic(err)
			}
			hasDied = true
		}

		age := calcAge(bday, diedTime, hasDied)

		var relatives []Relative
		for _, rel := range relCols {
			val := row[colIndex[rel]]
			if val == "null" || val == "" {
				continue
			}
			fn, ln := splitName(val)
			relatives = append(relatives, Relative{
				FirstName:    fn,
				LastName:     ln,
				Relationship: rel,
			})
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
