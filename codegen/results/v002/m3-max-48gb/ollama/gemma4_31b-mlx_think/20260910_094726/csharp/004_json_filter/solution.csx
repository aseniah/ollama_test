using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;

// Define a record to map the JSON structure
public record Record(string name, int age, bool active, float score);

try
{
    // Read the JSON file
    string filePath = "input/data.json";
    if (!File.Exists(filePath))
    {
        return;
    }

    string jsonString = File.ReadAllText(filePath);

    // Deserialize the JSON array into a list of Record objects
    var data = JsonSerializer.Deserialize<List<Record>>(jsonString);

    if (data != null)
    {
        // Filter records: active is true AND age is 30 or older
        // Sort records: by name ascending
        var result = data
            .Where(r => r.active && r.age >= 30)
            .OrderBy(r => r.name)
            .ToList();

        // Output the result as a JSON array
        Console.WriteLine(JsonSerializer.Serialize(result));
    }
}
catch (Exception)
{
    // Silently handle or ignore errors as per "Do not output anything else" 
    // although normally you would log these.
}