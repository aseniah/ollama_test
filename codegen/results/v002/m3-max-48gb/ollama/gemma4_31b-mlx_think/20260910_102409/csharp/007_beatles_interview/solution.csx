using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Globalization;

// Setup target date
DateTime referenceDate = new DateTime(2025, 7, 1);

// Read CSV file
string csvContent = File.ReadAllText("input/input.csv");
string[] lines = csvContent.Split(new[] { '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

if (lines.Length == 0)
{
    Console.WriteLine("[]");
    return;
}

// Parse header
string[] headers = lines[0].Split(',');

// JSON output array
JsonArray rootArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    string[] columns = lines[i].Split(',');
    if (columns.Length < headers.Length) continue;

    // Map CSV columns to dictionary for easy access
    var data = new Dictionary<string, string>();
    for (int j = 0; j < headers.Length; j++)
    {
        data[headers[j]] = columns[j].Trim();
    }

    // Helper to extract first and last name
    Func<string, (string first, string last)> splitName = (name) =>
    {
        if (string.IsNullOrWhiteSpace(name) || name.Equals("null", StringComparison.OrdinalIgnoreCase))
            return (null, null);
        
        string[] parts = name.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        if (parts.Length == 0) return (null, null);
        if (parts.Length == 1) return (parts[0], "");
        return (parts[0], parts[parts.Length - 1]);
    };

    // Parse Dates
    DateTime birthday = DateTime.ParseExact(data["Birthday"], "M/d/yyyy", CultureInfo.InvariantCulture);
    DateTime? deathDate = null;
    if (!string.IsNullOrEmpty(data["Died"]) && !data["Died"].Equals("null", StringComparison.OrdinalIgnoreCase))
    {
        deathDate = DateTime.ParseExact(data["Died"], "M/d/yyyy", CultureInfo.InvariantCulture);
    }

    // Calculate Age
    DateTime endDate = deathDate ?? referenceDate;
    int age = endDate.Year - birthday.Year;
    if (endDate < new DateTime(endDate.Year, birthday.Month, birthday.Day))
    {
        age--;
    }

    var (firstName, lastName) = splitName(data["Name"]);

    // Build person object
    JsonObject person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age
    };

    // Handle relatives
    JsonArray relatives = new JsonArray();
    string[] relativeCols = new[] { "Father", "Mother", "Brother", "Sister" };
    foreach (var col in relativeCols)
    {
        string relativeName = data[col];
        if (!string.IsNullOrEmpty(relativeName) && !relativeName.Equals("null", StringComparison.OrdinalIgnoreCase))
        {
            var (relFirst, relLast) = splitName(relativeName);
            JsonObject relative = new JsonObject
            {
                ["FirstName"] = relFirst,
                ["LastName"] = relLast,
                ["Relationship"] = col
            };
            relatives.Add(relative);
        }
    }
    person["Relatives"] = relatives;

    rootArray.Add(person);
}

// Output JSON to stdout
var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(rootArray, options));