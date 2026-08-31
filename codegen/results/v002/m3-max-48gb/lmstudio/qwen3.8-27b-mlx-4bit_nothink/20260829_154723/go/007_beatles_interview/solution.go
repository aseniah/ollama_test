package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func ageAsOf(birthday time.Time, asOf time.Time) int {
	age := asOf.Year() - birthday.Year()
	anniversary := time.Date(asOf.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, birthday.Location())
	if asOf.Before(anniversary) {
		age--
	}
	return age
}

func main() {
	f, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	asOf := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var output []Person

	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) < 7 {
			continue
		}

		name := record[0]
		parts := strings.Fields(name)
		firstName := parts[0]
		lastName := ""
		if len(parts) >= 2 {
			lastName = parts[len(parts)-1]
		}

		birthdayStr := strings.TrimSpace(record[1])
		birthday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		age := ageAsOf(birthday, asOf)
		birthdayISO := birthday.Format("2006-01-02")

		relatives := []Relative{}

		addRelative := func(csvVal, relationship string) {
			val := strings.TrimSpace(csvVal)
			if val == "" || val == "null" {
				return
			}
			p := strings.Fields(val)
			fn := p[0]
			ln := ""
			if len(p) >= 2 {
				ln = p[len(p)-1]
			}
			relatives = append(relatives, Relative{
				FirstName:    fn,
				LastName:     ln,
				Relationship: relationship,
			})
		}

		addRelative(record[3], "Father")
		addRelative(record[4], "Mother")
		addRelative(record[5], "Brother")
		addRelative(record[6], "Sister")

		if relatives == nil {
			relatives = []Relative{}
		}

		output = append(output, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayISO,
			Age:       age,
			Relatives: relatives,
		})
	}

	_ = strconv.Itoa

	out, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}