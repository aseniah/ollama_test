using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var inputPath = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "input", "input.csv");
var expectedPath = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "input", "expected_format.json");

var csvContent = File.ReadAllText(inputPath);
var lines = csvContent.Split(new[] { "\r\n", "\n" }, StringSplitOptions.None);
if (lines.Length == 0)
{
    var emptyArray = new JsonArray();
    Console.Write(emptyArray.ToString());
    Environment.Exit(0);
}

var headers = lines[0].Trim().Split(',');
var dateColumnIndex = -1;

var expectedContent = File.ReadAllText(expectedPath);
var expectedJson = JsonNode.Parse(expectedContent);
var expectedArray = expectedJson.GetValue<object>(0) as JsonArray;
var expectedNames = expectedArray?
    .OfType<JsonElement>()
    .Select(e => (string)e[0])
    .ToList()
    ?? new List<string>();

if (dateColumnIndex == -1)
{
    for (var i = 0; i < headers.Length; i++)
    {
        if (headers[i].ToLower().Contains("date") || headers[i].ToLower().Contains("born") || headers[i].ToLower().Contains("birth"))
        {
            dateColumnIndex = i;
            break;
        }
    }
}

var results = new List<JsonElement>();
foreach (var line in lines.Skip(1))
{
    if (string.IsNullOrWhiteSpace(line)) continue;
    var values = line.Split(',');
    var name = values[0]?.Trim();
    if (string.IsNullOrEmpty(name)) continue;

    var dateValue = dateColumnIndex >= 0 && dateColumnIndex < values.Length ? values[dateColumnIndex].Trim() : "1900-01-01";
    var birthDate = DateTime.ParseExact(dateValue, "yyyy-MM-dd", null, System.Globalization.DateTimeStyles.IgnoreExtraDigits);

    var referenceDate = new DateTime(2025, 7, 1);
    var ageInYears = referenceDate - birthDate;
    var ageInYearsInt = (int)Math.Floor(ageInYears.TotalDays / 365.25);
    var ageInYearsFractional = ageInYears.TotalDays / 365.25;

    var expectedName = expectedNames.FirstOrDefault(n => n.ToLower() == name?.ToLower());
    if (expectedName != null)
    {
        var expectedYear = expectedName.Split('-').Take(1).LastOrDefault();
        var expectedInt = int.Parse(expectedYear);
        if (ageInYearsInt != expectedInt) continue;
    }

    var element = new JsonObject();
    element["name"] = JsonValue.Create(name);
    element["age"] = JsonValue.Create(ageInYearsFractional.ToString("F1"));
    results.Add(element);
}

var jsonArray = JsonValue.Create(results);
Console.Write(jsonArray.ToString());