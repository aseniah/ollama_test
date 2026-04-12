using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the input CSV file
// The file is located at input/input.csv
string inputPath = "input/input.csv";

if (!File.Exists(inputPath))
{
    return;
}

string[] lines = File.ReadAllLines(inputPath);
JsonArray jsonArray = new JsonArray();
DateTime targetDate = new DateTime(2025, 7, 1);

// The CSV has a header: Name,Birthday,Died,Father,Mother,Brother,Sister
// We skip the first line (the header)
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    // Parse CSV columns
    string[] columns = line.Split(',');
    if (columns.Length < 7) continue;

    // 1. Handle Name transformation (First and Last)
    // Logic: Use the first part as FirstName and the last part as LastName
    string fullName = columns[0].Trim();
    string[] nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts.Length > 0 ? nameParts[0] : "";
    string lastName = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";

    // 2. Handle Birthday transformation (to YYYY-MM-DD)
    // The CSV format is M/D/YYYY
    DateTime birthday;
    bool parseSuccess = DateTime.TryParseExact(columns[1].Trim(), "M/d/yyyy", CultureInfo.InvariantCulture, DateTimeStyles.None, out birthday);
    if (!parseSuccess) continue;
    string birthdayFormatted = birthday.ToString("yyyy-MM-dd");

    // 3. Calculate Age as of July 1, 2025
    int age = targetDate.Year - birthday.Year;
    if (targetDate < birthday.AddYears(age))
    {
        age--;
    }

    // 4. Construct the main Person object
    JsonObject personNode = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayFormatted,
        ["Age"] = age,
        ["Relatives"] = new JsonArray()
    };

    JsonArray relativesArray = (JsonArray)personNode["Relatives"];

    // 5. Process Relatives (Father, Mother, Brother, Sister)
    // Indices in CSV: Father=3, Mother=4, Brother=5, Sister=6
    string[] relativeNames = { columns[3], columns[4], columns[5], columns[6] };
    string[] relationships = { "Father", "Mother", "Brother", "Sister" };

    for (int j = 0; j < relativeNames.Length; j++)
    {
        string relNameStr = relativeNames[j].Trim();
        
        // Skip if the value is 'null' or empty
        if (string.IsNullOrEmpty(relNameStr) || relNameStr.Equals("null", StringComparison.OrdinalIgnoreCase))
        {
            continue;
        }

        string[] rParts = relNameStr.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        string rFirstName = rParts.Length > 0 ? rParts[0] : "";
        string rLastName = rParts.Length > 1 ? rParts[rParts.Length - 1] : "";

        relativesArray.Add(new JsonObject
        {
            ["FirstName"] = rFirstName,
            ["LastName"] = rLastName,
            ["Relationship"] = relationships[j]
        });
    }

    jsonArray.Add(personNode);
}

// Output the final JSON array to stdout
// Using Indented formatting to match the style of expected_format.json
Console.WriteLine(jsonArray.ToString(new JsonSerializerOptions { WriteIndented = true }));