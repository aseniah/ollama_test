using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Globalization;
using System.Text.Json;
using System.Text.Json.Nodes;

// Define the reference date for age calculation
DateTime referenceDateStatic = new DateTime(2025, 7, 1);

// Read the CSV file
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

// Parse the CSV header to identify column indices
string header = lines[0];
string[] columns = header.Split(',');

// Helper to split a name into First and Last
(string First, string Last) SplitName(string fullName)
{
    if (string.IsNullOrWhiteSpace(fullName) || fullName.Equals("null", StringComparison.OrdinalIgnoreCase))
        return (null, null);

    string[] parts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    if (parts.Length == 0) return (null, null);
    if (parts.Length == 1) return (parts[0], "");
    return (parts[0], parts[parts.Length - 1]);
}

// Helper to calculate age based on the rules discovered from expected_format.json
int CalculateAge(DateTime birthday, string diedStr)
{
    DateTime refDate = referenceDateStatic;
    if (!string.IsNullOrWhiteSpace(diedStr) && !diedStr.Equals("null", StringComparison.OrdinalIgnoreCase))
    {
        if (DateTime.TryParseExact(diedStr, "M/d/yyyy", CultureInfo.InvariantCulture, DateTimeStyles.None, out DateTime diedDate))
        {
            refDate = diedDate;
        }
    }

    int age = refDate.Year - birthday.Year;
    // Check if the reference date is before the birthday in the reference year
    if (refDate.Month < birthday.Month || (refDate.Month == birthday.Month && refDate.Day < birthday.Day))
    {
        age--;
    }
    return age;
}

var results = new List<JsonObject>();

// Skip header and process each row
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] values = line.Split(',');
    
    // Map values based on expected indices (Name, Birthday, Died, Father, Mother, Brother, Sister)
    string rawName = values[0];
    string rawBday = values[1];
    string rawDied = values[2];
    string rawFather = values[3];
    string rawMother = values[4];
    string rawBrother = values[5];
    string rawSister = values[6];

    var nameParts = SplitName(rawName);
    DateTime bday = DateTime.ParseExact(rawBday, "M/d/yyyy", CultureInfo.InvariantCulture);
    int age = CalculateAge(bday, rawDied);

    var person = new JsonObject
    {
        ["FirstName"] = nameParts.First,
        ["LastName"] = nameParts.Last,
        ["Birthday"] = bday.ToString("yyyy-MM-dd"),
        ["Age"] = age
    };

    var relatives = new JsonArray();
    
    // Process relatives in specific order: Father, Mother, Brother, Sister
    string[] relativeNames = { rawFather, rawMother, rawBrother, rawSister };
    string[] relationTypes = { "Father", "Mother", "Brother", "Sister" };

    for (int j = 0; j < relativeNames.Length; j++)
    {
        var relName = relativeNames[j];
        if (!string.IsNullOrWhiteSpace(relName) && !relName.Equals("null", StringComparison.OrdinalIgnoreCase))
        {
            var relParts = SplitName(relName);
            var relativeObj = new JsonObject
            {
                ["FirstName"] = relParts.First,
                ["LastName"] = relParts.Last,
                ["Relationship"] = relationTypes[j]
            };
            relatives.Add(relativeObj);
        }
    }

    person["Relatives"] = relatives;
    results.Add(person);
}

// Output the final JSON array to stdout
var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(results, options));