using System;
using System.Collections.Generic;
using System.IO;
using System.Globalization;
using System.Text.Json;

// Define the path to the input file
string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    // If the file doesn't exist, we output an empty array to maintain valid JSON output
    Console.WriteLine("[]");
    return;
}

var results = new List<Dictionary<string, object>>();

try
{
    using (var reader = new StreamReader(filePath))
    {
        // Read the header line
        string headerLine = reader.ReadLine();
        if (headerLine == null)
        {
            Console.WriteLine("[]");
            return;
        }

        string[] headers = headerLine.Split(',');

        // Read the data rows
        while (!reader.EndOfStream)
        {
            string line = reader.ReadLine();
            if (string.IsNullOrWhiteSpace(line)) continue;

            string[] values = line.Split(',');

            // Ensure we have enough columns to map to our required fields
            // Expecting: Name, Age, Email, Score (at least 4 columns)
            if (values.Length >= 4)
            {
                var rowObject = new Dictionary<string, object>();

                // Parse Name (string)
                rowObject["Name"] = values[0].Trim();

                // Parse Age (int)
                if (int.TryParse(values[1].Trim(), NumberStyles.Integer, CultureInfo.InvariantCulture, out int age))
                {
                    rowObject["Age"] = age;
                }
                else
                {
                    rowObject["Age"] = 0; 
                }

                // Parse Email (string)
                rowObject["Email"] = values[2].Trim();

                // Parse Score (float/double)
                if (double.TryParse(values[3].Trim(), NumberStyles.Float, CultureInfo.InvariantCulture, out double score))
                {
                    rowObject["Score"] = score;
                }
                else
                {
                    rowObject["Score"] = 0.0;
                }

                results.Add(rowObject);
            }
        }
    }

    // Output the final JSON array to stdout
    var options = new JsonSerializerOptions 
    { 
        WriteIndented = true 
    };
    string jsonOutput = JsonSerializer.Serialize(results, options);
    Console.WriteLine(jsonOutput);
}
catch (Exception)
{
    // In case of error, output an empty array to ensure valid JSON structure is returned
    Console.WriteLine("[]");
}