using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text.Json;

// Define the data structure to match the JSON schema
public record Person(string name, int age, bool active, float score);

string filePath = "input/data.json";

// Ensure the file exists before attempting to read
if (File.Exists(filePath))
{
    try
    {
        string jsonString = File.ReadAllText(filePath);
        
        // Configure options to be case-insensitive to handle potential casing mismatches
        var options = new JsonSerializerOptions 
        { 
            PropertyNameCaseInsensitive = true 
        };

        var data = JsonSerializer.Deserialize<List<Person>>(jsonString, options);

        if (data != null)
        {
            // Apply filtering: active must be true AND age must be 30 or older
            // Apply sorting: name ascending
            var filteredResults = data
                .Where(p => p.active && p.age >= 30)
                .OrderBy(p => p.name)
                .ToList();

            // Output the result as a JSON array to stdout
            // Using WriteIndented for cleaner, readable output
            var outputOptions = new JsonSerializerOptions { WriteIndented = true };
            Console.WriteLine(JsonSerializer.Serialize(filteredResults, outputOptions));
        }
    }
    catch (Exception)
    {
        // In a production environment, we might log errors.
        // Per requirements, we output nothing else to stdout.
    }
}