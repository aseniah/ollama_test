using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var referenceDate = new DateTime(2025, 7, 1);

var result = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var fields = lines[i].Split(',');
    
    var nameParts = fields[0].Trim().Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];
    
    var birthday = DateTime.ParseExact(fields[1].Trim(), "M/d/yyyy", CultureInfo.InvariantCulture);
    var birthdayStr = birthday.ToString("yyyy-MM-dd");
    
    string diedStr = fields[2].Trim();
    DateTime? died = null;
    if (diedStr != "null" && diedStr != "")
    {
        died = DateTime.ParseExact(diedStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    }
    
    int age;
    if (died.HasValue && died.Value < referenceDate)
        age = CalcAge(birthday, died.Value);
    else
        age = CalcAge(birthday, referenceDate);
    
    var relatives = new JsonArray();
    var relationshipFields = new[] { "Father", "Mother", "Brother", "Sister" };
    for (int j = 3; j < fields.Length; j++)
    {
        var relName = fields[j].Trim();
        if (relName != "null" && relName != "")
        {
            var relParts = relName.Split(' ');
            relatives.Add(new JsonObject
            {
                ["FirstName"] = relParts[0],
                ["LastName"] = relParts[relParts.Length - 1],
                ["Relationship"] = relationshipFields[j - 3]
            });
        }
    }
    
    result.Add(new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayStr,
        ["Age"] = age,
        ["Relatives"] = relatives
    });
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.Write(result.ToJsonString(options));

int CalcAge(DateTime birth, DateTime reference)
{
    int age = reference.Year - birth.Year;
    if (reference.Month < birth.Month || (reference.Month == birth.Month && reference.Day < birth.Day))
        age--;
    return age;
}