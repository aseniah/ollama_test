using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var records = new List<JsonNode>();

for (int i = 1; i < lines.Length; i++)
{
    var fields = lines[i].Split(',');
    
    var record = new JsonObject();
    
    var fullName = fields[0].Split(' ');
    record["FirstName"] = fullName[0];
    record["LastName"] = string.Join(" ", fullName.Skip(1));
    
    var birthdayStr = fields[1];
    var birthday = DateTime.Parse(birthdayStr);
    record["Birthday"] = birthday.ToString("yyyy-MM-dd");
    
    var age = CalculateAge(birthday, new DateTime(2025, 7, 1));
    record["Age"] = age;
    
    var relatives = new JsonArray();
    
    var father = fields[3];
    if (!string.IsNullOrEmpty(father))
    {
        var fatherParts = father.Split(' ');
        var fatherObj = new JsonObject
        {
            ["FirstName"] = fatherParts[0],
            ["LastName"] = string.Join(" ", fatherParts.Skip(1)),
            ["Relationship"] = "Father"
        };
        relatives.Add(fatherObj);
    }
    
    var mother = fields[4];
    if (!string.IsNullOrEmpty(mother))
    {
        var motherParts = mother.Split(' ');
        var motherObj = new JsonObject
        {
            ["FirstName"] = motherParts[0],
            ["LastName"] = string.Join(" ", motherParts.Skip(1)),
            ["Relationship"] = "Mother"
        };
        relatives.Add(motherObj);
    }
    
    var brother = fields[5];
    if (!string.IsNullOrEmpty(brother) && brother != "null")
    {
        var brotherParts = brother.Split(' ');
        var brotherObj = new JsonObject
        {
            ["FirstName"] = brotherParts[0],
            ["LastName"] = string.Join(" ", brotherParts.Skip(1)),
            ["Relationship"] = "Brother"
        };
        relatives.Add(brotherObj);
    }
    
    var sister = fields[6];
    if (!string.IsNullOrEmpty(sister) && sister != "null")
    {
        var sisterParts = sister.Split(' ');
        var sisterObj = new JsonObject
        {
            ["FirstName"] = sisterParts[0],
            ["LastName"] = string.Join(" ", sisterParts.Skip(1)),
            ["Relationship"] = "Sister"
        };
        relatives.Add(sisterObj);
    }
    
    record["Relatives"] = relatives;
    
    records.Add(record);
}

Console.WriteLine(records.ToJsonString());

int CalculateAge(DateTime birthDate, DateTime referenceDate)
{
    int age = referenceDate.Year - birthDate.Year;
    if (birthDate > referenceDate.AddYears(-age)) age--;
    return age;
}