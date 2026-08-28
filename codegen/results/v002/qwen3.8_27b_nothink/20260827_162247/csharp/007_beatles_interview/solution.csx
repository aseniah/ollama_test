using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.Json;

string inputPath = "input/input.csv";
string[] lines = File.ReadAllLines(inputPath);

DateTime referenceDate = new DateTime(2025, 7, 1);

var people = new List<Dictionary<string, object>>();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    
    string[] fields = lines[i].Split(',');
    if (fields.Length < 7) continue;
    
    string fullName = fields[0].Trim();
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts.Length > 1 ? string.Join(" ", nameParts.Skip(1)) : "";
    
    // Parse birthday: format is M/D/YYYY
    DateTime birthday;
    if (!DateTime.TryParse(fields[1].Trim(), out birthday))
        continue;
    
    // Parse death date
    DateTime? deathDate = null;
    if (!string.Equals(fields[2].Trim(), "null", StringComparison.OrdinalIgnoreCase))
    {
        if (DateTime.TryParse(fields[2].Trim(), out var dd))
            deathDate = dd;
    }
    
    // Calculate age as of July 1, 2025
    int age;
    if (deathDate.HasValue && deathDate < referenceDate)
    {
        age = (int)(deathDate.Value.Year - birthday.Year);
        if (deathDate.Value.Month < birthday.Month || (deathDate.Value.Month == birthday.Month && deathDate.Value.Day < birthday.Day))
            age--;
    }
    else
    {
        age = (int)(referenceDate.Year - birthday.Year);
        if (referenceDate.Month < birthday.Month || (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
            age--;
    }
    
    var relatives = new List<Dictionary<string, string>>();
    
    // Father
    if (!string.Equals(fields[3].Trim(), "null", StringComparison.OrdinalIgnoreCase))
    {
        string fatherName = fields[3].Trim();
        string[] fParts = fatherName.Split(' ');
        relatives.Add(new Dictionary<string, string>
        {
            { "FirstName", fParts[0] },
            { "LastName", fParts.Length > 1 ? string.Join(" ", fParts.Skip(1)) : "" },
            { "Relationship", "Father" }
        });
    }
    
    // Mother
    if (!string.Equals(fields[4].Trim(), "null", StringComparison.OrdinalIgnoreCase))
    {
        string motherName = fields[4].Trim();
        string[] mParts = motherName.Split(' ');
        relatives.Add(new Dictionary<string, string>
        {
            { "FirstName", mParts[0] },
            { "LastName", mParts.Length > 1 ? string.Join(" ", mParts.Skip(1)) : "" },
            { "Relationship", "Mother" }
        });
    }
    
    // Brother
    if (!string.Equals(fields[5].Trim(), "null", StringComparison.OrdinalIgnoreCase))
    {
        string brotherName = fields[5].Trim();
        string[] bParts = brotherName.Split(' ');
        relatives.Add(new Dictionary<string, string>
        {
            { "FirstName", bParts[0] },
            { "LastName", bParts.Length > 1 ? string.Join(" ", bParts.Skip(1)) : "" },
            { "Relationship", "Brother" }
        });
    }
    
    // Sister
    if (!string.Equals(fields[6].Trim(), "null", StringComparison.OrdinalIgnoreCase))
    {
        string sisterName = fields[6].Trim();
        string[] sParts = sisterName.Split(' ');
        relatives.Add(new Dictionary<string, string>
        {
            { "FirstName", sParts[0] },
            { "LastName", sParts.Length > 1 ? string.Join(" ", sParts.Skip(1)) : "" },
            { "Relationship", "Sister" }
        });
    }
    
    people.Add(new Dictionary<string, object>
    {
        { "FirstName", firstName },
        { "LastName", lastName },
        { "Birthday", birthday.ToString("yyyy-MM-dd") },
        { "Age", age },
        { "Relatives", relatives }
    });
}

using var serializer = new System.Text.Json.JsonSerializerOptions
{
    WriteIndented = true
};

string json = JsonSerializer.Serialize(people, serializer);
Console.WriteLine(json);