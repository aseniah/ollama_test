using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text.Json;

// Define a class to represent the structure of the objects in the JSON array.
// We use properties that match the JSON keys exactly.
public class Person
{
    public string name { get; set; }
    public int age { get; set; }
    public bool active { get; set; }
    public float score { get; set; }
}

// The path to the input file as specified in the task.
string filePath = "input/data.json";

// Ensure the file exists before attempting to read it.
if (File.Exists(filePath))
{
    try
    {
        // Read the entire content of the JSON file.
        string jsonString = File.ReadAllText(filePath);

        // Deserialize the JSON array into a list of Person objects.
        var people = JsonSerializer.Deserialize<List<Person>>(jsonString);

        if (people != null)
        {
            // Perform the filtering and sorting:
            // 1. Filter: active must be true AND age must be 30 or older.
            // 2. Sort: Order by the name property in ascending order.
            var filteredResults = people
                .Where(p => p.active && p.age >= 30)
                .OrderBy(p => p.name)
                .ToList();

            // Configure serialization options for a clean, readable output.
            var serializeOptions = new JsonSerializerOptions
            {
                WriteIndented = true
            };

            // Output the resulting JSON array to stdout.
            string outputJson = JsonSerializer.Serialize(filteredResults, serializeOptions);
            Console.WriteLine(outputJson);
        }
    }
    catch (Exception)
    {
        // In a real-world production environment, you would handle specific exceptions.
        // For this task, we follow the instruction to output nothing else.
    }
}