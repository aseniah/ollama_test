using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Linq;

public class JsonProcessor
{
    public static void ProcessCSVFile(string filePath)
    {
        try
        {
            var lines = File.ReadAllLines(filePath);
            var data = lines.Skip(1).ToList(); // Skip header

            var jArray = new JArray();

            foreach (var line in data)
            {
                string[] fields = line.Split(',');
                if (fields.Length < 2) continue;

                var person = new JObject
                {
                    [JToken.Name("firstName")] = fields[0].Trim(),
                    [JToken.Name("lastName")] = fields[1].Trim(),
                    [JToken.Name("age")] = CalculateAge(fields[2])
                };

                jArray.Add(person);
            }

            var jsonString = JObject.ToString(jArray);
            Console.WriteLine(jsonString);
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error processing file: {ex.Message}");
        }
    }

    private static int CalculateAge(string birthYear)
    {
        int year = 2025;
        int birthYearInt = int.Parse(birthYear);
        if (birthYearInt < year)
            throw new ArgumentException("Invalid birth year");
        return year - birthYearInt;
    }
}

JsonProcessor.ProcessCSVFile("input/input.csv");