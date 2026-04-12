#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Serialization;

string referenceDate = "2025-07-01";
DateTime reference = DateTime.Parse(referenceDate);

string[] lines = File.ReadAllLines("input/input.csv");
string header = lines[0];
string[] headers = header.Split(',');

List<object> output = new List<object>();

for (int i = 1; i < lines.Length; i++) {
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    string[] fields = line.Split(',');
    
    // Parse name into FirstName and LastName
    string fullName = fields[0];
    int spaceIndex = fullName.LastIndexOf(' ');
    string firstName = spaceIndex > 0 ? fullName.Substring(0, spaceIndex) : fullName;
    string lastName = spaceIndex > 0 ? fullName.Substring(spaceIndex + 1) : "";
    
    // Parse birthday
    string birthdayStr = fields[1];
    DateTime birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", null);
    string birthdayFormatted = birthday.ToString("yyyy-MM-dd");
    
    // Calculate age (as of reference date or at death if deceased)
    int age;
    if (fields[2].ToLower() != "null" && !string.IsNullOrEmpty(fields[2])) {
        DateTime died = DateTime.ParseExact(fields[2], "M/d/yyyy", null);
        age = CalculateAge(birthday, died);
    } else {
        age = CalculateAge(birthday, reference);
    }
    
    // Parse relatives
    List<object> relatives = new List<object>();
    
    if (!string.IsNullOrEmpty(fields[3]) && fields[3].ToLower() != "null") {
        string[] parentName = fields[3].Split(' ', StringSplitOptions.TrimEntries);
        relatives.Add(new {
            FirstName = parentName.Length > 0 ? parentName[0] : "",
            LastName = parentName.Length > 1 ? parentName[^1] : "",
            Relationship = "Father"
        });
    }
    
    if (!string.IsNullOrEmpty(fields[4]) && fields[4].ToLower() != "null") {
        string[] parentName = fields[4].Split(' ', StringSplitOptions.TrimEntries);
        relatives.Add(new {
            FirstName = parentName.Length > 0 ? parentName[0] : "",
            LastName = parentName.Length > 1 ? parentName[^1] : "",
            Relationship = "Mother"
        });
    }
    
    if (!string.IsNullOrEmpty(fields[5]) && fields[5].ToLower() != "null") {
        string[] siblingName = fields[5].Split(' ', StringSplitOptions.TrimEntries);
        relatives.Add(new {
            FirstName = siblingName.Length > 0 ? siblingName[0] : "",
            LastName = siblingName.Length > 1 ? siblingName[^1] : "",
            Relationship = "Brother"
        });
    }
    
    if (!string.IsNullOrEmpty(fields[6]) && fields[6].ToLower() != "null") {
        string[] siblingName = fields[6].Split(' ', StringSplitOptions.TrimEntries);
        relatives.Add(new {
            FirstName = siblingName.Length > 0 ? siblingName[0] : "",
            LastName = siblingName.Length > 1 ? siblingName[^1] : "",
            Relationship = "Sister"
        });
    }
    
    output.Add(new {
        FirstName,
        LastName,
        Birthday = birthdayFormatted,
        Age = age,
        Relatives = relatives
    });
}

string jsonOutput = JsonSerializer.Serialize(output, new JsonSerializerOptions { WriteIndented = true });
Console.WriteLine(jsonOutput);

int CalculateAge(DateTime birth, DateTime end) {
    int years = end.Year - birth.Year;
    if (end.Month < birth.Month || (end.Month == birth.Month && end.Day < birth.Day)) {
        years--;
    }
    return years;
}