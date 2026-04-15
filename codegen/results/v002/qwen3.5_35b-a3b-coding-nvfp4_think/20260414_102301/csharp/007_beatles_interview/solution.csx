using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;

var inputLines = File.ReadAllLines("input/input.csv");
var referenceDate = new DateOnly(2025, 7, 1);

var people = new List<JsonNode>();

foreach (var line in inputLines.Skip(1))
{
    var parts = line.Split(',');
    var name = parts[0];
    var birthdayStr = parts[1];
    var died = parts[2];
    var father = parts[3];
    var mother = parts[4];
    var brother = parts[5];
    var sister = parts[6];
    
    var nameParts = name.Split(' ');
    var firstName = nameParts[0];
    var lastName = string.Join(" ", nameParts.Skip(1));
    
    var birthdayParts = birthdayStr.Split('/');
    var birthMonth = int.Parse(birthdayParts[0]);
    var birthDay = int.Parse(birthdayParts[1]);
    var birthYear = int.Parse(birthdayParts[2]);
    var birthDate = new DateOnly(birthYear, birthMonth, birthDay);
    
    var age = referenceDate.Year - birthDate.Year;
    if (referenceDate.DayOfYear < birthDate.DayOfYear)
        age--;
    
    var relativeList = new List<JsonNode>();
    
    if (!string.IsNullOrEmpty(father) && father != "null")
        relativeList.Add(new JsonObject
        {
            ["FirstName"] = firstName,
            ["LastName"] = father,
            ["Relationship"] = "Father"
        });
    
    if (!string.IsNullOrEmpty(mother) && mother != "null")
        relativeList.Add(new JsonObject
        {
            ["FirstName"] = firstName,
            ["LastName"] = mother,
            ["Relationship"] = "Mother"
        });
    
    if (!string.IsNullOrEmpty(brother) && brother != "null")
        relativeList.Add(new JsonObject
        {
            ["FirstName"] = firstName,
            ["LastName"] = brother,
            ["Relationship"] = "Brother"
        });
    
    if (!string.IsNullOrEmpty(sister) && sister != "null")
        relativeList.Add(new JsonObject
        {
            ["FirstName"] = firstName,
            ["LastName"] = sister,
            ["Relationship"] = "Sister"
        });
    
    people.Add(new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthDate.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = new JsonArray(relativeList)
    });
}

Console.WriteLine(JsonSerializer.Serialize(people, new JsonSerializerOptions
{
    WriteIndented = true
}));