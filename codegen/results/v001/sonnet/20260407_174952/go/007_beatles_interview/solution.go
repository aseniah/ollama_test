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

func splitName(full string) (string, string) {
	parts := strings.SplitN(strings.TrimSpace(full), " ", 2)
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

func parseDate(s string) (time.Time, error) {
	return time.Parse("1/2/2006", s)
}

func calcAge(birthday time.Time, asOf time.Time) int {
	years := asOf.Year() - birthday.Year()
	// Check if birthday hasn't occurred yet in asOf year
	anniversaryThisYear := time.Date(asOf.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)
	if asOf.Before(anniversaryThisYear) {
		years--
	}
	return years
}

func main() {
	f, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Build header index
	headers := records[0]
	idx := make(map[string]int)
	for i, h := range headers {
		idx[h] = i
	}

	asOf := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	relFields := []string{"Father", "Mother", "Brother", "Sister"}

	var people []Person

	for _, row := range records[1:] {
		fullName := row[idx["Name"]]
		firstName, lastName := splitName(fullName)

		birthdayStr := row[idx["Birthday"]]
		birthday, err := parseDate(birthdayStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "bad birthday:", birthdayStr, err)
			os.Exit(1)
		}

		// Format birthday as YYYY-MM-DD
		formattedBirthday := birthday.Format("2006-01-02")

		// Determine age
		diedStr := row[idx["Died"]]
		var age int
		if diedStr == "null" || diedStr == "" {
			age = calcAge(birthday, asOf)
		} else {
			died, err := parseDate(diedStr)
			if err != nil {
				fmt.Fprintln(os.Stderr, "bad died:", diedStr, err)
				os.Exit(1)
			}
			age = calcAge(birthday, died)
		}

		// Build relatives
		var relatives []Relative
		for _, rel := range relFields {
			val := row[idx[rel]]
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
			Birthday:  formattedBirthday,
			Age:       age,
			Relatives: relatives,
		})
	}

	out, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
