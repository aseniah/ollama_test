using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Collections.Generic;

// Define the structure to map the JSON objects
public class Person
{
    public string name { get; set; }
    public int age { get; set; }
    public bool active { get; set; }
    public float score { get; set; }
}

string filePath = "input/data.json";

if (!File.Exists(filePath))
{
    return;
}

try
{
    // Read the file content
    string jsonContent = File.ReadAllText(filePath);

    // Deserialize the JSON array into a list of Person objects
    var options = new JsonSerializerOptions
    {
        PropertyNameCaseInsensitive = true
    };
    
    var people = JsonSerializer.Deserialize<List<Person>>(jsonContent, options);

    if (people != null)
    {
        // Filter: active is true AND age is 30 or older
        // Sort: name ascending
        var filteredPeople = people
            .Where(p => p.active && p.age >= 30)
            .OrderBy(p => p.name)
            .ToList();

        // Serialize the filtered list back to JSON
        // Using WriteIndented to ensure clean output, though the prompt implies a standard JSON array
        string outputJson = JsonSerializer.Serialize(filteredPeople, new JsonSerializerOptions 
        { 
            WriteIndented = true 
        });

        // Write to stdout
        Console.WriteLine(outputJson);
    }
}
catch (Exception)
{
    // In case of error, the requirement specifies "Do not output anything else"
    // So we exit silently.
}