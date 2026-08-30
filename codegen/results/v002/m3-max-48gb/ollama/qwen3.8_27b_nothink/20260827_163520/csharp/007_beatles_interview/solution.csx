using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

string csvPath = "input/input.csv";
string lines = File.ReadAllText(csvPath);
string[] dataLines = lines.Split('\n', StringSplitOptions.RemoveEmptyEntries);

var result = new JsonArray();

DateTime referenceDate = new DateTime(2025, 7, 1);

for (int i = 1; i < dataLines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(dataLines[i])) continue;
    
    string[] fields = dataLines[i].Split(',');
    
    // Name, Birthday, Died, Father, Mother, Brother, Sister
    string fullName = fields[0].Trim();
    string birthdayStr = fields[1].Trim();
    string diedStr = fields[2].Trim();
    string fatherStr = fields[3].Trim();
    string motherStr = fields[4].Trim();
    string brotherStr = fields[5].Trim();
    string sisterStr = fields[6].Trim();
    
    // Split name into first and last
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[1];
    
    // Parse birthday - format is M/D/YYYY
    DateTime birthday = DateTime.Parse(birthdayStr);
    string birthdayFormatted = birthday.ToString("yyyy-MM-dd");
    
    // Calculate age as of July 1, 2025
    DateTime deathDate = null;
    if (!diedStr.Equals("null") && !string.IsNullOrWhiteSpace(diedStr))
    {
        deathDate = DateTime.Parse(diedStr);
    }
    
    int age;
    if (deathDate.HasValue)
    {
        // Age at death
        age = (int)(deathDate.Value - birthday).TotalDays / 365;
    }
    else
    {
        // Age as of reference date
        if (referenceDate.Year == birthday.Year)
            age = (referenceDate.Month > birthday.Month || 
                   (referenceDate.Month == birthday.Month && referenceDate.Day >= birthday.Day)) ? 0 : 0;
        else
        {
            age = referenceDate.Year - birthday.Year;
            if (referenceDate.Month < birthday.Month || 
                (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
            {
                age--;
            }
        }
    }
    
    // Build relatives array
    var relatives = new JsonArray();
    
    if (!fatherStr.Equals("null") && !string.IsNullOrWhiteSpace(fatherStr))
    {
        string[] fParts = fatherStr.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = fParts[0],
            ["LastName"] = fParts[1],
            ["Relationship"] = "Father"
        });
    }
    
    if (!motherStr.Equals("null") && !string.IsNullOrWhiteSpace(motherStr))
    {
        string[] mParts = motherStr.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = mParts[0],
            ["LastName"] = mParts[1],
            ["Relationship"] = "Mother"
        });
    }
    
    if (!brotherStr.Equals("null") && !string.IsNullOrWhiteSpace(brotherStr))
    {
        string[] bParts = brotherStr.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = bParts[0],
            ["LastName"] = bParts[1],
            ["Relationship"] = "Brother"
        });
    }
    
    if (!sisterStr.Equals("null") && !string.IsNullOrWhiteSpace(sisterStr))
    {
        string[] sParts = sisterStr.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = sParts[0],
            ["LastName"] = sParts[1],
            ["Relationship"] = "Sister"
        });
    }
    
    var person = new JsonObject {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayFormatted,
        ["Age"] = age,
        ["Relatives"] = relatives
    };
    
    result.Add(person);
}

Console.WriteLine(result.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));