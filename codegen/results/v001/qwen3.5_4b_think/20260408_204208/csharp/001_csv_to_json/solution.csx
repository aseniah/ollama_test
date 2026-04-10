using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var fileContent = File.ReadAllText("input/data.csv");
var lines = fileContent.Split(new[] { "\r\n", "\n" }, StringSplitOptions.RemoveEmptyEntries);

var header = lines[0].Split(',');
var headerNames = header.Select(h => h.Trim());

var dataLines = lines.Skip(1);
var objects = new List<JsonElement>();

using var sr = new StreamReader("input/data.csv");
var reader = new CsvReader(sr);

var records = reader.ReadAll().ToList();

var jsonArray = new List<object>();

foreach (var record in records)
{
    var name = record["name"].ToString();
    var age = int.Parse(record["age"].ToString());
    var email = record["email"].ToString();
    var score = float.Parse(record["score"].ToString());
    jsonArray.Add(new { Name = name, Age = age, Email = email, Score = score });
}

var json = JsonSerializer.Serialize(jsonArray);
Console.WriteLine(json);