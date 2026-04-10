using System;
using System.IO;
using System.Text.Json;

// Assuming input.csv is in the same directory as this script
string inputFilePath = "input/input.csv";
string outputFilePath = "output.json";

// Read the CSV file and parse it into a List of dictionaries
List<Dictionary<string, object>> csvData = File.ReadAllLines(inputFilePath).Select(line =>
    JsonConvert.DeserializeObject<Dictionary<string, object>>(line)
).ToList();

// Calculate ages as of July 1, 2025
var transformedData = csvData.Select(row =>
{
    string dateStr = row["birthdate"].ToString();
    DateTime birthDate = DateTime.ParseExact(dateStr, "yyyy-MM-dd", CultureInfo.InvariantCulture, DateTimeStyles.None);
    DateTime currentDate = new DateTime(2025, 7, 1);
    
    // Calculate age
    long age = currentDate.Year - birthDate.Year - ((currentDate.Month < birthDate.Month) || (currentDate.Month == birthDate.Month && currentDate.Day < birthDate.Day));

    return new Dictionary<string, object>
    {
        { "name", row["name"].ToString() },
        { "age", age }
    };
}).ToList();

// Serialize the transformed data to JSON and write to a file
string json = JsonSerializer.Serialize(transformedData, new JsonSerializerOptions { WriteIndented = true });
File.WriteAllText(outputFilePath, json);

Console.WriteLine("Process completed. JSON output written to " + outputFilePath);