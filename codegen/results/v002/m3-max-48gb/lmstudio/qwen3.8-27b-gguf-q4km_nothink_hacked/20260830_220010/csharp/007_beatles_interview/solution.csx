using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string[] lines = File.ReadAllLines("input/input.csv");
string[] headers = lines[0].Split(',');

var referenceDate = new DateTime(2025, 7, 1);
var results = new System.Collections.Generic.List<JsonNode>();

for (int i = 1; i < lines.Length; i++)
{
    var fields = lines[i].Split(',');
    
    string fullName = fields[0].Trim();
    string birthdayStr = fields[1].Trim();
    string diedStr = fields[2].Trim();
    string fatherStr = fields[3].Trim();
    string motherStr = fields[4].Trim();
    string brotherStr = fields[5].Trim();
    string sisterStr = fields[6].Trim();
    
    // Parse name
    var nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = string.Join(" ", nameParts.Skip(1));
    
    // Parse birthday (MM/DD/YYYY)
    DateTime birthday;
    if (DateTime.TryParseExact(birthdayStr, "M/d/yyyy", null, System.Globalization.DateTimeStyles.None, out birthday))
    {
        // ok
    }
    else if (DateTime.TryParseExact(birthdayStr, "M/dd/yyyy", null, System.Globalization.DateTimeStyles.None, out birthday))
    {
        // ok
    }
    else if (DateTime.TryParseExact(birthdayStr, "MM/dd/yyyy", null, System.Globalization.DateTimeStyles.None, out birthday))
    {
        // ok
    }
    else if (DateTime.TryParseExact(birthdayStr, "MM/DD/yyyy", null, System.Globalization.DateTimeStyles.None, out birthday))
    {
        // ok
    }
    else
    {
        birthday = DateTime.Parse(birthdayStr);
    }
    
    string birthdayIso = birthday.ToString("yyyy-MM-dd");
    
    // Calculate age
    int age;
    if (diedStr != "null" && diedStr != "")
    {
        DateTime died;
        if (!DateTime.TryParseExact(diedStr, "M/d/yyyy", null, System.Globalization.DateTimeStyles.None, out died))
        {
            if (!DateTime.TryParseExact(diedStr, "M/dd/yyyy", null, System.Globalization.DateTimeStyles.None, out died))
            {
                died = DateTime.Parse(diedStr);
            }
        }
        age = died.Year - birthday.Year;
    }
    else
    {
        age = referenceDate.Year - birthday.Year;
        // Check if birthday hasn't occurred yet this year
        var thisYearBirthday = new DateTime(referenceDate.Year, birthday.Month, birthday.Day);
        if (thisYearBirthday > referenceDate)
        {
            age--;
        }
    }
    
    // Build relatives
    var relatives = new System.Collections.Generic.List<JsonNode>();
    
    if (fatherStr != "null" && fatherStr != "")
    {
        var namePartsF = fatherStr.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = namePartsF[0],
            ["LastName"] = string.Join(" ", namePartsF.Skip(1)),
            ["Relationship"] = "Father"
        });
    }
    
    if (motherStr != "null" && motherStr != "")
    {
        var namePartsM = motherStr.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = namePartsM[0],
            ["LastName"] = string.Join(" ", namePartsM.Skip(1)),
            ["Relationship"] = "Mother"
        });
    }
    
    if (brotherStr != "null" && brotherStr != "")
    {
        var namePartsB = brotherStr.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = namePartsB[0],
            ["LastName"] = string.Join(" ", namePartsB.Skip(1)),
            ["Relationship"] = "Brother"
        });
    }
    
    if (sisterStr != "null" && sisterStr != "")
    {
        var namePartsS = sisterStr.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = namePartsS[0],
            ["LastName"] = string.Join(" ", namePartsS.Skip(1)),
            ["Relationship"] = "Sister"
        });
    }
    
    var person = new JsonObject {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayIso,
        ["Age"] = age,
        ["Relatives"] = new JsonArray(relatives.ToArray())
    };
    
    results.Add(person);
}

var output = new JsonArray(results.ToArray());
string json = output.ToJsonString(new JsonSerializerOptions { WriteIndented = true });
Console.WriteLine(json);