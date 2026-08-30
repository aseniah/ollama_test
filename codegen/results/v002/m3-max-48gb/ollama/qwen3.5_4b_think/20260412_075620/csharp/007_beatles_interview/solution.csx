#r "System.IO"
using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;

// Read CSV file
string csvPath = "input/input.csv";
string[] csvLines = File.ReadAllLines(csvPath);

// Parse header to get column indices
var headers = csvLines[0].Split(',').Select(h => h.Trim());
var nameIdx = headers.IndexOf("Name");
var birthdayIdx = headers.IndexOf("Birthday");
var fatherIdx = headers.IndexOf("Father");
var motherIdx = headers.IndexOf("Mother");
var brotherIdx = headers.IndexOf("Brother");
var sisterIdx = headers.IndexOf("Sister");

// Parse each person
var people = new List<object>();

foreach (var line in csvLines.Skip(1))
{
    var columns = line.Split(',').Select(c => c.Trim());
    
    var name = columns[nameIdx].Trim();
    var birthdayStr = columns[birthdayIdx].Trim();
    
    // Parse birthday (M/D/YYYY format)
    var dateParts = birthdayStr.Split('/');
    var birthMonth = int.Parse(dateParts[0]);
    var birthDay = int.Parse(dateParts[1]);
    var birthYear = int.Parse(dateParts[2]);
    
    // Calculate age as of July 1, 2025
    var birthDate = new DateTime(birthYear, birthMonth, birthDay);
    var targetDate = new DateTime(2025, 7, 1);
    var age = (int)(targetDate.Year - birthDate.Year);
    if (targetDate < new DateTime(birthDate.Year, birthMonth, birthDay))
    {
        age--;
    }
    
    // Parse relatives
    var father = columns[fatherIdx].Trim();
    var mother = columns[motherIdx].Trim();
    var brother = columns[brotherIdx].Trim();
    var sister = columns[sisterIdx].Trim();
    
    var relatives = new List<object>();
    
    if (!string.IsNullOrEmpty(father) && father != "null")
    {
        var fParts = father.Split(',');
        var fFirstName = fParts[0].Trim();
        var fLastName = fParts[1].Trim();
        relatives.Add(new
        {
            FirstName = fFirstName,
            LastName = fLastName,
            Relationship = "Father"
        });
    }
    
    if (!string.IsNullOrEmpty(mother) && mother != "null")
    {
        var mParts = mother.Split(',');
        var mFirstName = mParts[0].Trim();
        var mLastName = mParts[1].Trim();
        relatives.Add(new
        {
            FirstName = mFirstName,
            LastName = mLastName,
            Relationship = "Mother"
        });
    }
    
    if (!string.IsNullOrEmpty(brother) && brother != "null")
    {
        var bParts = brother.Split(',');
        var bFirstName = bParts[0].Trim();
        var bLastName = bParts[1].Trim();
        relatives.Add(new
        {
            FirstName = bFirstName,
            LastName = bLastName,
            Relationship = "Brother"
        });
    }
    
    if (!string.IsNullOrEmpty(sister) && sister != "null")
    {
        var sParts = sister.Split(',');
        var sFirstName = sParts[0].Trim();
        var sLastName = sParts[1].Trim();
        relatives.Add(new
        {
            FirstName = sFirstName,
            LastName = sLastName,
            Relationship = "Sister"
        });
    }
    
    people.Add(new
    {
        FirstName = name,
        LastName = "", // Will be filled from name
        Birthday = $"{birthYear:D}-{birthMonth:D2}-{birthDay:D2}",
        Age = age,
        Relatives = relatives
    });
}

// Now I realize I need to split the full name into first and last name
// Looking at the input, the Name field contains "FirstName LastName"
// So I need to extract the last name from the full name

// Re-reading the CSV, let me fix the parsing
var people2 = new List<object>();

foreach (var line in csvLines.Skip(1))
{
    var columns = line.Split(',').Select(c => c.Trim());
    
    var fullName = columns[nameIdx];
    var parts = fullName.Split(' ');
    var firstName = parts[0];
    var lastName = parts.Length > 1 ? parts[1] : "";
    
    var birthdayStr = columns[birthdayIdx].Trim();
    var dateParts = birthdayStr.Split('/');
    var birthMonth = int.Parse(dateParts[0]);
    var birthDay = int.Parse(dateParts[1]);
    var birthYear = int.Parse(dateParts[2]);
    
    var birthDate = new DateTime(birthYear, birthMonth, birthDay);
    var targetDate = new DateTime(2025, 7, 1);
    var age = (int)(targetDate.Year - birthDate.Year);
    if (targetDate < new DateTime(birthDate.Year, birthMonth, birthDay))
    {
        age--;
    }
    
    var relatives = new List<object>();
    
    if (!string.IsNullOrEmpty(columns[fatherIdx]))
    {
        var fParts = columns[fatherIdx].Trim().Split(',');
        if (fParts.Length >= 2)
        {
            relatives.Add(new
            {
                FirstName = fParts[0].Trim(),
                LastName = fParts[1].Trim(),
                Relationship = "Father"
            });
        }
    }
    
    if (!string.IsNullOrEmpty(columns[motherIdx]))
    {
        var mParts = columns[motherIdx].Trim().Split(',');
        if (mParts.Length >= 2)
        {
            relatives.Add(new
            {
                FirstName = mParts[0].Trim(),
                LastName = mParts[1].Trim(),
                Relationship = "Mother"
            });
        }
    }
    
    if (!string.IsNullOrEmpty(columns[brotherIdx]))
    {
        var bParts = columns[brotherIdx].Trim().Split(',');
        if (bParts.Length >= 2)
        {
            relatives.Add(new
            {
                FirstName = bParts[0].Trim(),
                LastName = bParts[1].Trim(),
                Relationship = "Brother"
            });
        }
    }
    
    if (!string.IsNullOrEmpty(columns[sisterIdx]))
    {
        var sParts = columns[sisterIdx].Trim().Split(',');
        if (sParts.Length >= 2)
        {
            relatives.Add(new
            {
                FirstName = sParts[0].Trim(),
                LastName = sParts[1].Trim(),
                Relationship = "Sister"
            });
        }
    }
    
    people2.Add(new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = $"{birthYear:D}-{birthMonth:D2}-{birthDay:D2}",
        Age = age,
        Relatives = relatives
    });
}

// Convert to JSON
var json = JsonSerializer.Serialize(people2, new JsonSerializerOptions
{
    WriterOptions = new WriterOptions { Indented = true },
    PropertyNameCase = JsonNamingConvention.CamelCase
});

Console.WriteLine(json);