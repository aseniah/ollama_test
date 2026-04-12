using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Linq;

class Program
{
    static void Main()
    {
        // Read the CSV file
        string csvContent = File.ReadAllLines("input/input.csv");

        // Parse the CSV content into a dictionary
        Dictionary<string, string> csvData = csvContent.Select(line =>
        {
            var fields = line.Split(',');
            return new
            {
                Name = fields[0],
                Birthday = fields[1],
                Died = fields[2],
                Father = fields[3],
                Mother = fields[4],
                Brother = fields[5],
                Sister = fields[6]
            }.ToString();
        }).ToDictionary(kvp => kvp.Key, kvp => kvp.Value);

        // Calculate ages and prepare the JSON array
        var jsonArray: List<JObject> = new List<JObject>();

        foreach (var entry in csvData)
        {
            var person = new JObject
            {
                "FirstName": entry.Key,
                "LastName": entry.Value.Split(' ')[1],
                "Birthday": DateTime.ParseExact(entry.Value.Split(' ')[0], "MM/dd/yyyy", CultureInfo.InvariantCulture),
                "Age": CalculateAge(DateTime.ParseExact(entry.Value.Split(' ')[0], "MM/dd/yyyy", CultureInfo.InvariantCulture)),
                "Relatives": GetRelatives(entry.Value)
            };
            jsonArray.Add(person);
        }

        // Output the JSON array
        var jsonString = JsonSerializer.Serialize(jsonArray, Formatting.Indented);
        Console.WriteLine(jsonString);
    }

    private static int CalculateAge(DateTime birthdate)
    {
        DateTime today = DateTime.Now;
        return (today.Year - birthdate.Year) - ((today.Month, today.Day) < (birthdate.Month, birthdate.Day));
    }

    private static List<JObject> GetRelatives(string relativesString)
    {
        return relativesString
            .Split('],')
            .Select(relative => new
            {
                FirstName = relative.Split('"')[1],
                LastName = relative.Split('"')[0].Split(' ')[0],
                Relationship = relative.Split('"')[2]
            })
            .ToList();
    }
}