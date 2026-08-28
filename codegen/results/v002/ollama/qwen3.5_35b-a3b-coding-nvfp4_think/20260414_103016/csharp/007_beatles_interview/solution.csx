using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

var csvContent = File.ReadAllLines("input/input.csv");
var header = csvContent[0];
var dataLines = csvContent.Skip(1);

var referenceDate = new DateTime(2025, 7, 1);

var people = new List<JsonObject>();

foreach (var line in dataLines)
{
    var parts = line.Split(',');
    
    var name = parts[0].Trim();
    var birthdayStr = parts[1].Trim();
    var father = parts[2].Trim();
    var mother = parts[3].Trim();
    var brother = parts[4].Trim();
    var sister = parts[5].Trim();
    
    var nameParts = name.Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];
    
    var birthdayParts = birthdayStr.Split('/');
    var birthday = new DateTime(
        int.Parse(birthdayParts[2]),
        int.Parse(birthdayParts[0]),
        int.Parse(birthdayParts[1])
    );
    
    var age = referenceDate.Year - birthday.Year;
    if (referenceDate < birthday.AddYears(age))
    {
        age--;
    }
    
    var relatives = new List<JsonObject>();
    
    if (father != "null" && !string.IsNullOrEmpty(father))
    {
        var fParts = father.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = fParts[0],
            ["LastName"] = fParts[fParts.Length - 1],
            ["Relationship"] = "Father"
        });
    }
    
    if (mother != "null" && !string.IsNullOrEmpty(mother))
    {
        var mParts = mother.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = mParts[0],
            ["LastName"] = mParts[mParts.Length - 1],
            ["Relationship"] = "Mother"
        });
    }
    
    if (brother != "null" && !string.IsNullOrEmpty(brother))
    {
        var bParts = brother.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = bParts[0],
            ["LastName"] = bParts[bParts.Length - 1],
            ["Relationship"] = "Brother"
        });
    }
    
    if (sister != "null" && !string.IsNullOrEmpty(sister))
    {
        var sParts = sister.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = sParts[0],
            ["LastName"] = sParts[sParts.Length - 1],
            ["Relationship"] = "Sister"
        });
    }
    
    people.Add(new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = new JsonArray(relatives)
    });
}

Console.WriteLine(JsonSerializer.Serialize(people, new JsonSerializerOptions { WriteIndented = true }));