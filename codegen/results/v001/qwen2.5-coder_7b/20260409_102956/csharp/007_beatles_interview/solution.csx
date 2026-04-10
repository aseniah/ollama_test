using System;
using System.IO;
using System.Text.Json;

var inputFilePath = "input/input.csv";
var outputFilePath = "output/output.json";

if (!File.Exists(inputFilePath))
{
    Console.WriteLine("Input file does not exist.");
    return;
}

var lines = File.ReadAllLines(inputFilePath);
var jsonArray = new JsonArray();

foreach (var line in lines)
{
    var columns = line.Split(',');
    if (columns.Length < 3)
    {
        continue;
    }

    var birthDateStr = columns[1];
    var name = columns[0];

    if (!DateTime.TryParseExact(birthDateStr, "MM/dd/yyyy", System.Globalization.CultureInfo.InvariantCulture, System.Globalization.DateTimeStyles.None, out DateTime birthDate))
    {
        continue;
    }

    var ageAsOfJuly1 = CalculateAgeAsOfJuly1(birthDate);
    var jsonNode = JsonNode.Parse($"{{\"Name\": \"{name}\", \"Age\": {ageAsOfJuly1}}}");
    jsonArray.Add(jsonNode);
}

Console.WriteLine(jsonArray.ToString());

static int CalculateAgeAsOfJuly1(DateTime birthDate)
{
    var today = new DateTime(2025, 7, 1);
    var age = today.Year - birthDate.Year;
    if (today < birthDate.AddYears(age)) age--;
    return age;
}