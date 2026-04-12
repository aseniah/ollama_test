using System;
using System.IO;
using System.Text.Json;

// Read data.json from the input directory and parse it into a JSON array.
var jsonData = File.ReadAllLines("input/data.json").ToArray();

// Parse the JSON array into a JSON array of objects.
var jsonArray = JsonDocument.Parse(new StringJsonDocument(jsonData));

// Filter records where active is true and age is 30 or older.
var filteredRecords = jsonArray.SelectMany(obj => new List<dynamic> { obj }).Where(record => record.active && record.age >= 30);

// Sort the filtered records by name in ascending order.
var sortedRecords = filteredRecords.OrderBy(record => record.name);

// Print the sorted records to stdout.
foreach (var record in sortedRecords)
{
    var name = record["name"];
    var age = record["age"];
    var active = record["active"];
    var score = record["score"];
    string jsonOutput = $"{{ \"name\": \"{name}\", \"age\": {age}, \"active\": {active}, \"score\": {score}}}";
    Console.WriteLine(jsonOutput);
}