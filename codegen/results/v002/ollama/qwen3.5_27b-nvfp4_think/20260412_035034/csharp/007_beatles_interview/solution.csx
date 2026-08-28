#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Serialization;

var referenceDate = new DateTime(2025, 7, 1);

var options = new JsonSerializerOptions
{
    WriteIndented = true,
    DefaultIgnoreCondition = JsonIgnoreCondition.WhenWritingNull
};

var results = new List<object>();

var lines = File.ReadAllLines("input/input.csv");

// Skip header line
for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    
    // Parse full name into FirstName and LastName
    var fullName = parts[0].Trim();
    var nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    var firstName = nameParts.First();
    var lastName = nameParts.Last();
    
    // Parse birthday - format is M/D/YYYY or MM/DD/YYYY
    var birthdayStr = parts[1].Trim();
    var birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    var formattedBirthday = birthday.ToString("yyyy-MM-dd");
    
    // Calculate age
    int age;
    if (parts[2].Trim() == "null" || string.IsNullOrEmpty(parts[2]))
    {
        // Still alive - calculate age as of July 1, 2025
        var years = referenceDate.Year - birthday.Year;
        if (referenceDate < birthday.AddYears(years))
        {
            years--;
        }
        age = years;
    }
    else
    {
        // Deceased - calculate age at death
        var diedStr = parts[2].Trim();
        var died = DateTime.ParseExact(diedStr, "M/d/yyyy", CultureInfo.InvariantCulture);
        var years = died.Year - birthday.Year;
        if (died < birthday.AddYears(years))
        {
            years--;
        }
        age = years;
    }
    
    // Build relatives list
    var relatives = new List<object>();
    
    // Father
    if (!string.IsNullOrEmpty(parts[3].Trim()) && parts[3].Trim() != "null")
    {
        relatives.Add(new { FirstName = parts[3].Trim(), LastName = "", Relationship = "Father" });
        // Parse father's name for proper last name if space exists
        var fatherParts = parts[3].Trim().Split(' ', StringSplitOptions.RemoveEmptyEntries);
        if (fatherParts.Length > 1)
        {
            relatives[^1] = new { FirstName = fatherParts[0], LastName = fatherParts.Skip(1).Join(" "), Relationship = "Father" };
        }
    }
    
    // Mother
    if (!string.IsNullOrEmpty(parts[4].Trim()) && parts[4].Trim() != "null")
    {
        relatives.Add(new { FirstName = parts[4].Trim(), LastName = "", Relationship = "Mother" });
        var motherParts = parts[4].Trim().Split(' ', StringSplitOptions.RemoveEmptyEntries);
        if (motherParts.Length > 1)
        {
            relatives[^1] = new { FirstName = motherParts[0], LastName = motherParts.Skip(1).Join(" "), Relationship = "Mother" };
        }
    }
    
    // Brother
    if (!string.IsNullOrEmpty(parts[5].Trim()) && parts[5].Trim() != "null")
    {
        relatives.Add(new { FirstName = parts[5].Trim(), LastName = "", Relationship = "Brother" });
        var brotherParts = parts[5].Trim().Split(' ', StringSplitOptions.RemoveEmptyEntries);
        if (brotherParts.Length > 1)
        {
            relatives[^1] = new { FirstName = brotherParts[0], LastName = brotherParts.Skip(1).Join(" "), Relationship = "Brother" };
        }
    }
    
    // Sister
    if (!string.IsNullOrEmpty(parts[6].Trim()) && parts[6].Trim() != "null")
    {
        relatives.Add(new { FirstName = parts[6].Trim(), LastName = "", Relationship = "Sister" });
        var sisterParts = parts[6].Trim().Split(' ', StringSplitOptions.RemoveEmptyEntries);
        if (sisterParts.Length > 1)
        {
            relatives[^1] = new { FirstName = sisterParts[0], LastName = sisterParts.Skip(1).Join(" "), Relationship = "Sister" };
        }
    }
    
    var person = new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = formattedBirthday,
        Age = age,
        Relatives = relatives
    };
    
    results.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(results, options));