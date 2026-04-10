using System;
using System.IO;
using System.Collections.Generic;
using System.Globalization;
using System.Text.Json;

string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    return;
}

string[] lines = File.ReadAllLines(filePath);
var jsonList = new List<object>();

// Skip the first line (header)
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    // Split by comma. Note: This assumes no commas are escaped within values.
    string[] columns = line.Split(',');

    if (columns.Length >= 4)
    {
        try
        {
            // Parse values using InvariantCulture to ensure '.' is the decimal separator
            string name = columns[0].Trim();
            int age = int.Parse(columns[1].Trim(), CultureInfo.InvariantCulture);
            string email = columns[2].Trim();
            float score = float.Parse(columns[3].Trim(), CultureInfo.InvariantCulture);

            jsonList.Add(new
            {
                Name = name,
                Age = age,
                Email = email,
                Score = score
            });
        }
        catch (Exception)
        {
            // In a production environment, you might log the error. 
            // Here we skip malformed rows to ensure valid JSON output.
            continue;
        }
    }
}

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

string jsonOutput = JsonSerializer.Serialize(jsonList, options);
Console.Write(jsonOutput);