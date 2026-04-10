using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Linq;

class Program
{
    static void Main(string[] args)
    {
        var inputFile = Path.Combine("input", "input.csv");
        var expectedFile = Path.Combine("input", "expected_format.json");

        var jsonArray = JsonDocument.ReadTree<JArray>(File.ReadAllText(inputFile));

        foreach (var person in jsonArray)
        {
            var age = CalculateAgeAsOf(person["birthdate"].ToObject<DateTime>());
            person["age"] = JObject.Create<int>(age);
        }

        var outputJson = JsonDocument.Serialize(jsonArray);
        Console.WriteLine(outputJson);
    }

    private static int CalculateAgeAsOf(DateTime birthdate)
    {
        var currentDate = DateTime.Parse("2025-07-01");
        return currentDate.Year - birthdate.Year - ((currentDate.Month, currentDate.Day) < (birthdate.Month, birthdate.Day));
    }
}