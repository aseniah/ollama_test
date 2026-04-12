using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Collections.Generic;

// Define the structure of the input data
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
    // If the file doesn't exist, we output nothing as per instructions or handle error.
    // However, the prompt implies the file exists.
    return;
}

try
{
    // Read the content from the file
    string jsonContent = File.ReadAllText(filePath);

    // Deserialize the JSON array into a list of Person objects
    var options = new JsonSerializerOptions
    {
        PropertyNameCaseInsensitive = true
    };
    
    List<Person> people = JsonSerializer.Deserialize<List<Person>>(jsonContent, options);

    if (people != null)
    {
        // Filter: active is true AND age is 30 or older
        // Sort: name ascending
        var filteredPeople = people
            .Where(p => p.active && p.age >= 30)
            .OrderBy(p => p.name)
            .ToList();

        // Serialize the resulting list back to JSON
        // Using WriteIndented for readability, though the prompt asks for the array.
        string outputJson = JsonSerializer.Serialize(filteredPeople, new JsonSerializerOptions 
        { 
            WriteIndented = true 
        });

        // Output to stdout
        Console.Write(outputJson);
    }
}
catch (Exception)
{
    // If there's an error parsing, we follow the instruction "Do not output anything else"
    // by staying silent.
}