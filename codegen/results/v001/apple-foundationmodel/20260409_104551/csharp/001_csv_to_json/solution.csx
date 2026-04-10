using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Serialization;

class Program
{
    static void Main()
    {
        var jsonArray = new List<JsonObject>();

        try
        {
            // Read CSV file
            var lines = File.ReadAllLines("input/data.csv");

            // Skip header
            if (lines.Length > 1)
            {
                foreach (var line in lines.Skip(1))
                {
                    var fields = line.Split(',');
                    if (fields.Length == 4)
                    {
                        var name = fields[0].Trim();
                        var age = int.Parse(fields[1].Trim());
                        var email = fields[2].Trim();
                        var score = float.Parse(fields[3].Trim());

                        var jsonObject = JsonSerializer.Serialize(new
                        {
                            Name = name,
                            Age = age,
                            Email = email,
                            Score = score
                        });

                        jsonArray.Add(JsonNode.Parse(jsonObject));
                    }
                }
            }
        }
        catch (Exception ex)
        {
            Console.WriteLine("Error reading file: " + ex.Message);
        }

        // Output JSON array
        Console.WriteLine(JsonSerializer.Serialize(jsonArray));
    }
}