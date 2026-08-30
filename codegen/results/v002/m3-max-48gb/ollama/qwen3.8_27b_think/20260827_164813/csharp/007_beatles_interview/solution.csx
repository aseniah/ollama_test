using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string csvPath = "input/input.csv";
string[] lines = File.ReadAllLines(csvPath);
string header = lines[0];
string[] headers = header.Split(',');

var people = new List<JsonObject>();
DateTime referenceDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++)
{
    string[] fields = lines[i].Split(',');
    string name = fields[0];
    string birthdayStr = fields[1];
    string diedStr = fields[2];
    string fatherStr = fields[3];
    string motherStr = fields[4];
    string brotherStr = fields[5];
    string sisterStr = fields[6];

    // Parse name
    string[] nameParts = name.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    // Parse birthday (MM/DD/YYYY)
    string[] bd = birthdayStr.Split('/');
    int bMonth = int.Parse(bd[0]);
    int bDay = int.Parse(bd[1]);
    int bYear = int.Parse(bd[2]);
    DateTime birthday = new DateTime(bYear, bMonth, bDay);
    string birthdayIso = $"{bYear:D4}-{bMonth:D2}-{bDay:D2}";

    // Calculate age
    int age;
    if (diedStr != "null")
    {
        string[] dd = diedStr.Split('/');
        int dMonth = int.Parse(dd[0]);
        int dDay = int.Parse(dd[1]);
        int dYear = int.Parse(dd[2]);
        age = dYear - bYear;
        if (dMonth < bMonth || (dMonth == bMonth && dDay < bDay))
            age--;
    }
    else
    {
        age = referenceDate.Year - bYear;
        if (referenceDate.Month < bMonth || (referenceDate.Month == bMonth && referenceDate.Day < bDay))
            age--;
    }

    // Build relatives
    var relatives = new ArrayNode();

    void AddRelative(string full, string relationship)
    {
        if (full == "null") return;
        string[] parts = full.Split(' ');
        var rel = new JsonObject
        {
            ["FirstName"] = parts[0],
            ["LastName"] = parts[parts.Length - 1],
            ["Relationship"] = relationship
        };
        relatives.Append(rel);
    }

    AddRelative(fatherStr, "Father");
    AddRelative(motherStr, "Mother");
    AddRelative(brotherStr, "Brother");
    AddRelative(sisterStr, "Sister");

    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayIso,
        ["Age"] = age,
        ["Relatives"] = relatives
    };

    people.Add(person);
}

var output = new ArrayNode();
foreach (var p in people)
    output.Append(p);

string json = output.ToJsonString(new JsonSerializerOptions { WriteIndented = true });
Console.Write(json);