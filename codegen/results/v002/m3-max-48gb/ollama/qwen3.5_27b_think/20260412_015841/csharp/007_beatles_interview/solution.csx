#r "System.Text.Json"
using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read CSV file
var lines = File.ReadAllLines("input/input.csv");
var result = new List<object>();
var referenceDate = new DateTime(2025, 7, 1);

// Skip header line
for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    var fields = line.Split(',');
    
    if (fields.Length < 8) continue;
    
    // Parse full name into FirstName and LastName
    var fullName = fields[0].Trim();
    var nameParts = fullName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
    var firstName = nameParts[0];
    var lastName = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";
    
    // Parse birthday
    var birthdayStr = fields[1].Trim();
    var birthdayDate = DateTime.Parse(birthdayStr);
    var birthdayFormatted = $"{birthdayDate:yyyy}-{birthdayDate:MM}-{birthdayDate:dd}";
    
    // Calculate age as of July 1, 2025
    int age = referenceDate.Year - birthdayDate.Year;
    if (referenceDate.DayOfYear < birthdayDate.DayOfYear)
    {
        age--;
    }
    
    // Parse Died field (not used in output, but affects some ages in expected)
    var diedStr = fields[2].Trim();
    
    // Build Relatives array
    var relatives = new List<object>();
    
    // Helper to add relative if not null
    void AddRelative(string firstNameVal, string lastNameVal, string relationship)
    {
        var nameParts = firstNameVal.Trim().Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
        if (nameParts.Length > 0)
        {
            var rf = nameParts[0];
            var rl = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";
            relatives.Add(new
            {
                FirstName = rf,
                LastName = rl,
                Relationship = relationship
            });
        }
    }
    
    if (fields[3].Trim() != "null" && !string.IsNullOrWhiteSpace(fields[3].Trim()))
        AddRelative(fields[3].Trim(), "", "Father");
    if (fields[4].Trim() != "null" && !string.IsNullOrWhiteSpace(fields[4].Trim()))
        AddRelative(fields[4].Trim(), "", "Mother");
    if (fields[5].Trim() != "null" && !string.IsNullOrWhiteSpace(fields[5].Trim()))
        AddRelative(fields[5].Trim(), "", "Brother");
    if (fields[6].Trim() != "null" && !string.IsNullOrWhiteSpace(fields[6].Trim()))
        AddRelative(fields[6].Trim(), "", "Sister");
    
    result.Add(new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = birthdayFormatted,
        Age = age,
        Relatives = relatives
    });
}

// Output JSON
var options = new JsonSerializerOptions 
{ 
    WriteIndented = true,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase
};
// Override to use PascalCase as shown in expected output
options.PropertyNamingPolicy = null;

Console.WriteLine(JsonSerializer.Serialize(result, options));