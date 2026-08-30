using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;

// Define the reference date as specified
DateTime referenceDate = new DateTime(2025, 7, 1);

string inputFilePath = "input/input.csv";

if (!File.Exists(inputFilePath))
{
    return;
    // If file doesn't exist, we can't do anything. 
    // But per instructions, we assume it exists for the task.
}

var lines = File.ReadAllLines(inputFilePath);
if (lines.Length <= 1)
{
    Console.Write("[]");
    return;
}

// Header: Name,Birthday,Died,Father,Mother,Brother,Sister
var headers = lines[0].Split(',');
var dataRows = lines.Skip(1);

var resultList = new List<object>();

foreach (var row in dataRows)
{
    if (string.IsNullOrWhiteSpace(row)) continue;

    var columns = row.Split(',');
    // Columns mapping:
    // 0: Name (Full Name)
    // 1: Birthday
    // 2: Died (not used in JSON but present)
    // 3: Father
    // 4: Mother
    // 5: Brother
    // 6: Sister

    string fullName = columns[0];
    string birthdayStr = columns[1];
    
    // Split FullName into First and Last
    // Logic: The expected format uses "John" and "Lennon" from "John Winston Lennon"
    // Looking at the expected JSON: "John Lennon" comes from "John Winston Lennon".
    // It seems it takes the first element and the last element.
    string[] nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";

    // Parse Birthday
    DateTime birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", CultureInfo.InvariantCulture);

    // Calculate Age as of July 1, 2025
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate.Month < birthday.Month || (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
    {
        age--;
    }

    var relatives = new List<object>();

    // Helper to add relative
    void AddRelative(string personData, string relationship)
    {
        if (string.IsNullOrWhiteSpace(personData) || personData.Equals("null", StringComparison.OrdinalIgnoreCase))
            return;

        string[] parts = personData.Split(' ');
        string relFirstName = parts[0];
        string relLastName = parts.Length > 1 ? parts[parts.Length - 1] : "";

        relatives.Add(new
        {
            FirstName = relFirstName,
            LastName = relLastName,
            Relationship = relationship
        });
    }

    // 3: Father, 4: Mother, 5: Brother, 6: Sister
    AddRelative(columns.Length > 3 ? columns[3] : null, "Father");
    AddRelative(columns.Length > 4 ? columns[4] : null, "Mother");
    AddRelative(columns.Length > 5 ? columns[5] : null, "Brother");
    AddRelative(columns.Length > 6 ? columns[6] : null, "Sister");

    resultList.Add(new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = birthday.ToString("yyyy-MM-dd"),
        Age = age,
        Relatives = relatives
    });
}

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

string jsonOutput = JsonSerializer.Serialize(resultList, options);
Console.WriteLine(jsonOutput);