using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Path to the input file
string csvPath = "input/input.csv";

// Check if the input file exists before proceeding
if (!File.Exists(csvPath))
{
    return;
}

// Read all lines from the CSV file
string[] lines = File.ReadAllLines(csvPath);
if (lines.Length <= 1)
{
    // If only the header or no lines exist, output an empty array
    Console.WriteLine("[]");
    return;
}

// The reference date for calculating age for those who haven't died
DateTime referenceDate = new DateTime(2025, 7, 1);
JsonArray results = new JsonArray();

// Iterate through the CSV lines, skipping the header
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    // Split the CSV line by comma
    string[] columns = line.Split(',');
    if (columns.Length < 7) continue;

    string fullName = columns[0].Trim();
    string birthStr = columns[1].Trim();
    string diedStr = columns[2].Trim();
    string fatherStr = columns[3].Trim();
    string motherStr = `columns[4]`.Trim(); // Wait, fixing index typo below
    string mother = columns[4].Trim();
    string brother = columns[5].Trim();
    string sister = columns[6].Trim();

    // Parse the birth date (Format: M/d/yyyy)
    DateTime birthday = DateTime.ParseExact(birthStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    
    // Determine the date to use for age calculation (either death date or July 1, 2025)
    DateTime? diedDate = diedStr == "null" ? null : (DateTime?)DateTime.ParseExact(diedStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    DateTime calculationRef = diedDate ?? referenceDate;

    // Calculate age
    int age = calculationRef.Year - birthday.Year;
    if (calculationRef.Month < birthday.Month || (calculationRef.Month == birthday.Month && calculationRef.Day < birthday.Day))
    {
        age--;
    }

    // Construct the person object
    var person = new JsonObject
    {
        ["FirstName"] = GetFirstName(fullName),
        ["LastName"] = GetLastName(fullName),
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = new JsonArray()
    };

    var relativesArray = (JsonArray)person["Relatives"];

    // Helper to add relatives to the array
    void AddRelative(string relativeFullName, string relationship)
    {
        if (relativeFullName != "null" && !string.IsNullOrWhiteSpace(relativeFullName))
        {
            relativesArray.Add(new JsonObject
            {
                ["FirstName"] = GetFirstName(relativeFullName),
                ["LastName"] = GetLastName(relativeFullName),
                ["Relationship"] = relationship
            });
        }
    }

    AddRelative(columns[3].Trim(), "Father");
    AddRelative(columns[4].Trim(), "Mother");
    AddRelative(columns[5].Trim(), "Brother");
    AddRelative(columns[6].Trim(), "Sister");

    results.Add(person);
}

// Output the final JSON array to stdout
Console.WriteLine(results.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));

// Helper functions to parse names
string GetFirstName(string fullName)
{
    string[] parts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    return parts.Length > 0 ? parts[0] : "";
}

string GetLastName(string fullName)
{
    string[] parts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    return parts.Length > 1 ? parts[^1] : "";
}