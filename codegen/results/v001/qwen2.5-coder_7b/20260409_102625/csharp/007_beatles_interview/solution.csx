using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

class Script
{
    public static void Main()
    {
        var lines = File.ReadAllLines("input/input.csv");
        var output = new List<JsonNode>();

        foreach (var line in lines)
        {
            if (string.IsNullOrWhiteSpace(line) || line.StartsWith("#"))
                continue;

            var parts = line.Split(',');
            if (parts.Length < 5)
                continue;

            var name = parts[0];
            var birthDateStr = parts[1];
            var address = parts[3];
            var phoneNumber = parts[4];

            if (!DateTime.TryParse(birthDateStr, out DateTime birthDate))
                continue;

            var age = CalculateAge(birthDate, new DateTime(2025, 7, 1));

            var jsonNode = JsonNode.Parse($@"{{""name"": ""{name}"", ""age"": {age}, ""address"": ""{address}"", ""phoneNumber"": ""{phoneNumber}""}}");
            output.Add(jsonNode);
        }

        Console.WriteLine(JsonSerializer.Serialize(output));
    }

    private static int CalculateAge(DateTime birthDate, DateTime today)
    {
        int age = today.Year - birthDate.Year;
        if (today < birthDate.AddYears(age))
            age--;
        return age;
    }
}