```go
package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName  string    `json:"FirstName"`
	LastName   string    `json:"LastName"`
	Birthday   string    `json:"Birthday"`
	Age        int       `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

func calculateAge(birthDate string, referenceDate time.Time) int {
	birth, err := time.Parse("1/2/2006", birthDate)
	if err != nil {
		return 0
	}

	years := referenceDate.Year() - birth.Year()
	
	// Adjust if birthday hasn't occurred yet this year
	if referenceDate.Month() < birth.Month() || 
		(referenceDate.Month() == birth.Month() && referenceDate.Day() < birth.Day()) {
		years--
	}
	
	return years
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Create reference date: July 1, 2025
	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	// Skip header row
	for i, record := range records {
		if i == 0 {
			continue
		}

		// Parse name to get first and last name
		fullName := record[0]
		nameParts := []string{}
		for _, part := range fullName {
			if part == ' ' {
				nameParts = append(nameParts, " ")
			} else {
				nameParts[len(nameParts)-1] = nameParts[len(nameParts)-1] + string(part)
			}
		}

		var firstName, lastName string
		firstName = nameParts[0]
		if len(nameParts) > 1 {
			lastName = nameParts[1]
		} else {
			lastName = ""
		}

		// Parse birthday (convert format)
		birthday := record[1]
		parts := []string{}
		for i, part := range birthday {
			if part == '/' {
				parts = append(parts, " ")
			} else {
				if len(parts) == 0 {
					parts = append(parts, string(part))
				} else {
					parts[len(parts)-1] = parts[len(parts)-1] + string(part)
				}
			}
		}
		
		// Convert to format: 1/2/1940 -> 1/2/1940 (keeping original for now)
		// Actually parse correctly:
		var month, day, year string
		monthPart := parts[0]
		dayPart := parts[1]
		yearPart := parts[2]
		
		// Format like M/D/YYYY (e.g., 10/9/1940 or 6/18/1942)
		// We'll handle parsing carefully to match our output expectation
		parts = []string{}
		for i, part := range birthday {
			if part == '/' {
				parts = append(parts, " ")
			} else {
				if len(parts) == 0 {
					parts = append(parts, string(part))
				} else {
					parts[len(parts)-1] = parts[len(parts)-1] + string(part)
				}
			}
		}
		if len(parts) == 3 {
			// Ensure right padding of single digits
			bMonth := parts[0]
			bDay := parts[1]
			bYear := parts[2]
			
			// Format: M/D/YYYY (but keep format as needed)
			// Need to check if we have "10/9/1940" or "6/18/1942"
			// Keep existing format for now but ensure correct parsing for date calculations
			bDay = fmt.Sprintf("%d", parseNumber(bDay))
			bMonth = fmt.Sprintf("%d", parseNumber(bMonth))
			bYear = fmt.Sprintf("%d", parseNumber(bYear))
			
			// Format to proper MM/DD/YYYY or similar - but the expected output shows same format?
			// Let's check the CSV and the expected:
			// CSV: "10/9/1940" -> Expected: "1940-10-09" 
			// Format conversion is required - we want YYYY-MM-DD in JSON but CSV shows M/D/YYYY
			
			// Let's just do date parsing as is now, then convert to desired format later
			// Use a more robust method:
			// Original: "10/9/1940" -> Need to parse as date, then output "1940-10-09"
			// We know the actual date strings are in M/D/YYYY (like 10/9/1940) so we need to parse correctly.
			
			// Convert MM/DD/YYYY string to a standard string we can parse later
			// For example "10/9/1940" -> Use Go time parser correctly
			// Since the date string format in CSV is M/D/YYYY like 10/9/1940, 
			// we can parse it using time.Parse with specific format.
			
			// Use Go's built-in parser for M/D/YYYY, converting to YYYY-MM-DD before output
			
			// For simplicity, let's first parse what we can:
			// Format: M/D/YYYY
			
			// But actually let's do it more carefully to be accurate:
			// First, parse the parts carefully as numbers.
			monthPart = parts[0]
			dayPart = parts[1] 
			yearPart = parts[2]

			// Ensure padding to get consistent format like 01/02/1940 (but not necessary)
			// Just parse correctly. Let's do a direct conversion:
			
			// If we want 1940-10-09 as output, we just need to convert string format.
			// Go's parser can understand 1/2/2006, and we need to output in the right format.
			
			bMonth = strings.TrimLeft(parts[0], "0")
			bDay = strings.TrimLeft(parts[1], "0")
			bYear = parts[2]
			
			// But the issue is the CSV format "10/9/1940" when we're converting it,
			// Go's time parser can't parse that easily without formatting.
			
			// Simplest approach:
			// Take "10/9/1940"
			// Go time.Parse("1/2/2006", ...) 
			// That should work...
			
			// Actually, Go's time.Parse("1/2/2006", "10/9/1940") will parse it.
			// So just use that to determine age, then return in required format.
			
			// Let's just make sure date formatting is compatible with the standard Go format parsing
			// to create accurate age, then we'll reformat appropriately for JSON output.
			
			// Let's just make a simple approach. For now, the age calculation is what matters.
			// Reformat properly for JSON output.

			// Try to parse correctly:
			birthDateRaw := fmt.Sprintf("%s/%s/%s", monthPart, dayPart, yearPart)
			parsedBirth, err := time.Parse("1/2/2006", birthDateRaw)
			if err != nil {
				// Let's fallback to just using the original string for age calculation, 
				// but use the correct format for output.
				// If it fails, use a simpler approach:
				
				bMonth = strings.TrimLeft(parts[0], "0")
				bDay = strings.TrimLeft(parts[1], "0") 
				bYear = parts[2]
				
				// Now for JSON output, we want "1940-10-09"
				// Let's just use the original string format, but convert it
				// to the standard date format with consistent two-digit months and days.
				if len(bMonth) == 1 {
					bMonth = "0" + bMonth
				}
				if len(bDay) == 1 {
					bDay = "0" + bDay
				}
				// Output format: YYYY-MM-DD
				birthdayFormatted := fmt.Sprintf("%s-%s-%s", bYear, bMonth, bDay)
				// Now for age calculation...
				
				// To get the age, we need to parse correctly. Let's make an actual date from the parts:
				birthDate, err := time.Parse("1/2/2006", fmt.Sprintf("%s/%s/%s", parts[0], parts[1], parts[2]))
				if err != nil {
					log.Fatal("Failed to parse date for person:", fullName)
				}
				
				// Calculate age correctly based on July 1, 2025 as reference
				age := calculateAge(fmt.Sprintf("%s/%s/%s", parts[0], parts[1], parts[2]), referenceDate)
				
				// Create the JSON object with the right format...
				
				// Reconstruct birthday string with leading zeros if necessary:
				bMonth = parts[0]
				if len(bMonth) == 1 {
					bMonth = "0" + bMonth
				}
				bDay = parts[1] 
				if len(bDay) == 1 {
					bDay = "0" + bDay
				}
				bYear = parts[2]
				
				// Create final birthday string in desired format for output:
				birthdayOutput := fmt.Sprintf("%s-%s-%s", bYear, bMonth, bDay)
				
				// Build relatives list:
				relatives := make([]Relative, 0)
				
				// Get father name if exists:
				father := record[3]
				if father != "null" && father != "" {
					fatherParts := strings.Split(father, " ")
					if len(fatherParts) >= 2 {
						relative := Relative{
							FirstName:    fatherParts[0],
							LastName:     fatherParts[1],
							Relationship: "Father",
						}
						relatives = append(relatives, relative)
					} else if len(fatherParts) == 1 {
						relative := Relative{
							FirstName:    fatherParts[0],
							LastName:     "",
							Relationship: "Father",
						}
						relatives = append(relatives, relative)
					}
				}
				
				// Get mother name if exists:
				mother := record[4]
				if mother != "null" && mother != "" {
					motherParts := strings.Split(mother, " ")
					if len(motherParts) >= 2 {
						relative := Relative{
							FirstName:    motherParts[0],
							LastName:     motherParts[1],
							Relationship: "Mother",
						}
						relatives = append(relatives, relative)
					} else if len(motherParts) == 1 {
						relative := Relative{
							FirstName:    motherParts[0],
							LastName:     "",
							Relationship: "Mother",
						}
						relatives = append(relatives, relative)
					}
				}
				
				// Get brother name if exists:
				brother := record[5]
				if brother != "null" && brother != "" {
					brotherParts := strings.Split(brother, " ")
					if len(brotherParts) >= 2 {
						relative := Relative{
							FirstName:    brotherParts[0],
							LastName:     brotherParts[1],
							Relationship: "Brother",
						}
						relatives = append(relatives, relative)
					} else if len(brotherParts) == 1 {
						relative := Relative{
							FirstName:    brotherParts[0],
							LastName:     "",
							Relationship: "Brother",
						}
						relatives = append(relatives, relative)
					}
				}
				
				// Get sister name if exists:
				sister := record[6]
				if sister != "null" && sister != "" {
					sisterParts := strings.Split(sister, " ")
					if len(sisterParts) >= 2 {
						relative := Relative{
							FirstName:    sisterParts[0],
							LastName:     sisterParts[1],
							Relationship: "Sister",
						}
						relatives = append(relatives, relative)
					} else if len(sisterParts) == 1 {
						relative := Relative{
							FirstName:    sisterParts[0],
							LastName:     "",
							Relationship: "Sister",
						}
						relatives = append(relatives, relative)
					}
				}
				
				// Add to results:
				person := Person{
					FirstName:  firstName,
					LastName:   lastName,
					Birthday:   birthdayOutput,
					Age:        age,
					Relatives:  relatives,
				}
				people = append(people, person)
				
			} else {
				// Successful parse
				// Do age calculation...
				age := calculateAge(fmt.Sprintf("%s/%s/%s", parts[0], parts[1], parts[2]), referenceDate)
				// Reformat the birthday string for JSON:
				bMonth = parts[0]
				if len(bMonth) == 1 {
					bMonth = "0" + bMonth
				}
				bDay = parts[1]
				if len(bDay) == 1 {
					bDay = "0" + bDay
				}
				bYear = parts[2]
				birthdayOutput := fmt.Sprintf("%s-%s-%s", bYear, bMonth, bDay)
				
				// Build relatives list:
				relatives := make([]Relative, 0)
				
				// Get father name if exists:
				father := record[3]
				if father != "null" && father != "" {
					fatherParts := strings.Split(father, " ")
					if len(fatherParts) >= 2 {
						relative := Relative{
							FirstName:    fatherParts[0],
							LastName:     fatherParts[1],
							Relationship: "Father",
						}
						relatives = append(relatives, relative)
					} else if len(fatherParts) == 1 {
						relative := Relative{
							FirstName:    fatherParts[0],
							LastName:     "",
							Relationship: "Father",
						}
						relatives = append(relatives, relative)
					}
				}
				
				// Get mother name if exists:
				mother := record[4]
				if mother != "null" && mother != "" {
					motherParts := strings.Split(mother, " ")
					if len(motherParts) >= 2 {
						relative := Relative{
							FirstName:    motherParts[0],
							LastName:     motherParts[1],
							Relationship: "Mother",
						}
						relatives = append(relatives, relative)
					} else if len(motherParts) == 1 {
						relative := Relative{
							FirstName:    motherParts[0],
							LastName:     "",
							Relationship: "Mother",
						}
						relatives = append(relatives, relative)
					}
				}
				
				// Get brother name if exists:
				brother := record[5]
				if brother != "null" && brother != "" {
					brotherParts := strings.Split(brother, " ")
					if len(brotherParts) >= 2 {
						relative := Relative{
							FirstName:    brotherParts[0],
							LastName:     brotherParts[1],
							Relationship: "Brother",
						}
						relatives = append(relatives, relative)
					} else if len(brotherParts) == 1 {
						relative := Relative{
							FirstName:    brotherParts[0],
							LastName:     "",
							Relationship: "Brother",
						}
						relatives = append(relatives, relative)
					}
				}
				
				// Get sister name if exists:
				sister := record[6]
				if sister != "null" && sister != "" {
					sisterParts := strings.Split(sister, " ")
					if len(sisterParts) >= 2 {
						relative := Relative{
							FirstName:    sisterParts[0],
							LastName:     sisterParts[1],
							Relationship: "Sister",
						}
						relatives = append(relatives, relative)
					} else if len(sisterParts) == 1 {
						relative := Relative{
							FirstName:    sisterParts[0],
							LastName:     "",
							Relationship: "Sister",
						}
						relatives = append(relatives, relative)
					}
				}
				
				// Add to results:
				person := Person{
					FirstName:  firstName,
					LastName:   lastName,
					Birthday:   birthdayOutput,
					Age:        age,
					Relatives:  relatives,
				}
				people = append(people, person)
			}
			
		} else {
			// Handle case where there's an unexpected format.
			// For simplicity, we could just skip the record, but let's make an assumption:
			// We'll assume the record format is always the same pattern.
			// But let's go back to just reading it and processing it:
			// Try again with another structure...
			// For brevity, just read and try to parse date properly...
			
			// The most critical part is to parse the date from "10/9/1940" 
			// and calculate age correctly on "2025-07-01"
			// We'll parse it manually and use the standard Go time parsing approach.
			
			// If the parsing fails due to a problem with format handling,
			// we can just skip and continue...
			
			// Let's take a step back...
			
			// Read CSV record: