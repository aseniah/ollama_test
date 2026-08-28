#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var referenceDate = new DateTime(2025, 7, 1);
var people = new List<object>();
var lines = File.ReadAllLines("input/input.csv");

// Skip header line if present and process data
for (int i = 1; i < lines.Length; i++)
{
    var row = lines[i];
    if (string.IsNullOrWhiteSpace(row)) continue;

    var parts = row.Split(',');
    if (parts.Length < 6) continue;

    // Parse Name (Split by last space for Last/First, or handle middle names)
    var fullName = parts[0].Trim();
    var nameParts = fullName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";

    // Parse Birthday (MM/D/YYYY or M/D/YYYY)
    var birthDateStr = parts[1].Trim();
    var birthDate = DateTime.Parse(birthDateStr);
    
    // Calculate Age as of July 1, 2025
    int age = referenceDate.Year - birthDate.Year;
    if (referenceDate.DayOfYear < birthDate.DayOfYear)
    {
        age--;
    }

    // Build Relatives List
    var relatives = new List<object>();

    // Helper to add relative if not null/empty
    void AddRelative(string name, string relationType)
    {
        if (!string.IsNullOrWhiteSpace(name))
        {
            var relParts = name.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
            var rFirstName = relParts[0];
            var rLastName = relParts.Length > 1 ? relParts[relParts.Length - 1] : "";

            relatives.Add(new Dictionary<string, object>
            {
                { "FirstName", rFirstName },
                { "LastName", rLastName },
                { "Relationship", relationType }
            });
        }
    }

    // Map CSV columns to relationships
    // Index: Name(0), Birthday(1), Died(2), Father(3), Mother(4), Brother(5), Sister(6)
    AddRelative(parts[3].Trim(), "Father");
    AddRelative(parts[4].Trim(), "Mother");
    if (parts.Length > 5) AddRelative(parts[5].Trim(), "Brother");
    if (parts.Length > 6) AddRelative(parts[6].Trim(), "Sister");

    // Construct the Person Object matching expected JSON structure
    var person = new Dictionary<string, object>
    {
        { "FirstName", firstName },
        { "LastName", lastName },
        { "Birthday", birthDate.ToString("yyyy-MM-dd") },
        { "Age", age },
        { "Relatives", relatives } // Will serialize to array of objects
    };

    people.Add(person);
}

// Configure JSON options to match expected output (no indentation not strictly required by prompt, 
// but standard formatting is usually safer. The prompt shows indented, so we use WriteIndented)
var options = new JsonSerializerOptions
{
    WriteIndented = true,
    PropertyNamingPolicy = null // Keep exact casing as provided in dictionaries
};

var jsonOutput = JsonSerializer.Serialize(people, options);
Console.WriteLine(jsonOutput);