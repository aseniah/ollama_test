using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var inputPath = "input/input.csv";
var lines = File.ReadAllLines(inputPath);

var referenceDate = new DateTime(2025, 7, 1);
var result = new List<JsonObject>();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    var parts = line.Split(',');
    
    // Parse Name: "First Last" or "First Middle Last" -> Take first word as FirstName, last word as LastName
    // Based on input: "John Winston Lennon" -> FirstName: "John", LastName: "Lennon"
    // "James Paul McCartney" -> FirstName: "James", LastName: "McCartney"
    var nameParts = parts[0].Trim().Split(' ', StringSplitOptions.RemoveEmptyEntries);
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];

    // Parse Birthday: MM/D/YYYY
    var birthdayStr = parts[1].Trim();
    var birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    var birthdayFormatted = birthday.ToString("yyyy-MM-dd");

    // Calculate Age as of July 1, 2025
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate < birthday.AddYears(age))
    {
        age--;
    }

    // Parse Relatives
    var relatives = new List<JsonObject>();
    
    // Helper to add relative if not null
    void AddRelative(string fullName, string relationship)
    {
        if (fullName != "null" && !string.IsNullOrWhiteSpace(fullName))
        {
            var relParts = fullName.Trim().Split(' ', StringSplitOptions.RemoveEmptyEntries);
            if (relParts.Length > 0)
            {
                var relFirstName = relParts[0];
                var relLastName = relParts[relParts.Length - 1];
                relatives.Add(new JsonObject
                {
                    ["FirstName"] = relFirstName,
                    ["LastName"] = relLastName,
                    ["Relationship"] = relationship
                });
            }
        }
    }

    // Father (index 3)
    AddRelative(parts[3], "Father");
    // Mother (index 4)
    AddRelative(parts[4], "Mother");
    // Brother (index 5)
    AddRelative(parts[5], "Brother");
    // Sister (index 6)
    AddRelative(parts[6], "Sister");

    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayFormatted,
        ["Age"] = age,
        ["Relatives"] = JsonNode.Parse(new System.Text.Json.JsonSerializer.Serialize(relatives, new JsonSerializerOptions { WriteIndented = false }))
    };
    
    // Ensure Relatives is an array in the node
    person["Relatives"] = new JsonArray(relatives);

    result.Add(person);
}

var options = new JsonSerializerOptions
{
    WriteIndented = true,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
    Encoder = System.Text.Json.Encoder.UnsafeRelaxedJsonEscaping
};

// Force PascalCase for keys to match expected output explicitly since we constructed manually
var output = new JsonArray(result);
var jsonOutput = output.ToJsonString(new JsonSerializerOptions
{
    WriteIndented = true,
    PropertyNamingPolicy = null // Default is PascalCase for our manually set keys
});

Console.Write(jsonOutput);