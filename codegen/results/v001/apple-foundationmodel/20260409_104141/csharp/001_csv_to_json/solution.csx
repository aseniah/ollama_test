using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Linq;

namespace CsvToJson
{
    class Program
    {
        static void Main()
        {
            // Read CSV file
            var csvFilePath = "input/data.csv";
            var csvLines = File.ReadAllLines(csvFilePath);
            
            // Skip header
            csvLines.RemoveAt(0);

            // Initialize JSON array
            var jsonArray: JArray = JArray.Create();

            // Process each line
            foreach (var line in csvLines)
            {
                // Split line into fields
                var fields = line.Split(',');
                var name = fields[0];
                var age = int.Parse(fields[1]);
                var email = fields[2];
                var score = float.Parse(fields[3]);

                // Create JSON object
                var jsonObject = JObject.Create(
                    JProperty("Name", name),
                    JProperty("Age", age),
                    JProperty("Email", email),
                    JProperty("Score", score)
                );

                // Add to JSON array
                jsonArray.Add(jsonObject);
            }

            // Output JSON array to stdout
            JsonSerializer.Serialize(jsonArray, Console.Out);
        }
    }
}