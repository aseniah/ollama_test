using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Collections.Generic;

// Read the JSON file from the specified path
string jsonString = File.ReadAllText("input/data.json");

// Deserialize the JSON array into a list of Record objects
var data = JsonSerializer.Deserialize<List<Record>>(jsonString);

// Filter records: active must be true and age must be 30 or older
// Then sort the resulting list by name in ascending order
var filteredData = data
    .Where(r => r.active && r.age >= 30)
    .OrderBy(r => r.name)
    .ToList();

// Serialize the filtered result back to a JSON string and write to stdout
Console.WriteLine(JsonSerializer.Serialize(filteredData));

// Define the data structure matching the JSON object fields
public class Record
{
    public string name { get; set; }
    public int age { get; set; }
    public bool active { get; set; }
    public float score { get; set; }
}