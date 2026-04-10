using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Linq;

class Program
{
    static void Main()
    {
        // Read the JSON file
        string filePath = "input/data.json";
        string jsonContent = File.ReadAllText(filePath);

        // Parse the JSON content
        JArray jsonArray = JArray.Parse(jsonContent);

        // Filter records where active is true and age is 30 or older
        JArray filteredData = jsonArray.Where(item => item["active"].GetBoolean() && item["age"].GetInt32() >= 30).ToArray();

        // Sort the filtered data by name in ascending order
        JArray sortedData = filteredData.OrderBy(item => item["name"].GetString()).ToArray();

        // Output the sorted JSON array
        Console.WriteLine(JsonConvert.SerializeObject(sortedData, Formatting.Indented));
    }
}