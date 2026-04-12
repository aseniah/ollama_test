using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;

// Define the structure of the JSON objects
public class Person
{
    public string name { get; set; }
    public int age { get; set; }
    public bool active { get; set; }
    public float score { get; set; }
}

string filePath = "input/data.json";

if (File.Exists(filePath))
{
    string jsonString = File.ReadAllText(filePath);
    
    // Deserialize the JSON array into a list of Person objects
    var options = new JsonSerializerOptions
    {
        PropertyNameCaseInsensitive = true
    };
    
    List<Person> people = JsonSerializer.Deserialize<List<Person>>(jsonString, options);

    if (people != null)
    {
        // Filter: active must be true AND age must be 30 or older
        // Sort: by name ascending
        var filteredPeople = people
            .Where(p => p.active && p.age >= 30)
            .OrderBy(p => p.name)
            .ToList();

        // Serialize the filtered list back to JSON
        string resultJson = JsonSerializer.Serialize(filteredPeople, new JsonSerializerOptions 
        { 
            WriteIndented = true 
        });

        // Output to stdout
        Console.Write(resultJson);
    }
}