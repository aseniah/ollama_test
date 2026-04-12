using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;

// Define the data structure to match the JSON input
public class Person
{
    public string name { get; set; }
    public int age { get; set; }
    public bool active { get; set; }
    public float score { get; set; }
}

// Define the path to the input file
string filePath = "input/data.json";

// Check if the file exists before proceeding
if (File.Exists(filePath))
{
    try
    {
        // Read the entire content of the JSON file
        string jsonString = File.ReadAllText(filePath);

        // Configure JSON deserialization options
        // PropertyNameCaseInsensitive helps if there are minor casing mismatches
        var options = new JsonSerializerOptions
        {
            PropertyNameCaseInsensitive = true
        };

        // Deserialize the JSON array into a List of Person objects
        var people = JsonSerializer.Deserialize<List<Person>>(jsonString, options);

        if (people != null)
        {
            // Apply the filter: active must be true AND age must be 30 or older
            // Then sort the resulting collection by name in ascending order
            var filteredPeople = people
                .Where(p => p.active && p.age >= 30)
                .OrderBy(p => p.name)
                .ToList();

            // Serialize the filtered list back to a JSON string and output to stdout
            // We use a simple serialization without indentation to keep it compact
            // but the structure remains a valid JSON array.
            string resultJson = JsonSerializer.Serialize(filteredPeople);
            Console.WriteLine(resultJson);
        }
    }
    catch (Exception)
    {
        // In a production environment, we would handle specific exceptions like JsonException or IOException.
        // Per the requirement "Do not output anything else", we exit silently on error.
    }
}