using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;

// Target date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

string inputPath = "input/input.csv";

if (!File.Exists(inputPath))
{
    return;
    // If the file doesn't exist, we can't do anything. 
    // However, per instructions, we should only output the JSON array.
}

string[] lines = File.ReadAllLines(inputPath);
if (lines.Length <= 1)
{
    Console.WriteLine("[]");
    return;
}

// Parse header
string[] headers = lines[0].Split(',');

var results = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] values = line.Split(',');
    
    // Map columns to a dictionary for easier lookup
    var row = new Dictionary<string, string>();
    for (int j = 0; j < headers.Length; j++)
    {
        row[headers[j].Trim()] = j < values.Length ? values[j].Trim() : null;
    }

    // Parse Full Name into First and Last
    // The input format shows "John Winston Lennon" -> "John" and "Lennon" 
    // but looking at the expected output:
    // John Winston Lennon -> FirstName: John, LastName: Lennon
    // James Paul McCartney -> FirstName: James, LastName: LastName
    // Ringo Starr -> FirstName: Ringo, LastName: Starr
    // It seems we take the first element as FirstName and the last as LastName.
    string fullName = row["Name"];
    string[] nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    if (nameParts.Length == 0) continue;

    string firstName = nameParts[0];
    string lastName = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";

    // Parse Birthday
    if (!DateTime.TryParseExact(row["Birthday"], "M/d/yyyy", CultureInfo.InvariantCulture, DateTimeStyles.None, out DateTime birthday))
    {
        // Fallback for different formats if necessary
        if (!DateTime.TryParse(row["Birthday"], out birthday)) continue;
    }

    // Calculate Age as of July 1, 2025
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate.Month < birthday.Month || (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
    {
        age--;
    }

    var relatives = new List<object>();

    // Helper to process relatives
    void AddRelative(string fullNameStr, string relationship)
    {
        if (string.IsNullOrEmpty(fullNameStr) || fullNameStr.Equals("null", StringComparison.OrdinalIgnoreCase))
            return;

        string[] parts = fullNameStr.Split(' ', StringSplitToptions.RemoveEmptyEntries);
        string relFirstName = parts[0];
        string relLastName = parts.Length > 1 ? parts[parts.Length - 1] : "";

        relatives.Add(new
        {
            FirstName = relFirstName,
            LastName = relLastName,
            Relationship = relationship
        });
    }

    AddRelative(row["Father"], "Father");
    AddRelative(row["Mother"], "Mother");
    AddRelative(row["Brother"], "Brother");
    AddRelative(row["Sister"], "Sister");

    results.Add(new
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

string jsonOutput = JsonSerializer.Serialize(results, options);
Console.Write(jsonOutput);