#r "System.Text.Json"
using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;

var referenceDate = new DateTime(2025, 7, 1);
var lines = File.ReadAllLines("input/input.csv");
var result = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;

    var parts = line.Split(',');
    if (parts.Length < 7) continue;

    var fullName = parts[0].Trim();
    var birthdayStr = parts[1].Trim();
    var father = parts[3].Trim();
    var mother = parts[4].Trim();
    var brother = parts[5].Trim();
    var sister = parts[6].Trim();

    // Parse name into FirstName and LastName
    var nameParts = fullName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
    var firstName = nameParts[0];
    var lastName = nameParts[^1];

    // Parse birthday (MM/DD/YYYY format)
    var dateParts = birthdayStr.Split('/');
    var month = int.Parse(dateParts[0]);
    var day = int.Parse(dateParts[1]);
    var year = int.Parse(dateParts[2]);
    var birthday = new DateTime(year, month, day);
    var birthdayStrFormatted = $"{year:D4}-{month:D2}-{day:D2}";

    // Calculate age as of July 1, 2025
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate.Month < birthday.Month ||
        (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
    {
        age--;
    }

    // Build relatives array
    var relatives = new List<object>();

    if (!string.IsNullOrEmpty(father) && father != "null")
    {
        var fatherNameParts = father.Split(' ');
        relatives.Add(new { FirstName = fatherNameParts[0], LastName = fatherNameParts.Length > 1 ? fatherNameParts[^1] : "", Relationship = "Father" });
    }
    if (!string.IsNullOrEmpty(mother) && mother != "null")
    {
        var motherNameParts = mother.Split(' ');
        relatives.Add(new { FirstName = motherNameParts[0], LastName = motherNameParts.Length > 1 ? motherNameParts[^1] : "", Relationship = "Mother" });
    }
    if (!string.IsNullOrEmpty(brother) && brother != "null")
    {
        var brotherNameParts = brother.Split(' ');
        relatives.Add(new { FirstName = brotherNameParts[0], LastName = brotherNameParts.Length > 1 ? brotherNameParts[^1] : "", Relationship = "Brother" });
    }
    if (!string.IsNullOrEmpty(sister) && sister != "null")
    {
        var sisterNameParts = sister.Split(' ');
        relatives.Add(new { FirstName = sisterNameParts[0], LastName = sisterNameParts.Length > 1 ? sisterNameParts[^1] : "", Relationship = "Sister" });
    }

    result.Add(new { FirstName = firstName, LastName = lastName, Birthday = birthdayStrFormatted, Age = age, Relatives = relatives });
}

var options = new JsonSerializerOptions { WriteIndented = true };
var json = JsonSerializer.Serialize(result, options);
Console.WriteLine(json);