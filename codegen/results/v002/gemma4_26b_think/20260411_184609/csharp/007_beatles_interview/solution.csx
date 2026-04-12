using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Define the path to the input CSV file
string csvPath = "input/input.csv";

// Verify if the file exists to prevent runtime errors
if (!File.Exists(csvPath))
{
    return;
    // Since we must output nothing other than the JSON, we just return.
}

// Read all lines from the CSV file
string[] lines = File.ReadAllLines(csvPath);
JsonArray root = new JsonArray();

// Iterate through each line of the CSV, skipping the header line
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line))
    {
        continue;
    }

    // Split the CSV line by comma
    string[] cols = line.Split(',');
    if (cols.Length < 7)
    {
        continue;
    }

    // Extract Name, Birthday, and Died columns
    string fullName = cols[0];
    string bdayStr = cols[1];
    string diedStr = cols[2];

    // Parse the Birthday using the M/d/yyyy format as observed in the input
    DateTime bday;
    if (!DateTime.TryParseExact(bdayStr, "M/d/yyyy", CultureInfo.InvariantCulture, DateTimeStyles.None, out bday))
    {
        continue;
    }

    // Parse the Died date if it is not "null"
    DateTime? died = null;
    if (diedStr != "null" && !string.IsNullOrWhiteSpace(diedStr))
    {
        if (DateTime.TryParseExact(diedStr, "M/d/yyyy", CultureInfo.InvariantCulture, DateTimeStyles.None, out DateTime d))
        {
            died = d;
        }
    }

    // Split the Name into First and Last name (First word and Last word)
    string[] nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    // Calculate Age: 
    // If Died is not null, calculate age at death.
    // If Died is null, calculate age as of July 1, 2025.
    DateTime referenceDate = died ?? new DateTime(2025, 7, 1);
    int age = referenceDate.Year - bday.Year;
    if (referenceDate < bday.AddYears(age))
    {
        age--;
    }

    // Construct the JSON object for the person
    JsonObject person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = bday.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = new JsonArray()
    };

    JsonArray relatives = (JsonArray)person["Relatives"];
    // The relative columns are Father, Mother, Brother, Sister
    string[] relCols = { "Father", "Mother", "Brother", "Sister" };

    // Iterate through columns representing relatives
    for (int j = 0; j < relCols.Length; j++)
    {
        int colIdx = 3 + j; // Index 3 is Father, 4 is Mother, etc.
        if (colIdx < cols.Length)
        {
            string relValue = cols[colIdx];
            // Check if the relative exists (not "null")
            if (relValue != "null" && !string.IsNullOrWhiteSpace(relValue))
            {
                string[] relNameParts = relValue.Split(' ', StringSplitOptions.RemoveEmptyEntries);
                if (relNameParts.Length > 0)
                {
                    relatives.Add(new JsonObject
                    {
                        ["FirstName"] = relNameParts[0],
                        ["LastName"] = relNameParts[relNameParts.Length - 1],
                        ["Relationship"] = relCols[j]
                    });
                }
            }
        }
    }

    // Add the person object to the main JSON array
    root.Add(person);
}

// Output the resulting JSON array to stdout, formatted for readability
var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(root.ToJsonString(options));