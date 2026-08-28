using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

// Read JSON data from a file
var jsonData = File.ReadAllLines("input/data.json").Select(Line => JsonSerializer.Deserialize<dynamic>(Line)).ToList();

// Filter and sort the records
var filteredRecords = jsonData.Where(record => record.Active && record.Age >= 30).OrderBy(record => record.Name).ToList();

// Output the filtered and sorted JSON array
foreach (var record in filteredRecords)
{
    JsonSerializer.Serialize(record).WriteLine();
}