using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var people = new List<JsonObject>();

for (int i = 1; i < lines.Length; i++) // Skip header
{
    var parts = lines[i].Split(',');
    if (parts.Length < 3) continue;
    
    var firstName = parts[0].Trim();
    var lastName = parts[1].Trim();
    var birthDateStr = parts[2].Trim();
    
    if (string.IsNullOrEmpty(firstName) || string.IsNullOrEmpty(lastName) || string.IsNullOrEmpty(birthDateStr))
        continue;
    
    if (DateTime.TryParseExact(birthDateStr, "MM/dd/yyyy", CultureInfo.InvariantCulture, DateTimeStyles.None, out DateTime birthDate))
    {
        var age = CalculateAge(birthDate, new DateTime(2025, 7, 1));
        
        var person = new JsonObject
        {
            ["firstName"] = firstName,
            ["lastName"] = lastName,
            ["age"] = age
        };
        
        people.Add(person);
    }
}

Console.WriteLine(JsonSerializer.Serialize(people));

int CalculateAge(DateTime birthDate, DateTime referenceDate)
{
    var age = referenceDate.Year - birthDate.Year;
    if (birthDate.AddYears(age) > referenceDate)
        age--;
    return age;
}