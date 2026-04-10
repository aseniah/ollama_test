using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text.Json;

// Define the data structure to match the JSON object fields
// Using lowercase names to match the input JSON keys directly
record Person(string name, int age, bool active, float score);

string filePath = "input/data.json";

if (File.Exists(filePath))
{
    try
    {
        // Read the raw JSON content from the file
        string jsonContent = File.ReadAllText(filePath);

        // Configure deserialization to be case-insensitive to ensure robust mapping
        var deserializeOptions = new JsonSerializerOptions
        {
            PropertyNameCaseInsensitive = true
        };

        // Parse the JSON array into a list of Person objects
        var people = JsonSerializer.Deserialize<List<Person>>(jsonContent, deserializeOptions);

        if (people != null)
        {
            // Apply the filtering logic: active must be true and age must be 30 or older
            // Then sort the resulting collection by the 'name' field in ascending order
            var processedList = people
                .Where(p => p.active && p.age >= 30)
                .OrderBy(p => p.name)
                .ToList();

            // Serialize the processed list back to a JSON string
            // Use WriteIndented for readable output, though a single line would also be valid JSON
            var serializeOptions = new JsonSerializerOptions
            {
                WriteIndented = true
            };

            string outputJson = JsonSerializer.Serialize(processedList, serializeOptions);

            // Write the final JSON array to stdout
            Console.WriteLine(outputJson);
        }
    }
    catch
    {
        // If any error occurs during reading or parsing, we exit silently 
        // to adhere to the "Do not output anything else" constraint.
    }
}