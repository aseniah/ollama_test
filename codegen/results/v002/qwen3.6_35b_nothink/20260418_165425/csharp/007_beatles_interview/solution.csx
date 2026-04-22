using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the input CSV
string[] lines = File.ReadAllLines("input/input.csv");

// Skip header
var peopleList = new List<JsonObject>();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;

    string[] parts = line.Split(',');
    
    // Parse Name into FirstName and LastName
    // The example shows:
    // John Winston Lennon -> John, Lennon
    // James Paul McCartney -> James, McCartney
    // Ringo Starr -> Ringo, Starr
    // George Harrison -> George, Harrison
    // It seems like first word is FirstName, last word is LastName. Middle names are ignored.
    
    string nameStr = parts[0].Trim();
    string[] nameParts = nameStr.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    // Parse Birthday
    string birthdayStr = parts[1].Trim();
    DateTime birthday = DateTime.Parse(birthdayStr);
    string birthdayFormatted = birthday.ToString("yyyy-MM-dd");

    // Calculate age as of July 1, 2025
    DateTime referenceDate = new DateTime(2025, 7, 1);
    int age = referenceDate.Year - birthday.Year;
    if (birthday.Date > referenceDate.AddYears(age).Date)
    {
        age--;
    }

    // Parse Died (not used in output format, but good to note)
    // string diedStr = parts[2].Trim();

    // Parse Relatives
    // Father, Mother, Brother, Sister
    string fatherStr = parts[3].Trim();
    string motherStr = parts[4].Trim();
    string brotherStr = parts[5].Trim();
    string sisterStr = parts[6].Trim();

    var relativesList = new JsonArray();

    // Helper to add relative if not null
    if (fatherStr != "null" && !string.IsNullOrEmpty(fatherStr))
    {
        var fParts = fatherStr.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        var relative = new JsonObject
        {
            ["FirstName"] = fParts[0],
            ["LastName"] = fParts[fParts.Length - 1],
            ["Relationship"] = "Father"
        };
        relativesList.Add(relative);
    }

    if (motherStr != "null" && !string.IsNullOrEmpty(motherStr))
    {
        var mParts = motherStr.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        var relative = new JsonObject
        {
            ["FirstName"] = mParts[0],
            ["LastName"] = mParts[mParts.Length - 1],
            ["Relationship"] = "Mother"
        };
        relativesList.Add(relative);
    }

    if (brotherStr != "null" && !string.IsNullOrEmpty(brotherStr))
    {
        var bParts = brotherStr.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        var relative = new JsonObject
        {
            ["FirstName"] = bParts[0],
            ["LastName"] = bParts[bParts.Length - 1],
            ["Relationship"] = "Brother"
        };
        relativesList.Add(relative);
    }

    if (sisterStr != "null" && !string.IsNullOrEmpty(sisterStr))
    {
        var sParts = sisterStr.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        var relative = new JsonObject
        {
            ["FirstName"] = sParts[0],
            ["LastName"] = sParts[sParts.Length - 1],
            ["Relationship"] = "Sister"
        };
        relativesList.Add(relative);
    }

    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayFormatted,
        ["Age"] = age,
        ["Relatives"] = relativesList
    };

    peopleList.Add(person);
}

// Create final array
var json = new JsonArray();
foreach (var person in peopleList)
{
    json.Add(person);
}

// Serialize to string with pretty printing to match expected format style
var options = new JsonSerializerOptions { WriteIndented = true };
string jsonOutput = json.ToJsonString(options);

Console.WriteLine(jsonOutput);