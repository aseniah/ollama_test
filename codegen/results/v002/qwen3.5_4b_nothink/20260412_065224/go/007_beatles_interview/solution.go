package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Relative struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName  string  `json:"FirstName"`
	LastName   string  `json:"LastName"`
	Birthday   string  `json:"Birthday"`
	Age        int     `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

func parseDate(dateStr string) time.Time {
	parts := strings.Split(dateStr, "/")
	if len(parts) != 3 {
		return time.Time{}
	}
	d, _ := time.Parse("01/02/06", dateStr)
	return d
}

func calculateAge(birthday time.Time, referenceTime time.Time) int {
	diff := referenceTime.Sub(birthday)
	years := int(diff / timeYear)
	days := diff.Sub(time.Duration(timeYear * time.Duration(years)))
	daysInYear := timeYear / time.Duration(years) + time.Second
	if days >= daysInYear {
		years++
	}
	return years
}

func main() {
	referenceTime, _ := time.Parse("2006-01-02", "2025-07-01")
	var people []Person

	// Read input CSV
	data, err := os.ReadFile("input/input.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content := strings.TrimSpace(string(data))
	lines := strings.Split(content, "\n")
	if len(lines) < 2 {
		fmt.Println("Error: Not enough lines in input file")
		os.Exit(1)
	}

	// Parse CSV
	header := strings.Split(lines[0], ",")
	colIdx := make(map[string]int)
	for i, h := range header {
		colIdx[h] = i
	}

	for i := 1; i < len(lines); i++ {
		fields := strings.Split(lines[i], ",")
		name := strings.Split(fields[colIdx["Name"]], " ")
		if len(name) != 2 {
			continue
		}
		firstName := name[0]
		lastName := name[1]

		birthdayStr := fields[colIdx["Birthday"]]
		birthday := parseDate(birthdayStr)

		age := calculateAge(birthday, referenceTime)

		fatherName := fields[colIdx["Father"]]
		motherName := fields[colIdx["Mother"]]
		brotherName := fields[colIdx["Brother"]]
		sisterName := fields[colIdx["Sister"]]

		var relatives []Relative
		if fatherName != "null" {
			relatives = append(relatives, Relative{
				FirstName:  fatherName,
				LastName:   fatherName,
				Relationship: "Father",
			})
		}
		if motherName != "null" {
			relatives = append(relatives, Relative{
				FirstName:  motherName,
				LastName:   motherName,
				Relationship: "Mother",
			})
		}
		if brotherName != "null" {
			relatives = append(relatives, Relative{
				FirstName:  brotherName,
				LastName:   brotherName,
				Relationship: "Brother",
			})
		}
		if sisterName != "null" {
			relatives = append(relatives, Relative{
				FirstName:  sisterName,
				LastName:   sisterName,
				Relationship: "Sister",
			})
		}

		people = append(people, Person{
			FirstName:  firstName,
			LastName:   lastName,
			Birthday:   birthdayStr,
			Age:        age,
			Relatives:  relatives,
		})
	}

	// Generate JSON
	jsonBytes, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonBytes))
}

const timeYear = time.Hour * 24 * 365.25