using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string csvPath = "input/input.csv";
string[] lines = File.ReadAllLines(csvPath);

DateTime asOfDate = new DateTime(2025, 7, 1);

var people = new List<JsonObject>();

// Skip header
for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    string[] fields = lines[i].Split(',');
    
    string fullName = fields[0].Trim();
    string birthdayStr = fields[1].Trim();
    string diedStr = fields[2].Trim();
    string fatherStr = fields[3].Trim();
    string motherStr = fields[4].Trim();
    string brotherStr = fields[5].Trim();
    string sisterStr = fields[6].Trim();
    
    // Parse first and last name
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts.Length > 1 ? string.Join(" ", nameParts.Skip(1)) : "";
    
    // Parse birthday
    DateTime birthday;
    string[] bdParts = birthdayStr.Split('/');
    int bMonth = int.Parse(bdParts[0]);
    int bDay = int.Parse(bdParts[1]);
    int bYear = int.Parse(bdParts[2]);
    birthday = new DateTime(bYear, bMonth, bDay);
    
    // Calculate age as of July 1, 2025
    int age = asOfDate.Year - birthday.Year;
    if (asOfDate.Month < birthday.Month || (asOfDate.Month == birthday.Month && asOfDate.Day < birthday.Day))
    {
        age--;
    }
    
    // Build relatives list
    var relatives = new List<JsonObject>();
    
    if (!string.IsNullOrEmpty(fatherStr) && fatherStr != "null")
    {
        string[] fp = fatherStr.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = fp[0],
            ["LastName"] = fp.Length > 1 ? string.Join(" ", fp.Skip(1)) : "",
            ["Relationship"] = "Father"
        });
    }
    
    if (!string.IsNullOrEmpty(motherStr) && motherStr != "null")
    {
        string[] mp = motherStr.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = mp[0],
            ["LastName"] = mp.Length > 1 ? string.Join(" ", mp.Skip(1)) : "",
            ["Relationship"] = "Mother"
        });
    }
    
    if (!string.IsNullOrEmpty(brotherStr) && brotherStr != "null")
    {
        string[] bp = brotherStr.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = bp[0],
            ["LastName"] = bp.Length > 1 ? string.Join(" ", bp.Skip(1)) : "",
            ["Relationship"] = "Brother"
        });
    }
    
    if (!string.IsNullOrEmpty(sisterStr) && sisterStr != "null")
    {
        string[] sp = sisterStr.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = sp[0],
            ["LastName"] = sp.Length > 1 ? string.Join(" ", sp.Skip(1)) : "",
            ["Relationship"] = "Sister"
        });
    }
    
    var person = new JsonObject {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayStr,
        ["Age"] = age,
        ["Relatives"] = new JsonArray(relatives)
    };
    
    people.Add(person);
}

var result = new JsonArray(people);
string output = result.ToJsonString(new JsonSerializerOptions { WriteIndented = true });
Console.Write(output);