using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var inputLines = File.ReadAllLines("input/input.csv");
var header = inputLines[0].Split(',');
var dataLines = inputLines.Skip(1);

var people = new List<JsonObject>();

var referenceDate = new DateTime(2025, 7, 1);

foreach (var line in dataLines)
{
    var values = line.Split(',');
    var person = new JsonObject();
    
    var fullName = values[0];
    var birthdayStr = values[1];
    
    // Parse name
    var nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    person["FirstName"] = nameParts[0];
    person["LastName"] = nameParts.Length > 1 ? string.Join(" ", nameParts.Skip(1)) : "";
    
    // Parse birthday
    var birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    person["Birthday"] = birthday.ToString("yyyy-MM-dd");
    
    // Calculate age as of July 1, 2025
    var age = referenceDate.Year - birthday.Year;
    if (referenceDate.Month < birthday.Month || 
        (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
        age--;
    person["Age"] = age;
    
    // Parse relatives
    var relatives = new JsonArray();
    
    // Father
    if (!string.IsNullOrEmpty(values[3]))
    {
        var fatherName = values[3];
        var fatherParts = fatherName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        var relative = new JsonObject
        {
            ["FirstName"] = fatherParts[0],
            ["LastName"] = fatherParts.Length > 1 ? string.Join(" ", fatherParts.Skip(1)) : "",
            ["Relationship"] = "Father"
        };
        relatives.Add(relative);
    }
    
    // Mother
    if (!string.IsNullOrEmpty(values[4]))
    {
        var motherName = values[4];
        var motherParts = motherName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        var relative = new JsonObject
        {
            ["FirstName"] = motherParts[0],
            ["LastName"] = motherParts.Length > 1 ? string.Join(" ", motherParts.Skip(1)) : "",
            ["Relationship"] = "Mother"
        };
        relatives.Add(relative);
    }
    
    // Brother
    if (!string.IsNullOrEmpty(values[5]) && !values[5].Equals("null", StringComparison.OrdinalIgnoreCase))
    {
        var brotherName = values[5];
        var brotherParts = brotherName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        var relative = new JsonObject
        {
            ["FirstName"] = brotherParts[0],
            ["LastName"] = brotherParts.Length > 1 ? string.Join(" ", brotherParts.Skip(1)) : "",
            ["Relationship"] = "Brother"
        };
        relatives.Add(relative);
    }
    
    // Sister
    if (!string.IsNullOrEmpty(values[6]) && !values[6].Equals("null", StringComparison.OrdinalIgnoreCase))
    {
        var sisterName = values[6];
        var sisterParts = sisterName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        var relative = new JsonObject
        {
            ["FirstName"] = sisterParts[0],
            ["LastName"] = sisterParts.Length > 1 ? string.Join(" ", sisterParts.Skip(1)) : "",
            ["Relationship"] = "Sister"
        };
        relatives.Add(relative);
    }
    
    person["Relatives"] = relatives;
    people.Add(person);
}

var json = new JsonArray();
foreach (var person in people)
{
    json.Add(person);
}

Console.WriteLine(json.ToJsonString());