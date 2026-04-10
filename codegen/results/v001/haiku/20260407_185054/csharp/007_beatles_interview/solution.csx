using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var headers = lines[0].Split('\t');

var result = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var values = lines[i].Split('\t');

    var nameData = values[0].Split(' ');
    var firstName = nameData[0];
    var lastName = string.Join(" ", nameData.Skip(1));

    var birthdayStr = values[1];
    var birthDate = DateTime.Parse(birthdayStr);

    var referenceDate = new DateTime(2025, 7, 1);
    var age = referenceDate.Year - birthDate.Year;
    if (birthDate.AddYears(age) > referenceDate)
    {
        age--;
    }

    var obj = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthDate.ToString("yyyy-MM-dd"),
        ["Age"] = age
    };

    var relatives = new JsonArray();

    var father = values[3];
    if (father != "null")
    {
        var fatherData = father.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = fatherData[0],
            ["LastName"] = string.Join(" ", fatherData.Skip(1)),
            ["Relationship"] = "Father"
        });
    }

    var mother = values[4];
    if (mother != "null")
    {
        var motherData = mother.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = motherData[0],
            ["LastName"] = string.Join(" ", motherData.Skip(1)),
            ["Relationship"] = "Mother"
        });
    }

    var brother = values[5];
    if (brother != "null")
    {
        var brotherData = brother.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = brotherData[0],
            ["LastName"] = string.Join(" ", brotherData.Skip(1)),
            ["Relationship"] = "Brother"
        });
    }

    var sister = values[6];
    if (sister != "null")
    {
        var sisterData = sister.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = sisterData[0],
            ["LastName"] = string.Join(" ", sisterData.Skip(1)),
            ["Relationship"] = "Sister"
        });
    }

    obj["Relatives"] = relatives;
    result.Add(obj);
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(result, options));
