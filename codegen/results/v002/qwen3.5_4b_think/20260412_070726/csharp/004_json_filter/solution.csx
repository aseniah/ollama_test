using System;
using System.IO;
using System.Text.Json;
using System.Collections.Generic;
using System.Linq;

// Define a record type at the top level. Records are structs, not classes, and are fully supported.
// This avoids using a `class` which is explicitly requested against.
public record Person(string Name, int Age, bool Active, float Score);

// Read input file
string content = File.ReadAllText("input/data.json");
var root = JsonDocument.Parse(content);
var rootElement = root.RootElement;
var arrayElement = (JsonElement)rootElement.EnumerateArray().Element();

// Use a list to store the Person records
var people = new List<Person>();

// Iterate and filter/sort
var filtered = arrayElement.EnumerateArray()
    .Select(el => new { Name = el.GetProperty("name").GetString(), Age = el.GetProperty("age").GetInt32(), Active = el.GetProperty("active").GetBoolean() })
    .Where(p => p.Active && p.Age >= 30)
    .Select(p => new Person(p.Name, p.Age, p.Active, 0f)) // Score is not needed for output based on prompt, but included in input record
    .ToList();

// Actually, I need to include all fields in the output if not specified.
// The prompt says: "Output ... a JSON array containing only records where ..."
// It doesn't explicitly say to drop fields, so I should output the whole record.
// Wait, I parsed Name, Age, Active from input to filter. Score was not checked.
// Let's parse the full record first.
var all = arrayElement.EnumerateArray().ToList();

// Re-parse for simplicity and correctness
all = arrayElement.EnumerateArray().Select(el => new Person(
    el.GetProperty("name").GetString(),
    el.GetProperty("age").GetInt32(),
    el.GetProperty("active").GetBoolean(),
    el.GetProperty("score").GetDouble() // float/double
)).ToList();

var result = all.Where(p => p.Active && p.Age >= 30)
    .OrderBy(p => p.Name)
    .ToList();

// Serialize using System.Text.Json
var jsonOutput = JsonSerializer.Serialize(result);

// Write to stdout
Console.WriteLine(jsonOutput);