using System;
using System.IO;
using System.Text.Json;

// Define a function to read and process JSON data
void ProcessData(string filePath)
{
    try
    {
        // Read all lines from the JSON file
        string jsonData = File.ReadAllLines(filePath).Aggregate(string.Empty, (accumulator, line) => accumulator + line + Environment.NewLine);

        // Parse the JSON data into a JSON array
        var jsonArray = JsonSerializer.Deserialize<JsonArray>(jsonData);

        // Filter records where active is true and age is 30 or older
        var filteredRecords = jsonArray.Select(record => new { name = record.GetProperty("name").GetString(), age = record.GetProperty("age").GetInt32(), active = record.GetProperty("active").GetBool(), score = record.GetProperty("score").GetFloat() });

        // Sort records by name in ascending order
        filteredRecords.Sort((a, b) => string.Compare(a.name, b.name));

        // Serialize the filtered and sorted records back to JSON
        string jsonOutput = JsonSerializer.Serialize(filteredRecords, new JsonSerializerOptions { WriteIndented = true });

        // Output the result to stdout
        Console.WriteLine(jsonOutput);
    }
    catch (Exception ex)
    {
        Console.WriteLine($"An error occurred: {ex.Message}");
    }
}

// Execute the function with the file path
ProcessData("input/data.json");