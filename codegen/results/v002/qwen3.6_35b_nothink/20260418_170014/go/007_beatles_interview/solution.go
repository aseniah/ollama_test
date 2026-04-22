package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func calculateAge(birthdayStr string, asOf time.Time) int {
	birthday, err := time.Parse("1/2/2006", birthdayStr)
	if err != nil {
		return 0
	}
	age := asOf.Year() - birthday.Year()
	// Check if birthday has occurred this year
	if asOf.YearDay() < birthday.YearDay() {
		age--
	}
	return age
}

func main() {
	// Read the CSV file
	data, err := ioutil.ReadFile("input/input.csv")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(strings.NewReader(string(data)))
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Skip header
	if len(records) < 2 {
		panic("No data rows in CSV")
	}

	// Reference date for age calculation
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	for _, record := range records[1:] {
		if len(record) < 7 {
			continue
		}

		name := record[0]
		nameParts := strings.SplitN(name, " ", 2)
		firstName := nameParts[0]
		lastName := ""
		if len(nameParts) > 1 {
			lastName = nameParts[1]
		}

		birthday := record[1]
		// died := record[2] // Not used

		father := record[3]
		mother := record[4]
		brother := record[5]
		sister := record[6]

		age := calculateAge(birthday, refDate)

		var relatives []Relative

		if father != "null" && father != "" {
			fParts := strings.SplitN(father, " ", 2)
			fFirst := fParts[0]
			fLast := ""
			if len(fParts) > 1 {
				fLast = fParts[1]
			}
			relatives = append(relatives, Relative{
				FirstName:    fFirst,
				LastName:     fLast,
				Relationship: "Father",
			})
		}

		if mother != "null" && mother != "" {
			mParts := strings.SplitN(mother, " ", 2)
			mFirst := mParts[0]
			mLast := ""
			if len(mParts) > 1 {
				mLast = mParts[1]
			}
			relatives = append(relatives, Relative{
				FirstName:    mFirst,
				LastName:     mLast,
				Relationship: "Mother",
			})
		}

		if brother != "null" && brother != "" {
			bParts := strings.SplitN(brother, " ", 2)
			bFirst := bParts[0]
			bLast := ""
			if len(bParts) > 1 {
				bLast = bParts[1]
			}
			relatives = append(relatives, Relative{
				FirstName:    bFirst,
				LastName:     bLast,
				Relationship: "Brother",
			})
		}

		if sister != "null" && sister != "" {
			sParts := strings.SplitN(sister, " ", 2)
			sFirst := sParts[0]
			sLast := ""
			if len(sParts) > 1 {
				sLast = sParts[1]
			}
			relatives = append(relatives, Relative{
				FirstName:    sFirst,
				LastName:     sLast,
				Relationship: "Sister",
			})
		}

		// Format birthday as YYYY-MM-DD
		bd, err := time.Parse("1/2/2006", birthday)
		if err != nil {
			panic(err)
		}
		birthdayFormatted := bd.Format("2006-01-02")

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		})
	}

	// Output as JSON
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}