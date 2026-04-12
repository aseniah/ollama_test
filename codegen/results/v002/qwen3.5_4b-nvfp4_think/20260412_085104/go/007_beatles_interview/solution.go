package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Member struct {
	FirstName string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Birthday   string `json:"Birthday"`
	Age        int    `json:"Age"`
	Relatives  []RelativeEntry `json:"Relatives"`
}

type RelativeEntry struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
		os.Exit(1)
	}
	defer file.Close()

	var records []map[string]string
	lines, _ := readAllLines(file)

	// Skip header
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" || line == "Name,Birthday,Died,Father,Mother,Brother,Sister" {
			continue
		}
		fields := strings.Split(line, ",")
		records = append(records, map[string]string{
			"FullName":  fields[0],
			"Birthday":  fields[1],
			"Died":      fields[2],
			"Father":    fields[3],
			"Mother":    fields[4],
			"Brother":   fields[5],
			"Sister":    fields[6],
		})
	}

	var members []Member
	targetDate := int64(1920250701) // July 1, 2025 as int64 timestamp

	for _, rec := range records {
		member := Member{}

		// Parse first and last name
		fields := strings.Split(rec["FullName"], " ")
		if len(fields) >= 2 {
			member.FirstName = fields[0]
			member.LastName = fields[1]
		} else if len(fields) == 1 {
			member.FirstName = fields[0]
			member.LastName = ""
		}

		// Parse birthday and convert to YYYY-MM-DD
		bf := strings.Fields(rec["Birthday"])
		if len(bf) >= 3 {
			var month, day, year int
			for i := 0; i < 3; i++ {
				val, _ := strconv.Atoi(bf[i])
				if i == 0 {
					month = val
				} else if i == 1 {
					day = val
				} else {
					year = val
				}
			}
			member.Birthday = fmt.Sprintf("%d-%02d-%02d", year, month, day)

			// Calculate age as of July 1, 2025
			bday := int64(year*10000 + month*100 + day)
			tbdy := int64(25 * 10000 + 7 * 100 + 1) // July 1, 2025

			if tbdy < bday {
				age = year - 1
			} else if tbdy == bday {
				age = year
			} else {
				age = year - 1
			}
			member.Age = int(math.Round(float64(age)))
		}

		// Build relatives array
		if rec["Father"] != "" && strings.TrimSpace(rec["Father"]) != "null" {
			f := strings.Split(rec["Father"], ",")
			memberRelatives = append(memberRelatives, RelativeEntry{
				FirstName:  f[0],
				LastName:   f[1],
				Relationship: "Father",
			})
		}

		if rec["Mother"] != "" && strings.TrimSpace(rec["Mother"]) != "null" {
			m := strings.Split(rec["Mother"], ",")
			memberRelatives = append(memberRelatives, RelativeEntry{
				FirstName:  m[0],
				LastName:   m[1],
				Relationship: "Mother",
			})
		}

		if rec["Brother"] != "" && strings.TrimSpace(rec["Brother"]) != "null" {
			b := strings.Split(rec["Brother"], ",")
			memberRelatives = append(memberRelatives, RelativeEntry{
				FirstName:  b[0],
				LastName:   b[1],
				Relationship: "Brother",
			})
		}

		if rec["Sister"] != "" && strings.TrimSpace(rec["Sister"]) != "null" {
			s := strings.Split(rec["Sister"], ",")
			memberRelatives = append(memberRelatives, RelativeEntry{
				FirstName:  s[0],
				LastName:   s[1],
				Relationship: "Sister",
			})
		}

		members = append(members, member)
	}

	jsonData, _ := json.MarshalIndent(members, "", "   ")
	fmt.Println(string(jsonData))
}