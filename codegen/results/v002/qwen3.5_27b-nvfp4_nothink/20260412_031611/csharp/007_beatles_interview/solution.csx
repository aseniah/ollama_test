#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Reference date for age calculation: July 1, 2025
DateTime referenceDate = new DateTime(2025, 7, 1);

string filePath = "input/input.csv";
string[] lines = File.ReadAllLines(filePath);
if (lines.Length == 0)
{
    Console.WriteLine("[]");
    return;
}

// Skip header
var records = new List<JsonNode>();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    // Split by comma. Note: Names might contain spaces but not commas in this dataset.
    string[] parts = line.Split(',');

    if (parts.Length < 6) continue;

    // Parse Name
    string fullName = parts[0].Trim();
    int lastSpaceIndex = fullName.LastIndexOf(' ');
    string firstName = (lastSpaceIndex > -1) ? fullName.Substring(0, lastSpaceIndex).Trim() : fullName;
    string lastName = (lastSpaceIndex > -1) ? fullName.Substring(lastSpaceIndex + 1).Trim() : "";

    // Parse Birthday: M/D/YYYY or MM/DD/YYYY
    string birthdayStr = parts[1].Trim();
    DateTime birthDate;
    
    if (!DateTime.TryParseExact(birthdayStr, "M/d/yyyy", null, System.Globalization.DateTimeStyles.None, out birthDate))
    {
        // Fallback to invariant culture parsing just in case
        if (!DateTime.TryParse(birthdayStr, System.Globalization.CultureInfo.InvariantCulture, System.Globalization.DateTimeStyles.AssumeLocal, out birthDate))
        {
            continue; 
        }
    }

    // Calculate Age as of July 1, 2025
    int age = referenceDate.Year - birthDate.Year;
    if (referenceDate.DayOfYear < birthDate.DayOfYear)
    {
        age--;
    }

    // Ensure Birthday is formatted as "yyyy-MM-dd"
    string formattedBirthday = birthDate.ToString("yyyy-MM-dd");

    // Parse Relatives
    var relatives = new List<JsonNode>();
    
    // Helper to add relative if not null/empty
    void AddRelative(string firstNameRel, string lastNameRel, string relationship)
    {
        if (string.IsNullOrWhiteSpace(firstNameRel) || string.IsNullOrWhiteSpace(lastNameRel)) return;
        
        string fName = firstNameRel.Trim();
        string lName = lastNameRel.Trim();
        
        // Check for "null" string explicitly passed in CSV
        if (fName.Equals("null", StringComparison.OrdinalIgnoreCase) || 
            lName.Equals("null", StringComparison.OrdinalIgnoreCase)) return;

        var relNode = new JsonObject
        {
            ["FirstName"] = fName,
            ["LastName"] = lName,
            ["Relationship"] = relationship
        };
        relatives.Add(relNode);
    }

    // Father (Index 3)
    string fatherPart = parts[3].Trim();
    AddRelative(fatherPart.Split(' ')[0], fatherPart.Substring(fatherPart.LastIndexOf(' ')).Trim(), "Father");
    
    // Mother (Index 4)
    string motherPart = parts[4].Trim();
    AddRelative(motherPart.Split(' ')[0], motherPart.Substring(motherPart.LastIndexOf(' ')).Trim(), "Mother");

    // Brother (Index 5)
    string brotherPart = parts.Length > 5 ? parts[5].Trim() : "";
    AddRelative(brotherPart.Split(' ')[0], brotherPart.Substring(brotherPart.LastIndexOf(' ')).Trim(), "Brother");

    // Sister (Index 6)
    string sisterPart = parts.Length > 6 ? parts[6].Trim() : "";
    AddRelative(sisterPart.Split(' ')[0], sisterPart.Substring(sisterPart.LastIndexOf(' ')).Trim(), "Sister");

    var personNode = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = formattedBirthday,
        ["Age"] = age,
        ["Relatives"] = JsonArray.FromObject(relatives)
    };

    records.Add(personNode);
}

// Output JSON
var outputOptions = new JsonSerializerOptions
{
    WriteIndented = true
};

string jsonOutput = JsonSerializer.Serialize(records, outputOptions);
Console.WriteLine(jsonOutput);