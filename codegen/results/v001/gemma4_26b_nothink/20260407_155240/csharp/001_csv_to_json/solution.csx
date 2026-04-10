using System;
using System.IO;
using System.Collections.Generic;
using System.Globalization;
using System.Text.Json;

// Define the path to the input file
string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    // If file doesn't exist, we output an empty array to maintain valid JSON structure
    Console.WriteLine("[]");
    return;
}

var resultList = new List<Dictionary<string, object>>();

try
{
    using (var reader = new StreamReader(filePath))
    {
        // Read header line
        string headerLine = reader.ReadLine();
        if (headerLine == null)
        {
            Console.WriteLine("[]");
            return;
        }

        string[] headers = headerLine.Split(',');

        // Read data rows
        while (!reader.EndOfStream)
        {
            string line = reader.ReadLine();
            if (string.IsNullOrWhiteSpace(line)) continue;

            string[] values = line.Split(',');

            // We expect 4 columns based on requirements: Name, Age, Email, Score
            // We use a dictionary to represent the object fields
            var rowObject = new Dictionary<string, object>();

            // Mapping logic based on prompt requirements
            // Assuming CSV order: Name, Age, Email, Score
            if (values.Length >= 4)
            {
                rowObject["Name"] = values[0].Trim();
                
                if (int.TryParse(values[1].Trim(), NumberStyles.Integer, CultureInfo.InvariantCulture, out int age))
                    rowObject["Age"] = age;
                else
                    rowObject["Age"] = 0;

                rowObject["Email"] = values[2].Trim();

                if (float.TryParse(values[3].Trim(), NumberStyles.Float, CultureInfo.InvariantCulture, out float score))
                    rowObject["Score"] = score;
                else
                    rowObject["Score"] = 0.0f;

                resultList.Add(rowObject);
            }
        }
    }

    // Serialize the list to a JSON array string
    var options = new JsonSerializerOptions 
    { 
        WriteIndented = true 
    };
    string jsonOutput = JsonSerializer.Serialize(resultList, options);

    // Output only the valid JSON to stdout
    Console.Write(jsonOutput);
}
catch (Exception)
{
    // In case of any error, output an empty array to ensure valid JSON output requirement
    Console.Write("[]");
}