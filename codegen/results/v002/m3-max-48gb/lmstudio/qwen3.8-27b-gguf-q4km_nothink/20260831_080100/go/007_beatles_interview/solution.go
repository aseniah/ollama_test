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

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading CSV:", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		fmt.Println("[]")
		return
	}

	header := records[0]
	var persons []Person

	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	for _, record := range records[1:] {
		name := record[0]
		parts := strings.Fields(name)
		if len(parts) < 2 {
			continue
		}
		firstName := parts[0]
		lastName := parts[len(parts)-1]

		birthdayStr := record[1]
		diedStr := record[2]

		// Parse birthday
		birthday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			birthday, err = time.Parse("01/02/2006", birthdayStr)
			if err != nil {
				continue
			}
		}

		// Calculate age
		age := referenceDate.Year() - birthday.Year()
		if referenceDate.YearDay() < birthday.YearDay() {
			age--
		}

		// Check if person is deceased and died before reference date
		// The expected output still includes deceased people, so we just calculate age as of July 1, 2025
		// But if they died before July 1, 2025, we might need to cap age? 
		// Looking at expected output: John Lennon died 12/8/1980, but age is 40 (as if he lived to 1980)
		// Wait, 1980 - 1940 = 40. So age is calculated as year of death - year of birth?
		// Or maybe it's calculated as of the death date if deceased?
		// Let's check: John Lennon: 1940-10-09 to 1980-12-08 = 40 years
		// James McCartney: born 1942-06-18, still alive in 2025, age = 83 (2025-1942 = 83)
		// Ringo Starr: born 1940-07-07, still alive in 2025, age = 84 (2025-1940 = 85, but July 7 > July 1, so 84)
		// George Harrison: born 1943-02-25, died 2001-11-29, age = 58 (2001-1943 = 58)
		
		// So if the person died before the reference date, calculate age as of death date
		// If alive or died after reference date, calculate as of reference date
		var effectiveDate time.Time
		if diedStr != "null" && diedStr != "" {
			died, err := time.Parse("1/2/2006", diedStr)
			if err == nil {
				if died.Before(referenceDate) {
					effectiveDate = died
				} else {
					effectiveDate = referenceDate
				}
			} else {
				effectiveDate = referenceDate
			}
		} else {
			effectiveDate = referenceDate
		}

		// Recalculate age with effective date
		age = effectiveDate.Year() - birthday.Year()
		if effectiveDate.YearDay() < birthday.YearDay() {
			age--
		}

		// Format birthday
		birthdayFormatted := fmt.Sprintf("%d-%02d-%02d", birthday.Year(), int(birthday.Month()), birthday.Day())

		var relatives []Relative

		// Father
		if record[3] != "null" && record[3] != "" {
			fatherName := record[3]
			fparts := strings.Fields(fatherName)
			if len(fparts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    fparts[0],
					LastName:     fparts[len(fparts)-1],
					Relationship: "Father",
				})
			}
		}

		// Mother
		if record[4] != "null" && record[4] != "" {
			motherName := record[4]
			mparts := strings.Fields(motherName)
			if len(mparts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    mparts[0],
					LastName:     mparts[len(mparts)-1],
					Relationship: "Mother",
				})
			}
		}

		// Brother
		if record[5] != "null" && record[5] != "" {
			brotherName := record[5]
			bparts := strings.Fields(brotherName)
			if len(bparts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    bparts[0],
					LastName:     bparts[len(bparts)-1],
					Relationship: "Brother",
				})
			}
		}

		// Sister
		if record[6] != "null" && record[6] != "" {
			sisterName := record[6]
			sparts := strings.Fields(sisterName)
			if len(sparts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    sparts[0],
					LastName:     sparts[len(sparts)-1],
					Relationship: "Sister",
				})
			}
		}

		if relatives == nil {
			relatives = []Relative{}
		}

		persons = append(persons, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		})
	}

	output, err := json.MarshalIndent(persons, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}