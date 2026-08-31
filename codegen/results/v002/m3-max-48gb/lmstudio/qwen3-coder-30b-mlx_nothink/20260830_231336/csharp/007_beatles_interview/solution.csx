using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var inputLines = File.ReadAllLines("input/input.csv");
var result = new List<JsonObject>();

var referenceDate = new DateTime(2025, 7, 1);

for (int i = 1; i < inputLines.Length; i++)
{
    var line = inputLines[i];
    var parts = line.Split(',');
    
    var person = new JsonObject();
    
    var fullName = parts[0].Trim();
    var birthdayStr = parts[1].Trim();
    var diedStr = parts[2].Trim();
    var father = parts[3].Trim();
    var mother = parts[4].Trim();
    var brother = parts[5].Trim();
    var sister = parts[6].Trim();
    
    var nameParts = fullName.Split(' ');
    person["FirstName"] = nameParts[0];
    person["LastName"] = nameParts.Length > 1 ? string.Join(" ", nameParts.Skip(1)) : "";
    
    var birthday = DateTime.Parse(birthdayStr);
    person["Birthday"] = birthday.ToString("yyyy-MM-dd");
    
    var age = referenceDate.Year - birthday.Year;
    if (referenceDate.Month < birthday.Month || 
        (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
    {
        age--;
    }
    person["Age"] = age;
    
    var relatives = new List<JsonObject>();
    
    if (!string.IsNullOrEmpty(father) && father != "null")
    {
        var fatherParts = father.Split(' ');
        var fatherObj = new JsonObject();
        fatherObj["FirstName"] = fatherParts[0];
        fatherObj["LastName"] = fatherParts.Length > 1 ? string.Join(" ", fatherParts.Skip(1)) : "";
        fatherObj["Relationship"] = "Father";
        relatives.Add(fatherObj);
    }
    
    if (!string.IsNullOrEmpty(mother) && mother != "null")
    {
        var motherParts = mother.Split(' ');
        var motherObj = new JsonObject();
        motherObj["FirstName"] = motherParts[0];
        motherObj["LastName"] = motherParts.Length > 1 ? string.Join(" ", motherParts.Skip(1)) : "";
        motherObj["Relationship"] = "Mother";
        relatives.Add(motherObj);
    }
    
    if (!string.IsNullOrEmpty(brother) && brother != "null")
    {
        var brotherParts = brother.Split(' ');
        var brotherObj = new JsonObject();
        brotherObj["FirstName"] = brotherParts[0];
        brotherObj["LastName"] = brotherParts.Length > 1 ? string.Join(" ", brotherParts.Skip(1)) : "";
        brotherObj["Relationship"] = "Brother";
        relatives.Add(brotherObj);
    }
    
    if (!string.IsNullOrEmpty(sister) && sister != "null")
    {
        var sisterParts = sister.Split(' ');
        var sisterObj = new JsonObject();
        sisterObj["FirstName"] = sisterParts[0];
        sisterObj["LastName"] = sisterParts.Length > 1 ? string.Join(" ", sisterParts.Skip(1)) : "";
        sisterObj["Relationship"] = "Sister";
        relatives.Add(sisterObj);
    }
    
    person["Relatives"] = new JsonArray(relatives.ToArray());
    
    result.Add(person);
}

Console.WriteLine(result.ToJsonString());