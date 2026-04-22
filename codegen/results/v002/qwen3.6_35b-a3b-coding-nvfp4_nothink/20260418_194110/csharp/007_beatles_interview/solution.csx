using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the input CSV
string[] lines = File.ReadAllLines("input/input.csv");

// Define the reference date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

// Skip header line
List<JsonNode> people = new List<JsonNode>();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;

    string[] parts = line.Split(',');
    if (parts.Length < 7) continue;

    string fullName = parts[0].Trim();
    string birthdayStr = parts[1].Trim();
    string fatherName = parts[3].Trim();
    string motherName = parts[4].Trim();
    string brotherName = parts[5].Trim();
    string sisterName = parts[6].Trim();

    // Parse name into FirstName and LastName
    // The format seems to be "FirstName MiddleName LastName" or "FirstName LastName"
    // Based on expected output:
    // John Winston Lennon -> FirstName: John, LastName: Lennon
    // James Paul McCartney -> FirstName: James, LastName: McCartney
    // Ringo Starr -> FirstName: Ringo, LastName: Starr
    // George Harrison -> FirstName: George, LastName: Harrison
    
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    // Parse birthday
    DateTime birthday;
    // The date format is M/D/YYYY or MM/D/YYYY etc.
    if (!DateTime.TryParse(birthdayStr, out birthday))
    {
        // Fallback or error handling if needed, but let's assume valid input
        continue;
    }

    // Calculate age as of July 1, 2025
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate.Month < birthday.Month || (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
    {
        age--;
    }

    // Create the person object
    JsonObject person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age
    };

    // Create relatives list
    JsonArray relatives = new JsonArray();

    // Add Father if not null
    if (fatherName != "null" && !string.IsNullOrEmpty(fatherName))
    {
        string[] fParts = fatherName.Split(' ');
        string fFirst = fParts[0];
        string fLast = fParts[fParts.Length - 1];
        
        relatives.Add(new JsonObject
        {
            ["FirstName"] = fFirst,
            ["LastName"] = fLast,
            ["Relationship"] = "Father"
        });
    }

    // Add Mother if not null
    if (motherName != "null" && !string.IsNullOrEmpty(motherName))
    {
        string[] mParts = motherName.Split(' ');
        string mFirst = mParts[0];
        string mLast = mParts[mParts.Length - 1];
        
        relatives.Add(new JsonObject
        {
            ["FirstName"] = mFirst,
            ["LastName"] = mLast,
            ["Relationship"] = "Mother"
        });
    }

    // Add Brother if not null
    if (brotherName != "null" && !string.IsNullOrEmpty(brotherName))
    {
        string[] bParts = brotherName.Split(' ');
        string bFirst = bParts[0];
        string bLast = bParts[bParts.Length - 1];
        
        relatives.Add(new JsonObject
        {
            ["FirstName"] = bFirst,
            ["LastName"] = bLast,
            ["Relationship"] = "Brother"
        });
    }

    // Add Sister if not null
    if (sisterName != "null" && !string.IsNullOrEmpty(sisterName))
    {
        string[] sParts = sisterName.Split(' ');
        string sFirst = sParts[0];
        string sLast = sParts[sParts.Length - 1];
        
        relatives.Add(new JsonObject
        {
            ["FirstName"] = sFirst,
            ["LastName"] = sLast,
            ["Relationship"] = "Sister"
        });
    }

    person["Relatives"] = relatives;
    people.Add(person);
}

// Output the JSON array
JsonArray result = new JsonArray();
foreach (JsonObject person in people)
{
    result.Add(person);
}

Console.WriteLine(result.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));