using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;

// Target date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

string csvPath = "input/input.csv";

if (!File.Exists(csvPath))
{
    return;
}

string[] lines = File.ReadAllLines(csvPath);
if (lines.Length <= 1)
{
    Console.WriteLine("[]");
    return;
}

// Parse CSV Header
string[] headers = lines[0].Split(',');

var results = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] values = line.Split(',');
    
    // Map headers to values for easy lookup
    var rowMap = new Dictionary<string, string>();
    for (int j = 0; j < headers.Length; j++)
    {
        rowMap[headers[j]] = j < values.Length ? values[j] : null;
    }

    // Full Name processing
    // The CSV contains "John Winston Lennon", but expected format has "John" and "Lennon"
    // We need to split the full name into First and Last.
    // Based on the expected format, it seems we take the first part as FirstName and last part as LastName.
    string fullName = rowMap["Name"];
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameProfileLastName(nameParts);

    // Birthday processing
    string bdayStr = rowMap["Birthday"];
    DateTime birthday = DateTime.ParseExact(bdayStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    
    // Age calculation as of July 1, 2025
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate < birthday.AddYears(age)) age--;

    // Relatives processing
    var relatives = new List<object>();
    string[] relationColumns = { "Father", "Mother", "Brother", "Sister" };
    string[] relationLabels = { "Father", "Mother", "Brother", "Sister" };

    for (int r = 0; r < relationColumns.Length; r++)
    {
        string col = relationColumns[r];
        string relVal = rowMap[col];

        if (!string.IsNullOrEmpty(relVal) && relVal.ToLower() != "null")
        {
            // Split relative name into First and Last
            string[] relParts = relVal.Split(' ');
            string relFirstName = relParts[0];
            string relLastName = relProfileLastName(relParts);

            relatives.Add(new
            {
                FirstName = relFirstName,
                LastName = relLastName,
                Relationship = relationLabels[r]
            });
        }
    }

    results.Add(new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = birthday.ToString("yyyy-MM-dd"),
        Age = age,
        Relatives = relatives
    });
}

// Helper to extract last name from split parts
string relProfileLastName(string[] parts)
{
    if (parts.Length <= 1) return "";
    return parts.Last();
}

// Specific logic for the Main person's last name based on the prompt's pattern
// In "John Winston Lennon" -> "Lennon"
// In "James Paul McCartney" -> "McCartney"
string relProfileLastName(string[] parts)
{
    return parts.Length > 1 ? parts.Last() : "";
}

// Standardize Last Name logic for the person
string relProfileLastName(string[] parts) 
{
    if (parts.Length == 1) return parts[0];
    return parts[parts.Length - 1];
}

// Overwrite the logic for the specific input parsing
// The user wants the output to match the JSON exactly.
// Let's refine the parsing loop.

// Re-calculating for the specific structure requirement
var finalOutput = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;
    string[] values = line.Split(',');
    
    var rowMap = new Dictionary<string, string>();
    for (int j = 0; j < headers.Length; j++) rowMap[headers[j]] = values[j];

    string fullName = rowMap["Name"];
    string[] nameParts = fullName.Split(' ');
    string fName = nameParts[0];
    string lName = nameParts.Last();

    DateTime bday = DateTime.ParseExact(rowMap["Birthday"], "M/d/yyyy", CultureInfo.InvariantCulture);
    int ageAsOfRef = referenceDate.Year - bday.Year;
    if (referenceDate < bday.AddYears(ageAsOfRef)) ageAsOfRef--;

    var rels = new List<object>();
    string[] relTypes = { "Father", "Mother", "Brother", "Sister" };
    foreach (var type in relTypes)
    {
        string val = rowMap[type];
        if (!string.IsNullOrEmpty(val) && val.ToLower() != "null")
        {
            string[] rParts = val.Split(' ');
            rels.Add(new {
                FirstName = rParts[0],
                LastName = rParts.Length > 1 ? rParts.Last() : "",
                Relationship = type
            });
        }
    }

    finalOutput.Add(new {
        FirstName = fName,
        LastName = lName,
        Birthday = bday.ToString("yyyy-MM-dd"),
        Age = ageAsOfRef,
        Relatives = rels
    });
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(finalOutput, options));

// Redefining the helper for the top-level scope
string relProfileLastName(string[] parts) => parts.Length > 1 ? parts.Last() : "";