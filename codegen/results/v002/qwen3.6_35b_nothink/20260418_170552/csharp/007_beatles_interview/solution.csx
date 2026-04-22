using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read input CSV
string[] lines = File.ReadAllLines("input/input.csv");
// Skip header
List<string[]> dataLines = new List<string[]>();
for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    string[] cols = lines[i].Split(',');
    dataLines.Add(cols);
}

// Reference date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

// Helper to calculate age
int CalculateAge(DateTime birthDate, DateTime asOfDate)
{
    int age = asOfDate.Year - birthDate.Year;
    // If birthday hasn't occurred yet in the reference year, subtract 1
    if (asOfDate.Month < birthDate.Month || (asOfDate.Month == birthDate.Month && asOfDate.Day < birthDate.Day))
    {
        age--;
    }
    return age;
}

// Helper to format date as YYYY-MM-DD
string FormatDate(DateTime date)
{
    return date.ToString("yyyy-MM-dd");
}

// Helper to parse date from M/d/yyyy format
DateTime ParseDate(string dateStr)
{
    return DateTime.Parse(dateStr);
}

// Helper to split name into first and last name
// The last name is the last word, everything else is first name
void SplitName(string fullName, out string firstName, out string lastName)
{
    string[] parts = fullName.Split(' ');
    if (parts.Length < 2)
    {
        firstName = "";
        lastName = parts.Length > 0 ? parts[0] : "";
    }
    else
    {
        lastName = parts[parts.Length - 1];
        firstName = string.Join(" ", parts, 0, parts.Length - 1);
    }
}

// Build JSON array
JsonArray jsonArray = new JsonArray();

foreach (var row in dataLines)
{
    // Columns: Name, Birthday, Died, Father, Mother, Brother, Sister
    string name = row[0];
    string birthdayStr = row[1];
    // string diedStr = row[2]; // Not needed for output
    string father = row[3];
    string mother = row[4];
    string brother = row[5];
    string sister = row[6];

    // Parse birthday
    DateTime birthday = ParseDate(birthdayStr);
    string birthdayFormatted = FormatDate(birthday);
    int age = CalculateAge(birthday, referenceDate);

    // Split name
    string firstName, lastName;
    SplitName(name, out firstName, out lastName);

    // Create person object
    JsonObject person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayFormatted,
        ["Age"] = age,
        ["Relatives"] = new JsonArray()
    };

    // Add relatives if not null
    if (!string.IsNullOrEmpty(father) && father != "null")
    {
        string fFirst, fLast;
        SplitName(father, out fFirst, out fLast);
        JsonObject fatherObj = new JsonObject
        {
            ["FirstName"] = fFirst,
            ["LastName"] = fLast,
            ["Relationship"] = "Father"
        };
        person["Relatives"].Add(fatherObj);
    }

    if (!string.IsNullOrEmpty(mother) && mother != "null")
    {
        string mFirst, mLast;
        SplitName(mother, out mFirst, out mLast);
        JsonObject motherObj = new JsonObject
        {
            ["FirstName"] = mFirst,
            ["LastName"] = mLast,
            ["Relationship"] = "Mother"
        };
        person["Relatives"].Add(motherObj);
    }

    if (!string.IsNullOrEmpty(brother) && brother != "null")
    {
        string bFirst, bLast;
        SplitName(brother, out bFirst, out bLast);
        JsonObject brotherObj = new JsonObject
        {
            ["FirstName"] = bFirst,
            ["LastName"] = bLast,
            ["Relationship"] = "Brother"
        };
        person["Relatives"].Add(brotherObj);
    }

    if (!string.IsNullOrEmpty(sister) && sister != "null")
    {
        string sFirst, sLast;
        SplitName(sister, out sFirst, out sLast);
        JsonObject sisterObj = new JsonObject
        {
            ["FirstName"] = sFirst,
            ["LastName"] = sLast,
            ["Relationship"] = "Sister"
        };
        person["Relatives"].Add(sisterObj);
    }

    jsonArray.Add(person);
}

// Write to stdout
string json = jsonArray.ToJsonString(new JsonSerializerOptions { WriteIndented = true });
Console.WriteLine(json);