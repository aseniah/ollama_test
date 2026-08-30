using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var csvPath = "input/input.csv";
var referenceDate = new DateTime(2025, 7, 1);

var lines = File.ReadAllLines(csvPath);
var csvData = lines.SelectMany(l => l.Split(','));

var results = new List<JsonNode>();

foreach (var csvRow in csvData.Skip(1))
{
    var nameParts = csvRow[0].Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];

    var birthday = DateTime.ParseExact(csvRow[1], "M/d/yyyy", null);
    var age = referenceDate - birthday;
    var ageInt = age.Days / 365;

    var father = csvRow[3];
    var mother = csvRow[4];
    var brother = csvRow[5];
    var sister = csvRow[6];

    var relatives = new List<JsonNode>();

    if (!string.IsNullOrEmpty(father))
    {
        var fatherParts = father.Split(' ');
        relatives.Add(JsonNode.Parse($"{{\"FirstName\": \"{fatherParts[0]}\", \"LastName\": \"{fatherParts[fatherParts.Length - 1]}\", \"Relationship\": \"Father\"}}"));
    }

    if (!string.IsNullOrEmpty(mother))
    {
        var motherParts = mother.Split(' ');
        relatives.Add(JsonNode.Parse($"{{\"FirstName\": \"{motherParts[0]}\", \"LastName\": \"{motherParts[motherParts.Length - 1]}\", \"Relationship\": \"Mother\"}}"));
    }

    if (!string.IsNullOrEmpty(brother))
    {
        var brotherParts = brother.Split(' ');
        relatives.Add(JsonNode.Parse($"{{\"FirstName\": \"{brotherParts[0]}\", \"LastName\": \"{brotherParts[brotherParts.Length - 1]}\", \"Relationship\": \"Brother\"}}"));
    }

    if (!string.IsNullOrEmpty(sister))
    {
        var sisterParts = sister.Split(' ');
        relatives.Add(JsonNode.Parse($"{{\"FirstName\": \"{sisterParts[0]}\", \"LastName\": \"{sisterParts[sisterParts.Length - 1]}\", \"Relationship\": \"Sister\"}}"));
    }

    var record = JsonNode.Parse($"{{\"FirstName\": \"{firstName}\", \"LastName\": \"{lastName}\", \"Birthday\": \"{birthday:yyyy-MM-dd}\", \"Age\": {ageInt}, \"Relatives\": {relativeJson}}})") where relativeJson in relatives
    results.Add(JsonNode.Parse($"{{\"FirstName\": \"{firstName}\", \"LastName\": \"{lastName}\", \"Birthday\": \"{birthday:yyyy-MM-dd}\", \"Age\": {ageInt}, \"Relatives\": {relativeJson}}})") where relativeJson in relatives.Add(relativeJson).ToString();
}

var jsonOutput = System.Text.Json.JsonSerializer.Serialize(results);

Console.WriteLine(jsonOutput);