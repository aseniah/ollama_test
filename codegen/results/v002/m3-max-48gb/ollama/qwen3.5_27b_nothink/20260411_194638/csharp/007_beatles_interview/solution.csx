using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var dateToCalculate = new DateTime(2025, 7, 1);
var lines = File.ReadAllLines("input/input.csv");
var result = new List<object>();

// Skip header
for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    
    string fullName = parts[0];
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];
    
    // Parse Birthday
    DateTime birthday;
    try 
    {
        birthday = DateTime.ParseExact(parts[1], "M/d/yyyy", null);
    }
    catch (FormatException)
    {
        birthday = DateTime.ParseExact(parts[1], "d/M/yyyy", null);
    }

    // Calculate Age as of July 1, 2025
    int age = dateToCalculate.Year - birthday.Year;
    if (dateToCalculate < birthday.AddYears(age))
    {
        age--;
    }

    var relatives = new List<object>();

    // Father
    if (!string.IsNullOrEmpty(parts[3]) && parts[3].ToLower() != "null")
    {
        string[] fParts = parts[3].Trim().Split(' ');
        relatives.Add(new JsonObject 
        { 
            ["FirstName"] = fParts[0], 
            ["LastName"] = fParts.Length > 1 ? string.Join(" ", fParts, 1, fParts.Length - 1) : "", 
            ["Relationship"] = "Father" 
        });
    }

    // Mother
    if (!string.IsNullOrEmpty(parts[4]) && parts[4].ToLower() != "null")
    {
        string[] mParts = parts[4].Trim().Split(' ');
        relatives.Add(new JsonObject 
        { 
            ["FirstName"] = mParts[0], 
            ["LastName"] = mParts.Length > 1 ? string.Join(" ", mParts, 1, mParts.Length - 1) : "", 
            ["Relationship"] = "Mother" 
        });
    }

    // Brother
    if (!string.IsNullOrEmpty(parts[5]) && parts[5].ToLower() != "null")
    {
        string[] bParts = parts[5].Trim().Split(' ');
        relatives.Add(new JsonObject 
        { 
            ["FirstName"] = bParts[0], 
            ["LastName"] = bParts.Length > 1 ? string.Join(" ", bParts, 1, bParts.Length - 1) : "", 
            ["Relationship"] = "Brother" 
        });
    }

    // Sister
    if (!string.IsNullOrEmpty(parts[6]) && parts[6].ToLower() != "null")
    {
        string[] sParts = parts[6].Trim().Split(' ');
        relatives.Add(new JsonObject 
        { 
            ["FirstName"] = sParts[0], 
            ["LastName"] = sParts.Length > 1 ? string.Join(" ", sParts, 1, sParts.Length - 1) : "", 
            ["Relationship"] = "Sister" 
        });
    }

    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = new JsonArray(relatives)
    };

    result.Add(person);
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(result, options));