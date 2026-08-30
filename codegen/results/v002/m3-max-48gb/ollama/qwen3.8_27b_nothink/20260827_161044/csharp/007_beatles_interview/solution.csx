using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");

var referenceDate = new DateTime(2025, 7, 1);

var result = new JsonArray();

for (int i = 1; i < lines.Length; i++) // skip header
{
    var fields = lines[i].Split(',');
    
    // Parse name - split into first and last
    var fullName = fields[0].Trim();
    var nameParts = fullName.Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";
    
    // For names with more than 2 parts (e.g., "John Winston Lennon"), we need to figure out which is first and which is last
    // Looking at expected output: "John" "Lennon" - so middle names are dropped
    // "James Paul McCartney" -> "James" "McCartney"
    // "Ringo Starr" -> "Ringo" "Starr"
    // "George Harrison" -> "George" "Harrison"
    
    // So first name is the first part, last name is the last part
    // But what about middle names? They seem to be dropped.
    // Let's use first word as FirstName and last word as LastName
    
    // Actually, looking more carefully:
    // "John Winston Lennon" -> FirstName: "John", LastName: "Lennon"
    // "James Paul McCartney" -> FirstName: "James", LastName: "McCartney"
    // So it's first word and last word.
    
    // Parse birthday: M/D/YYYY format
    var birthdayStr = fields[1].Trim();
    var birthdayParts = birthdayStr.Split('/');
    var birthday = new DateTime(int.Parse(birthdayParts[2]), int.Parse(birthdayParts[0]), int.Parse(birthdayParts[1]));
    
    // Calculate age as of July 1, 2025
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate.Month < birthday.Month || (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
    {
        age--;
    }
    
    var entry = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = new JsonArray()
    };
    
    var relatives = entry["Relatives"] as JsonArray;
    
    // Parse father
    var fatherStr = fields[3].Trim();
    if (fatherStr != "null")
    {
        var fatherNameParts = fatherStr.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = fatherNameParts[0],
            ["LastName"] = fatherNameParts[fatherNameParts.Length - 1],
            ["Relationship"] = "Father"
        });
    }
    
    // Parse mother
    var motherStr = fields[4].Trim();
    if (motherStr != "null")
    {
        var motherNameParts = motherStr.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = motherNameParts[0],
            ["LastName"] = motherNameParts[motherNameParts.Length - 1],
            ["Relationship"] = "Mother"
        });
    }
    
    // Parse brother
    var brotherStr = fields[5].Trim();
    if (brotherStr != "null")
    {
        var brotherNameParts = brotherStr.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = brotherNameParts[0],
            ["LastName"] = brotherNameParts[brotherNameParts.Length - 1],
            ["Relationship"] = "Brother"
        });
    }
    
    // Parse sister
    var sisterStr = fields[6].Trim();
    if (sisterStr != "null")
    {
        var sisterNameParts = sisterStr.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = sisterNameParts[0],
            ["LastName"] = sisterNameParts[sisterNameParts.Length - 1],
            ["Relationship"] = "Sister"
        });
    }
    
    result.Add(entry);
}

Console.WriteLine(JsonSerializer.Serialize(result, new JsonSerializerOptions { WriteIndented = true }));