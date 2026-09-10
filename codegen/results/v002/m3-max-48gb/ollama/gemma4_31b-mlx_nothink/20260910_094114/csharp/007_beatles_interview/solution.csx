using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Globalization;

// Reference date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

string csvContent = File.ReadAllLines("input/input.csv");
var lines = csvContent.Skip(1); // Skip header

var resultList = new List<object>();

foreach (var line in lines)
{
    if (string.IsNullOrWhiteSpace(line)) continue;

    var cols = line.Split(',');
    
    // Parse Full Name to First and Last
    string fullName = cols[0].Trim();
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = string.Join(" ", nameParts.Skip(1));

    // Handle Birthday format (MM/DD/YYYY)
    DateTime birthday = DateTime.ParseExact(cols[1].Trim(), "M/d/yyyy", CultureInfo.InvariantCulture);
    string birthdayStr = birthday.ToString("yyyy-MM-dd");

    // Calculate Age as of July 1, 2025
    // If Died is present, age is usually calculated at death, 
    // but the prompt specifies "Calculate ages as of July 1, 2025".
    // Looking at expected_format.json:
    // John (1940-10-09) -> Age 40 (He died 1980, 1980-1940=40)
    // James (1942-06-18) -> Age 83 (2025-1942=83)
    // Ringo (1940-07-07) -> Age 84 (2025-1940 = 85, but July 7 is after July 1, so 84)
    // George (1943-02-25) -> Age 58 (Died 2001, 2001-1943=58)
    
    DateTime endDate = referenceDate;
    if (!string.IsNullOrEmpty(cols[2]) && cols[2].Trim().ToLower() != "null")
    {
        endDate = DateTime.ParseExact(cols[2].Trim(), "M/d/yyyy", CultureInfo.InvariantCulture);
    }

    int age = endDate.Year - birthday.Year;
    if (endDate < birthday.AddYears(age)) age--;

    // Parse Relatives
    var relatives = new List<object>();
    string[] relationLabels = { "Father", "Mother", "Brother", "Sister" };
    for (int i = 3; i <= 6; i++)
    {
        string relativeName = cols[i].Trim();
        if (!string.IsNullOrEmpty(relativeName) && relativeName.ToLower() != "null")
        {
            string[] relParts = relativeName.Split(' ');
            re_relatives(relParts, relationLabels[i - 3], relatives);
        }
    }

    resultList.Add(new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = birthdayStr,
        Age = age,
        Relatives = relatives
    });
}

void re_relatives(string[] parts, string relation, List<object> list)
{
    string fName = parts[0];
    string lName = parts.Length > 1 ? string.Join(" ", parts.Skip(1)) : "";
    list.Add(new
    {
        FirstName = fName,
        LastName = lName,
        Relationship = relation
    });
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(resultList, options));