using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Collections.Generic;

// Define a simple class to represent the data structure
public class Record
{
    public string name { get; set; }
    public int age { get; set; }
    public bool active { get; set; }
    public float score { get; set; }
}

string filePath = "input/data.json";

// Ensure the file exists before attempting to read
if (File.Exists(filePath))
{
    try
    {
        string jsonString = File.ReadAllText(filePath);
        
        // Deserialize the JSON array into a list of Record objects
        // PropertyNameCaseInsensitive is used to be robust against casing differences
        var options = new JsonSerializerOptions { PropertyNameCaseInsensitive = true };
        var records = JsonSerializer.Deserialize<List<Record>>(jsonString, options);

        if (records != null)
        {
            // Apply filtering logic: active is true and age is 30 or older
            // Apply sorting logic: sort by name ascending
            var processedRecords = records
                .Where(r => r.active && r.age >= 30)
                .OrderBy(r => r.name)
                .ToList();

            // Serialize the result back to a JSON array and write to stdout
            // Using default settings (no indentation) to keep output clean as requested
            Console.WriteLine(JsonSerializer.Serialize(processedRecords));
        }
    }
    catch
    {
        // In a real-world scenario, error handling would go here.
        // Per instructions, we output nothing else.
    }
}